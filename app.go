package kob

import (
	"context"
	"net/http"

	koa "github.com/snowuly/koa-go"
)

type App struct {
	*koa.App
}

type FuncRenderF func(http.ResponseWriter, interface{}, ...string)
type FuncRenderG func(http.ResponseWriter, interface{}, string)

type Key int

const (
	KeyParams = Key(iota)
	KeyRenderF
	KeyRenderG
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

func (app *App) Listen(addr string, ctx context.Context) {
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = context.WithValue(ctx, KeyRenderF, FuncRenderF(renderFiles))
	ctx = context.WithValue(ctx, KeyRenderG, FuncRenderG(renderGlob))
	app.App.Listen(addr, ctx)
}

func renderFiles(w http.ResponseWriter, data interface{}, files ...string) {
	tpl, err := GetTpl(files...)
	if err != nil {
		panic(err)
	}
	err = tpl.Execute(w, data)
	if err != nil {
		panic(err)
	}
}
func renderGlob(w http.ResponseWriter, data interface{}, pattern string) {
	tpl, err := GetTplGlob(pattern)
	if err != nil {
		panic(err)
	}
	err = tpl.Execute(w, data)
	if err != nil {
		panic(err)
	}
}
