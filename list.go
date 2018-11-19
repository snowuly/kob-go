package kob

import (
	"context"
	"net/http"
)

type wrkey int

const wrKey wrkey = 0

type List struct {
	Queue
}

type wr struct {
	w http.ResponseWriter
	r *http.Request
}

func (list *List) Add(fn HandlerFunc) {
	list.Queue.Add(func(ctx context.Context, next func(context.Context)) {
		data := ctx.Value(wrKey).(*wr)
		fn(ctx, data.w, data.r, next)
	})
}
func (list *List) Run(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	ctx = context.WithValue(ctx, wrKey, &wr{w, r})
	list.Queue.Run(ctx)
}
