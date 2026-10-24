package todo

import "time"

type Task struct {
	Title       string
	Description string
	IsCompleted bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

func NewTask(title, description string) Task {
	return Task{
		Title:       title,
		Description: description,
		IsCompleted: false,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}
}

func (t *Task) Complete() {
	t.IsCompleted = true

	now := time.Now()
	t.CompletedAt = &(now)
}
