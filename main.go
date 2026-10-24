package main

import (
	"todo-cli/scanner"
	"todo-cli/todo"
)

func main() {
	todoList := todo.NewList()

	scanner := scanner.NewScanner(todoList)

	scanner.Start()
}
