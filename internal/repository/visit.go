package repository

import "time"

type Visit struct {
	Time time.Time `json:"time" db:"time"`
	IP   string    `json:"ip" db:"ip"`
	City string    `json:"city" db:"city"`
}
