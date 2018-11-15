// +build ignore

package main

import (
	"context"
	"kob"
	"net/http"
)

func main() {

	app := kob.NewApp()

	app.Get("/name/:name", func(ctx context.Context, w http.ResponseWriter, r *http.Request, next func(context.Context)) {
		params := ctx.Value(kob.KeyParams).(map[string]string)
		render := ctx.Value(kob.KeyRenderF).(kob.FuncRenderF)
		render(w, params, "head")
	})

	app.Listen(":8080", nil)
}
