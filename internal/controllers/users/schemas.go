package users

import "time"

type UserCreation struct {
	Username string
	Password string
}

type User struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserResponse struct {
	Data User `json:"data"`
}

type SessionCreate struct {
	UserCreation
}

type Session struct {
	Token string `json:"token"`
}

type SessonResponse struct {
	Data Session `json:"data"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}
