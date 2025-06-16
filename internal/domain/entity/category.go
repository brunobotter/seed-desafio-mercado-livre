package entity

import (
	"time"
)

type Category struct {
	Name       string
	ParentName *string
	CreatedAt  time.Time
}
