package genericstack

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/google/uuid"
	"github.com/yougaindra/stash-stack/internal/commands"
)

type GenericStack[T fmt.Stringer] interface {
	Pop() (T, error)
	Push(element T) bool
	Cleanup() error
}

type GenericStashStack[T fmt.Stringer] struct {
	gitCommands commands.GitCommands
	name        string
	parser      func(elemString string) (T, error)
}

func NewGenericStashStack[T fmt.Stringer](name string, parser func(elemString string) (T, error)) (GenericStack[T], error) {
	stack := GenericStashStack[T]{
		name:        name,
		gitCommands: &commands.GitCommandsImpl{},
		parser:      parser,
	}
	if err := stack.gitCommands.GitInit(name); err != nil {
		return nil, err
	}
	createFirst := exec.Command("touch", name+"/1")
	err := createFirst.Run()
	if err != nil {
		log.Printf("%+v", err.Error())
		return nil, err
	}
	if err := stack.gitCommands.GitAdd(name, "1"); err != nil {
		return nil, err
	}
	if err := stack.gitCommands.GitCommit(name, "initCommit"); err != nil {
		return nil, err
	}
	return &stack, nil
}

func (s *GenericStashStack[T]) Pop() (T, error) {

	val, err := s.gitCommands.GitStashList(s.name)
	if err != nil {
		panic(err)
	}
	vals := strings.Split(string(val), "\n")
	firstLine := strings.Split(vals[0], " ")
	ret := firstLine[len(firstLine)-1]
	if err = s.gitCommands.GitDrop(s.name); err != nil {
		panic(err)
	}
	return s.parser(ret)
}

func (s *GenericStashStack[T]) Push(x T) bool {
	fname := uuid.New().String()
	createFile := exec.Command("touch", s.name+"/"+fname)
	err := createFile.Run()
	if err != nil {
		println(err.Error())
		return false
	}
	err = s.gitCommands.GitAdd(s.name, fname)
	if err != nil {
		return false
	}
	return s.gitCommands.GitStash(s.name, fmt.Sprintf("%v", x)) == nil
}

func (s *GenericStashStack[T]) Cleanup() error {
	return exec.Command("rm", "-rf", s.name).Run()
}
