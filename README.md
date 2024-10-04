# CLI Todo List


A simple command-line todo list application written in Go.

## Getting Started
To get started with the application, simply clone this repository and build the application:
```bash
go build
```

## Usage

The application provides the following commands:
- `add <description>`: Add a new task to the list with the given description.
- `complete <id>`: Mark a task with the given ID as complete.
- `delete <id>`: Delete a task with the given ID.
- `list`: List all non-completed tasks. Use the `-a` flag to list all tasks, including completed ones.

### Example Use Cases

- Add a new task: `cli-todo-list add "Buy milk"`
- Mark a task as complete: `cli-todo-list complete 1`
- Delete a task: `cli-todo-list delete 1`
- List all tasks: `cli-todo-list list`
- List all tasks, including completed ones: `cli-todo-list list -a`

## Dependencies

This project uses the following dependencies:
- <https://github.com/inconshreveable/mousetrap>
- <https://github.com/mergestat/timediff>
- <https://github.com/spf13/cobra>
- <https://github.com/spf13/pflag>
