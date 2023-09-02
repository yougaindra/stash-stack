package genericstack

import "fmt"

type StackElement[T any] interface {
	fmt.Stringer
	New(elemStr string) (T, error)
}
