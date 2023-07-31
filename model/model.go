package model

type User struct {
	ID             int    `json:"id"`
	Username       string `json:"user_name"`
	FullName       string `json:"full_name"`
	ProfilePicture string `json:"profile_picture"`
	Followed       bool   `json:"followed"`
	Password       string `json:"pass"`
	Email          string `json:"email"`
}

type Post struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Owner       *User  `json:"owner"`
	Image       string `json:"image"`
	CreatedAt   int    `json:"created_at"`
	Liked       bool   `json:"liked"`
}

type LoginSuccessResponse struct {
	ID int `json:"id"`
}
