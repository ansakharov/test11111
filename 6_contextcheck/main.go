package main

import (
	"context"
	"time"
)

func main() {
	ctx := context.Background()
	call1(ctx)
}

func call1(ctx context.Context) {
	call2(ctx) // OK

	ctx, cancel := context.WithTimeout(ctx, time.Second)
	_ = cancel

	call2(ctx) // Non-inherited new context, use function like `context.WithXXX` instead

	call3(ctx) // Function `call3` should pass the context parameter

	call4() // Function `call4->call3` should pass the context parameter

	call5(ctx)
}

func call2(ctx context.Context) {
	//
}

func call3(ctx context.Context) {
	call2(ctx)
}

//nolint:contextcheck
func call4() {
	call3(context.Background()) // OK
}

func call5(_ context.Context) {
	ctx := context.Background()
	_ = ctx
}
