package kob

import (
	"context"
	"net/http"

	koa "github.com/snowuly/koa-go"
)

type App struct {
	*koa.App
}

type Key int

const (
	KeyParams = Key(iota)
)

func NewApp() *App {
	return &App{koa.NewApp()}
}

func (app *App) Route(method, path string, fns ...koa.Handler) koa.Handler {
	reg, keys, err := PathToReg(path)
	if err != nil {
		panic(err)
	}
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request, next func(context.Context)) {
		if r.Method == method {
			values := reg.FindStringSubmatch(r.URL.Path)
			if len(values) > 0 {
				params := make(map[string]string)
				for i, key := range keys {
					params[key] = values[i+1]
				}
				compose(fns)(context.WithValue(ctx, KeyParams, params))
				return
			}
		}
		if next != nil {
			next(ctx)
		} else {
			http.NotFound(w, r)
		}
	}
}

func (app *App) Get(path string, fns ...koa.Handler) {
	app.Use(app.Route("GET", path, fns...))
}

func compose(fns []koa.Handler) func(context.Context) {
	return func(ctx context.Context) {
		var list koa.List
		for _, fn := range fns {
			list.Add(fn)
		}
		list.Run(ctx)
	}

}
