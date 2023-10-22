package main

import "time"

type Event struct {
	ID string    `db:"event_id"`
	At time.Time `db:"event_at"`
}
