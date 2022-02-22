package main

import (
	"fmt"
	"gomysql/db"
)

func main() {
	db.Connect()
	//db.Ping()

	//db.CreateTable(models.UserSchema, "user")
	if db.ExistTable("user") {
		fmt.Println("La tabla existe")
	} else {
		fmt.Println("La tabla no existe")
	}
	db.Close()
}
