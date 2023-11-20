package server

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"main/db"
)

func UserSignUp(user *db.User) *gorm.DB {
	user.Uid = uuid.NewV4()
	user.SetPassword(user.Password)
	result := db.Conn.Create(&user)

	return result
}

func UserLogin(username, password string) (db.User, error) {
	var user db.User
	result := db.Conn.Find(&user).Where("username = ?", username)
	fmt.Printf("result => %+v", result)
	if result.Error != nil {
		fmt.Println("err=======+>>>>>", result.Error)
		return db.User{}, result.Error
	}
	if !user.CheckPassword(password) {
		return db.User{}, fmt.Errorf("Wrong Password!")
	}
	fmt.Printf("\n\nuser after checks ==> %+v\n\n", user)
	return user, nil
}
