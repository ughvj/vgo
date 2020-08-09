package dto

import (
	"time"
)

type User struct {
	ID                int
	Name              string
	Pass              string
	Administrator     bool
	Token             string
	Token_Expire_Date *time.Time
}