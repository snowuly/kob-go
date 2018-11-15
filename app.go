package kob

import (
	"context"
	"koa"
	"net/http"
)

type App struct {
	*koa.App
}

func NewApp() *App {
	return &App{koa.NewApp()}
}

func (app *App) Get(path string, fns ...koa.Handler) {
	app.Use(func(ctx context.Context, w http.ResponseWriter, r *http.Request, next func(context.Context)) {
		if r.URL.Path == path {
			compose(fns)(ctx)
		} else {
			if next != nil {
				next(ctx)
			} else {
				http.NotFound(w, r)
			}
		}
	})
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
