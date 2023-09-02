package main

import (
	"log"

	"github.com/yougaindra/stash-stack/internal/stack"
)

func main() {
	stk, err := stack.NewStack("s3")
	if err != nil {
		log.Printf("error creating new stack. Exiting: %+v", err)
	}
	defer stk.Cleanup()
	done := stk.Push(1000000)
	println(done)
	v, _ := stk.Pop()
	println(v)
	println("Done")
	stk.Push(1)
	stk.Push(2)
	stk.Push(3)
	println(stk.Pop())
	println(stk.Pop())
	println(stk.Pop())
}
