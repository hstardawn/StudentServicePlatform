package model

type User struct {
	ID       int    `json:"user_id"`
	Username int    `json:"username"`
	Name     string `json:"name"`
	Sex      string `json:"sex"`
	PhoneNum int    `json:"phone_num"`
	Email    string `json:"email"`
	UserType int    `json:"user_type"`
	Password string `json:"password"`
}
