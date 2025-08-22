package todo

import (
	"time"
)

type Task struct {
	Title     string
	Text      string
	Completed bool

	CreatedAt   time.Time
	CompletedAt *time.Time
}

func NewTask(title string, text string) Task {

	task := Task{
		Title:     title,
		Text:      text,
		Completed: false,

		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}
	return task

}

func (t *Task) Complete() {
	t.Completed = true
	now := time.Now()
	t.CompletedAt = &now
}
