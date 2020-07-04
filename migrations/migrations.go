package migrations

import (
	"duomly.com/go-bank-backend/helpers"
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialect/postgres"
)

type User struct {
	gorm.Model
	Username string
	Email    string
	Password string
}

type Account struct {
	gorm.Model
	Type    string
	Name    string
	Balance uint
	UserID  uint
}

func connectDB() *gorm.DB {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=bankapp password=postgres "+
		"sslmode=disabled")
	helpers.HandleErr(err)
	return db
}
func createAccounts() {
	db := connectDB()
	users := [2]User{
		{Username: "Nursultan", Email: "abdnurs2001@gmail.com"},
		{Username: "Sultan", Email: "abdnurs2001@gmail.com"},
	}
	for i := 0; i < len(users); i++ {
		generatedPassword := helpers.HashAndSalt([]byte(user[i].Username))
		user := User{Username: users[i].Username, Email: users[i].Email, Password: generatedPassword}
		db.Create(&user)

	}
}
