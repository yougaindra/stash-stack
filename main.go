package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/yougaindra/stash-stack/internal/genericstack"
	"github.com/yougaindra/stash-stack/internal/stack"
)

type StackStruct struct {
	x int
	y float64
}

func (s StackStruct) String() string {
	return strconv.Itoa(s.x) + "," + strconv.FormatFloat(s.y, 'f', -1, 32)
}

func IntParser(s string) (StackStruct, error) {
	splits := strings.Split(s, ",")
	intVal, err := strconv.Atoi(splits[0])
	floatVal, err := strconv.ParseFloat(splits[1], 64)
	return StackStruct{intVal, floatVal}, err
}

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

	gs, err := genericstack.NewGenericStashStack("s4", IntParser)
	defer gs.Cleanup()
	gs.Push(StackStruct{1, 2.3})
	top, err := gs.Pop()
	fmt.Printf("%v --- %v", top.x, top.y)
}
