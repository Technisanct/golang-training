package user

import "time"

type CreateUserRequest struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Phone     int64  `json:"phone"`
}

type GetUserRequest struct {
	UUID string `uri:"uuid"`
}

type User struct {
	ID        string    `json:"id"`
	UUID      string    `json:"uuid"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
}
