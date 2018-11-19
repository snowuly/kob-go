// +build ignore

package main

import (
	"context"
	"fmt"
	"kob"
	"net/http"
	"time"
)

func main() {

	var app kob.App

	app.Get(
		"/name/:name",
		func(ctx context.Context, w http.ResponseWriter, r *http.Request, next kob.NextFunc) {
			start := time.Now()
			next(ctx)
			w.Write([]byte(fmt.Sprintf("\nprocess cost: %d ns", time.Since(start))))
		},
		func(ctx context.Context, w http.ResponseWriter, r *http.Request, next kob.NextFunc) {
			params := kob.GetParams(ctx)
			w.Write([]byte("hello "))
			w.Write([]byte(params["name"]))
		},
	)

	app.Listen(":8080")
}
