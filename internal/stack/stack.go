package stack

import (
	"log"
	"os/exec"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/yougaindra/stash-stack/internal/commands"
)

type Stack interface {
	Pop() (int, error)
	Push(x int) bool
	Cleanup() error
}

type stashStack struct {
	gitCommands commands.GitCommands
	name        string
}

func NewStack(name string) (Stack, error) {
	stack := stashStack{}
	stack.name = name
	stack.gitCommands = &commands.GitCommandsImpl{}
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

func (s *stashStack) Pop() (int, error) {

	val, err := s.gitCommands.GitStashList(s.name)
	if err != nil {
		return -1, err
	}
	vals := strings.Split(string(val), "\n")
	// for _, v := range vals {
	// 	println(v)
	// }
	firstLine := strings.Split(vals[0], " ")
	ret := firstLine[len(firstLine)-1]
	if err = s.gitCommands.GitDrop(s.name); err != nil {
		return -1, err
	}
	return strconv.Atoi(ret)

}

func (s *stashStack) Push(x int) bool {
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
	stash := exec.Command("git", "-C", s.name, "stash", "push", "-m", strconv.Itoa(x))
	err = stash.Run()
	return err == nil
}

func (s *stashStack) Cleanup() error {
	return exec.Command("rm", "-rf", s.name).Run()
}
