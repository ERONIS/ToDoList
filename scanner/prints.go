package scanner

import (
	"fmt"
	"todolist/todo"

	"github.com/k0kubun/pp"
)

func PrintResult(result string) {
	fmt.Println("Результат выполнения: ", result)
	fmt.Println("")
}

func PrintPromp() {
	fmt.Print("Ведите команду: ")
}

func PrintExit() {
	fmt.Print("Завершение работы... До скорого  ")
}

func PrintAdd(title string) {
	fmt.Println("Задача '" + title + "'было успешно добавлено")
	fmt.Println("")
}
func PrintTasks(tasks map[string]todo.Task) {
	pp.Print("Список дел:", tasks)
	fmt.Println("")
}
func PrintDone(title string) {
	fmt.Println("Задача'" + title + "'помечена как выполненной")
}

func PrintDel(title string) {
	fmt.Println("Задача'" + title + "'успешно удалена")
	fmt.Println("")
}
func PrintHelp() {
	fmt.Println("Список команд:")
	fmt.Println("1. help")
	fmt.Println("-- эта команда позволяет узнать доступные команды и их формат")
	fmt.Println("")
	fmt.Println("2. add {заголовок задачи из одного слова} {текст задачи из одного или нескольких слов}")
	fmt.Println("-- эта команда позволяет добавлять новые задачи в список задач")
	fmt.Println("")
	fmt.Println("3. list")
	fmt.Println("-- эта команда позволяет получить полный список всех задач")
	fmt.Println("")
	fmt.Println("4. del {заголовок существующей задачи}")
	fmt.Println("-- эта команда позволяет удалить задачу по её заголовку")
	fmt.Println("")
	fmt.Println("5. done {заголовок существующей задачи}")
	fmt.Println("-- эта команда позволяет отменить задачу как выполненную")
	fmt.Println("")
	fmt.Println("6. events")
	fmt.Println("-- эта команда позволяет получить список всех событий")
	fmt.Println("")
	fmt.Println("7. exit")
	fmt.Println("-- эта команда позволяет завершить выполнение программы")
	fmt.Println("")
}
func PrintEvent(events []Event) {
	pp.Println("События: ", events)
	fmt.Println()
}
