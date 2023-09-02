# Git Stash as Stack

## Why 


![Why not](https://media.tenor.com/_mY-jX8E_yIAAAAC/not.gif)


Goes without saying DON'T USE IT FOR ANY PRACTICAL PURPOSES


## How
Check `main.go` for example.

- use `NewStack()` to create new stack, it creates, new git repo which is used as stack via git stash command
- Stack supports 3 methods:
   - `Push(int) (bool)` : pushes an integer to top of stack and return if it was successfully stored
   - `Pop() (int,error)` : pop top element from stack and returns it alongwith error if any
   - `Cleanup (error)` : delete the git repo created when initalising stack
