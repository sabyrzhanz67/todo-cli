package scanner

import "time"

type Event struct {
	Description string
	UserInput   string
	CreatedAt   time.Time
}

func NewEvent(description, userInput string) Event {
	return Event{
		Description: description,
		UserInput:   userInput,
		CreatedAt:   time.Now(),
	}
}
