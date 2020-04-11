package entities

import "time"

type Event struct {
	Id          Id          `db:"id"`
	Title       Title       `db:"title"`
	Description Description `db:"description"`
	Time        time.Time   `db:"date"`
}
