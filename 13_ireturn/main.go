package main

import (
	"context"
	"fmt"
	"time"
)

type MyStruct struct {
	value *time.Time
}

func (m *MyStruct) DoSomething(ctx context.Context) {
	fmt.Println(m.value.Unix())
}

type MyInterface interface {
	DoSomething(ctx context.Context)
}

/*func NewMyStruct() MyInterface {*/
func NewMyStruct() *MyStruct {
	var res *MyStruct // nil

	fmt.Println(res == nil)

	return res
}

func main() {
	ctx := context.Background()
	m := NewMyStruct()

	fmt.Println(m == nil)

	if m != nil {
		m.DoSomething(ctx)
	}
}
