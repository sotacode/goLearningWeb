package main

import (
	"fmt"
	"gomysql/db"
	"gomysql/models"
)

func main() {
	db.Connect()
	//db.Ping()

	//db.CreateTable(models.UserSchema, "user")
	/* if db.ExistTable("user") {
		fmt.Println("La tabla existe")
	} else {
		fmt.Println("La tabla no existe")
	} */

	models.CreateUser("Sota", "211097", "holarina1@gmail.com")
	models.CreateUser("Isi", "211097", "Isi@gmail.com")
	models.CreateUser("Sofi", "211097", "sofi@gmail.com")
	models.CreateUser("viky", "211097", "viky@gmail.com")
	//println(user)
	/* users := models.ListUsers()
	fmt.Println(users) */

	user := models.GetUser(2)
	fmt.Println(user)

	db.TruncateTable("user")
	db.Close()
}
