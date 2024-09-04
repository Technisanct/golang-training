package user

import (
	"time"
)

type CreateUserRequest struct {
	Firstname string
	Lastname  string
	Email     string
	Phone     int64
}

type User struct {
	ID        string
	UUID      string
	Firstname string
	Lastname  string
	Email     string
	CreatedAt time.Time
}
