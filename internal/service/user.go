package service

import "StudentServicePlatform/internal/model"

// "CMS/internal/model"
// "unicode"

// func IsNumericUsername(username string) bool {
// 	for _, char := range username {
// 		if !unicode.IsDigit(char) {
// 			return false
// 		}
// 	}
// 	return true
// }

func Register(username int, name string, email string, password string) error {
	return d.CreateUser(ctx, &model.User{
		Username: username,
		Name:     name,
		// Sex:      sex,
		// PhoneNum: phone_num,
		Email:    email,
		Password: password,
		// UserType: user_type,
	})
}

// func ComparePwd(pwd1 string, pwd2 string) bool {
// 	return pwd1 == pwd2
// }

func GetUserPassword(username int) (*model.User, error) {
	return d.GetUserPassword(ctx, username)
}

func GetUserByUsername(username int) (*model.User, error) {
	return d.GetUserByUsername(ctx, username)
}

func UpdateUser(username int, name string, sex string, phone_num int, email string, password string) error {
	return d.UpdateUser(ctx, username, name, sex, phone_num, email, password)
}
