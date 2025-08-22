package todo

type List struct {
	tasks map[string]Task
}

func NewList() *List {
	list := List{
		tasks: make(map[string]Task),
	}
	return &list
}

func (l *List) AddTasK(task Task) string {
	if _, ok := l.tasks[task.Title]; ok {
		return ErrTaskNotFound
	}
	l.tasks[task.Title] = task
	return ""
}

func (l *List) ListTasks() map[string]Task {
	tmp := make(map[string]Task, len(l.tasks))
	for k, v := range tmp {
		tmp[k] = v
	}
	return tmp
}

func (l *List) NotCompletedTask() map[string]Task {
	notCompletedTask := make(map[string]Task)

	for title, task := range l.tasks {
		if !task.Completed {
			notCompletedTask[title] = task

		}
	}
	return notCompletedTask
}

func (l *List) CompleteTask(title string) string {
	task, ok := l.tasks[title]
	if !ok {
		return ErrTaskNotFound
	}

	task.Complete()

	l.tasks[title] = task

	return ""
}

func (l *List) DeleteTasK(title string) string {
	_, ok := l.tasks[title]
	if !ok {
		return ErrTaskNotFound
	}
	delete(l.tasks, title)
	return ""
}
