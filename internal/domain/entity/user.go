package entity

import "time"

type User struct {
	Username     string
	Password     string
	Registration time.Time
}
