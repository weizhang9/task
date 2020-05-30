A TODO list management CLI tool as part of Gophercises

This tool is built on top of Gophercises' author's solution: 

- changed certain function signature and implementations.
- added multiple adding/deleting features and fixed the deleting bug.

### How to use

```
Task is a CLI task manager

Usage:
  task [command]

Available Commands:
  add         Adds task(s) to your task list. Use comma to separate multiple tasks.
  done        Makrs task(s) as done. Use comma or space to separate multiple tasks.
  help        Help about any command, e.g task help add
  list        List all your task(s)

Flags:
  -a, --all    Marks all task(s) as done
  -h, --help   help for task
```