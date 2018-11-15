// +build ignore

package main

import (
	"context"
	"kob"
	"net/http"
)

func main() {

	app := kob.NewApp()

	app.Get("/tony", func(ctx context.Context, w http.ResponseWriter, r *http.Request, next func(context.Context)) {
		w.Write([]byte("hello tont time to home"))
	})
	app.Get("/tonyx", func(ctx context.Context, w http.ResponseWriter, r *http.Request, next func(context.Context)) {
		w.Write([]byte("hello tont time to homex"))
	})

	app.Listen(":8080")
}
