package commands

import (
	"os/exec"
)

type GitCommands interface {
	GitAdd(repoName string, fileName string) error
	GitCommit(repoName string, msg string) error
	GitStash(repoName string, msg string) error
	GitStashList(repoName string) ([]byte, error)
	GitInit(repoName string) error
	GitDrop(repoName string) error
}

type GitCommandsImpl struct {
}

func (g *GitCommandsImpl) GitAdd(repoName string, fileName string) error {
	return exec.Command("git", "-C", repoName, "add", fileName).Run()
}

func (g *GitCommandsImpl) GitCommit(repoName string, msg string) error {
	return exec.Command("git", "-C", repoName, "commit", "-m", msg).Run()
}

func (g *GitCommandsImpl) GitStashList(repoName string) ([]byte, error) {
	return exec.Command("git", "-C", repoName, "stash", "list").Output()
}

func (g *GitCommandsImpl) GitStash(repoName string, msg string) error {
	return exec.Command("git", "-C", repoName, "stash", "push", "-m", msg).Run()
}

func (g *GitCommandsImpl) GitInit(repoName string) error {
	return exec.Command("git", "init", repoName).Run()
}

func (g *GitCommandsImpl) GitDrop(repoName string) error {
	return exec.Command("git", "-C", repoName, "stash", "pop").Run()
}
