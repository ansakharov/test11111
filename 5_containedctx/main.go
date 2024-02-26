package main

import (
	"context"
	"time"
)

type MyStruct struct {
	// found a struct that contains a context.Context field (containedctx)

	//ctx context.Context
}

func main() {
	my := MyStruct{}

	//my.someFunc()
	my.someFunc(context.Background())
}

/*func (m *MyStruct) someFunc() {
	ctx := m.ctx
	ctx, cancel := context.WithTimeout(ctx, time.Second)

	_ = ctx
	_ = cancel
}*/

func (m *MyStruct) someFunc(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, time.Second)

	_ = ctx
	_ = cancel
}
