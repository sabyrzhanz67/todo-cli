package scanner

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"todo-cli/todo"
)

type Scanner struct {
	todoList *todo.List
	events   []Event
}

func NewScanner(todoList *todo.List) *Scanner {
	return &Scanner{
		todoList: todoList,
	}
}

func (s *Scanner) Start() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		printPrompt()

		if ok := scanner.Scan(); !ok {
			break
		}

		inputString := scanner.Text()

		result := s.process(inputString)
		if result != "" {
			if result == needExit {
				printExit()
				return
			}

			printResult(result)
		}

		event := NewEvent(result, inputString)
		s.events = append(s.events, event)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка при чтении ввода:", err)
	}
}

func (s *Scanner) process(inputString string) string {
	fields := strings.Fields(inputString)

	if len(fields) == 0 {
		return emptyInput
	}

	cmd := fields[0]

	if cmd == "exit" {
		return needExit
	}

	if cmd == "add" {
		return s.cmdAdd(fields)
	}

	if cmd == "list" {
		return s.cmdList(fields)
	}

	if cmd == "complete" {
		return s.cmdComplete(fields)
	}

	if cmd == "delete" {
		return s.cmdDelete(fields)
	}

	if cmd == "help" {
		return s.cmdHelp(fields)
	}

	if cmd == "events" {
		return s.cmdEvents(fields)
	}

	return unknownCommand
}

func (s *Scanner) cmdAdd(fields []string) string {
	if len(fields) < 3 {
		return wrongArgs
	}

	title := fields[1]
	description := strings.Join(fields[2:], " ")

	task := todo.NewTask(title, description)

	s.todoList.AddTask(task)

	printAdd(title)

	return ""
}

func (s *Scanner) cmdList(fields []string) string {
	if len(fields) != 1 {
		return wrongArgs
	}

	tasks := s.todoList.ListTasks()

	printTasks(tasks)

	return ""
}

func (s *Scanner) cmdComplete(fields []string) string {
	if len(fields) != 2 {
		return wrongArgs
	}

	title := fields[1]

	completeTaskResult := s.todoList.CompleteTask(title)
	if completeTaskResult != "" {
		return completeTaskResult
	}

	printComplete(title)

	return ""
}

func (s *Scanner) cmdDelete(fields []string) string {
	if len(fields) != 2 {
		return wrongArgs
	}

	title := fields[1]

	deleteTaskResult := s.todoList.DeleteTask(title)
	if deleteTaskResult != "" {
		return deleteTaskResult
	}

	printDelete(title)

	return ""
}

func (s *Scanner) cmdHelp(fields []string) string {
	if len(fields) != 1 {
		return wrongArgs
	}

	printHelp()

	return ""
}

func (s *Scanner) cmdEvents(fields []string) string {
	if len(fields) != 1 {
		return wrongArgs
	}

	printEvents(s.events)

	return ""
}
