package kob

import (
	"context"
	"net/http"
	"regexp"
)

type NextFunc func(context.Context)
type HandlerFunc func(context.Context, http.ResponseWriter, *http.Request, NextFunc)
type EntryFunc func(context.Context, http.ResponseWriter, *http.Request)

type Route struct {
	Method string
	Reg    *regexp.Regexp
	keys   []string
	Entry  EntryFunc
}

type key int

const paramsKey key = 0

type App struct {
	Routes []*Route
}

func (app *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, route := range app.Routes {
		if r.Method == route.Method {
			values := route.Reg.FindStringSubmatch(r.URL.Path)
			if len(values) > 0 {
				params := make(map[string]string)
				for i, key := range route.keys {
					params[key] = values[i+1]
				}
				ctx := context.WithValue(r.Context(), paramsKey, params)
				route.Entry(ctx, w, r)
				return
			}
		}
	}
	http.NotFound(w, r)
}

func (app *App) Route(method, path string, fns ...HandlerFunc) {
	reg, keys, err := PathToReg(path)
	if err != nil {
		panic(err)
	}
	app.Routes = append(app.Routes, &Route{method, reg, keys, compose(fns)})
}

func (app *App) Get(path string, fns ...HandlerFunc) {
	app.Route("GET", path, fns...)
}

func (app *App) Listen(addr string) {
	http.ListenAndServe(addr, app)
}

func GetParams(ctx context.Context) map[string]string {
	if params, ok := ctx.Value(paramsKey).(map[string]string); ok {
		return params
	}
	panic("params is not map[string]string")

}

func compose(fns []HandlerFunc) EntryFunc {
	var list List
	for _, fn := range fns {
		list.Add(fn)
	}
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
		list.Run(ctx, w, r)
	}
}
