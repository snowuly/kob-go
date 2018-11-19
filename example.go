// +build ignore

package main

import (
	"context"
	"kob"
	"net/http"
)

func main() {

	var app kob.App

	app.Get("/name/:name", func(ctx context.Context, w http.ResponseWriter, r *http.Request, next kob.NextFunc) {

		params := kob.GetParams(ctx)
		w.Write([]byte("hello "))
		w.Write([]byte(params["name"]))
	})

	app.Listen(":8080")
}
