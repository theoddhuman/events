package model

import "time"

type Event struct {
	Id          int       `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"datetime"`
	UserId      int       `json:"user_id"`
}

var events []Event

func (event Event) Save() {
	events = append(events, event)
}

func GetEvents() []Event {
	return events
}
