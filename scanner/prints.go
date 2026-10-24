package scanner

import (
	"fmt"
	"todo-cli/todo"

	"github.com/k0kubun/pp"
)

func printResult(result string) {
	fmt.Println("Результат выполнения команды:", result)
}

func printPrompt() {
	fmt.Print("Введите команду: ")
}

func printExit() {
	fmt.Println("Завершение работы...До скорого!")
}

func printAdd(title string) {
	fmt.Printf("Задача %s добавлена!\n", title)
}

func printTasks(tasks map[string]todo.Task) {
	pp.Println("Список дел:", tasks)
}

func printComplete(title string) {
	fmt.Printf("Задача %q отмечена как выполненная\n", title)
}

func printDelete(title string) {
	fmt.Printf("Задача %s удалена!\n", title)
}

func printHelp() {
	fmt.Println(`Список команд:
- help — эта команда позволяет узнать доступные команды и их формат
- add {заголовок} {текст задачи} — добавляет новую задачу в список
- list — выводит полный список всех задач
- delete {заголовок} — удаляет задачу по её заголовку
- complete {заголовок} — отмечает задачу как выполненную
- events — выводит список всех событий
- exit — завершает выполнение программы`)
}

func printEvents(events []Event) {
	pp.Println("События:", events)
}
