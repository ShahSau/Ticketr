package models

import "time"

type Event struct {
	ID          int       `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"dateTime" binding:"required"`
	UserID      int       `json:"userId"`
}

var events = []Event{
	{ID: 1, Name: "Event 1", Description: "Description 1", Location: "Location 1", DateTime: time.Now(), UserID: 1},
	{ID: 2, Name: "Event 2", Description: "Description 2", Location: "Location 2", DateTime: time.Now(), UserID: 1},
	{ID: 3, Name: "Event 3", Description: "Description 3", Location: "Location 3", DateTime: time.Now(), UserID: 2},
}

func (e Event) SaveEvent() {
	events = append(events, e)
}

func GetAllEvents() []Event {
	return events
}
