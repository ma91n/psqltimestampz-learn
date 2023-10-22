package main

import "time"

type Event struct {
	ID string
	At time.Time
	TZ string
}
