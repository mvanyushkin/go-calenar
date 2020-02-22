package entities

import "time"

type Event struct {
	Id          Id
	Title       Title
	Description Description
	Time        time.Time
}
