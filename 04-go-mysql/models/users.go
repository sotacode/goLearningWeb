package models

import (
	"fmt"
	"gomysql/db"
)

type User struct {
	Id       int
	Username string
	Password string
	Email    string
}

type Users []User

const UserSchema string = `CREATE TABLE user (
	id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	username VARCHAR(30) NOT NULL,
	password VARCHAR(100) NOT NULL,
	email VARCHAR(50),
	create_data TIMESTAMP DEFAULT CURRENT_TIMESTAMP)`

func NewUser(userName, password, email string) *User {
	user := &User{
		Username: userName,
		Password: password,
		Email:    email,
	}
	return user
}

func (user *User) insert() {
	sql := "INSERT user SET username=?, password=?, email=?"
	if result, err := db.Exec(sql, user.Username, user.Password, user.Email); err != nil {
		panic(err)
	} else {
		fmt.Println("Inserccion realizada correctamente:", result)
	}
}

//crea usuarios
func CreateUser(username, password, email string) *User {
	user := NewUser(username, password, email)
	user.insert()
	return user
}

//lista todos los usuarios
func ListUsers() Users {
	query := "SELECT id, username, password, email FROM user"
	users := Users{}
	rows, _ := db.Query(query)

	for rows.Next() {
		user := User{}
		rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email)
		users = append(users, user)
	}
	return users
}

//retorna un usuario
func GetUser(id int64) *User {
	user := NewUser("", "", "")

	query := "SELECT id, username, password, email FROM user WHERE id=?"
	rows, _ := db.Query(query, id)
	for rows.Next() {
		rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email)
	}
	return user
}

//actualiza un usuario
func (user *User) update() {
	query := "UPDATE user SET username=?, password=?, email=? WHERE id=?"
	db.Exec(query, user.Username, user, user.Password, user.Email, user.Id)
}

//insertar si no existe, actualizar si existe
func (user *User) Save() {
	if user.Id == 0 {
		user.insert()
	} else {
		user.update()
	}
}

//eliminar registro
func (user *User) Delete() {
	query := "DELETE FROM user WHERE id=?"
	db.Exec(query, user.Id)
}
