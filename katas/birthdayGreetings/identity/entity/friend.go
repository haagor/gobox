package entity

import "time"

type Friend struct {
	First_name string
	Last_name  string
	Email      string
	Birth      time.Time
}
