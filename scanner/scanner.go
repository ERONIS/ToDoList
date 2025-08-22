package scanner

import (
	"bufio"
	"os"
	"strings"
	"todolist/todo"
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
		PrintPromp()

		ok := scanner.Scan()
		if !ok {
			return
		}
		inputString := scanner.Text()
		result := s.process(inputString)
		if result != "" {
			if result == needExit {
				PrintExit()
				return
			}
			PrintResult(result)

		}
		event := NewEvent(result, inputString)
		s.events = append(s.events, event)
	}
}

func (s Scanner) process(inputString string) string {
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
	if cmd == "del" {
		return s.cmdDel(fields)
	}
	if cmd == "help" {
		return s.cmdHelp(fields)
	}
	if cmd == "events" {
		return s.cmdEvent(fields)
	}
	return unknownCommand
}

func (s *Scanner) cmdAdd(fields []string) string {
	if len(fields) < 3 {
		return wrongArgs
	}

	title := fields[1]
	taskText := ""
	for i := 2; i < len(fields); i++ {
		taskText += fields[i]

		if i != len(fields)-1 {
			taskText += " "
		}
	}
	task := todo.NewTask(title, taskText)

	s.todoList.AddTasK(task)
	PrintAdd(title)
	return ""
}

func (s *Scanner) cmdList(fields []string) string {
	if len(fields) != 1 {
		return wrongArgs
	}

	tasks := s.todoList.ListTasks()

	PrintTasks(tasks)

	return ""
}

func (s *Scanner) cmdComplete(fields []string) string {
	if len(fields) != 2 {
		return wrongArgs
	}
	title := fields[1]
	doneTaskResult := s.todoList.CompleteTask(title)
	if doneTaskResult != "" {
		return doneTaskResult
	}
	return ""

}

func (s *Scanner) cmdDel(fields []string) string {
	if len(fields) != 2 {
		return wrongArgs
	}

	title := fields[1]
	delTaskresult := s.todoList.DeleteTasK(title)
	if delTaskresult != "" {
		return delTaskresult
	}
	PrintDel(title)
	return ""

}
func (s *Scanner) cmdHelp(fields []string) string {
	if len(fields) != 1 {
		return wrongArgs
	}
	PrintHelp()
	return ""
}
func (s *Scanner) cmdEvent(fields []string) string {
	if len(fields) != 1 {
		return wrongArgs
	}
	PrintEvent(s.events)

	return ""
}
