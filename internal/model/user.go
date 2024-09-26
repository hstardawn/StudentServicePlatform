package model


type User struct {
	Id             int    `json:"user_id"`
	UserName       string    `json:"username"`
	Name        string `json:"name"`
	Sex         string `json:"sex"`
	PhoneNum int `json:"phone_num"`
	Email string `json:"email"`
	Usertype int `json:"user_type"`
	Password string `json:"password"`
}