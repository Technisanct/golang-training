package user

type CreateUserRequest struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Phone     int64  `json:"phone"`
}
