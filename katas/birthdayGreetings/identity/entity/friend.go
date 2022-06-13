package entity

import "time"

type Friend struct {
	FirstName string
	LastName  string
	Email     string
	Birth     time.Time
}
