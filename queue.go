package kob

import "context"

type Queue struct {
	list []handler
}

type handler func(context.Context, func(context.Context))

func (q *Queue) Add(fn handler) {
	q.list = append(q.list, fn)
}

func (q *Queue) genNext(index int) func(context.Context) {
	if index >= len(q.list) {
		return nil
	}
	return func(ctx context.Context) {
		q.list[index](ctx, q.genNext(index+1))
	}
}

func (q *Queue) Run(ctx context.Context) {
	if len(q.list) > 0 {
		q.list[0](ctx, q.genNext(1))
	}

}
