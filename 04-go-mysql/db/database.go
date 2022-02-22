package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//username:pasword@tcp(localhost:3306)/database

const url = "root:211097@tcp(localhost:3306)/goweb_db"

var db *sql.DB

func Connect() {
	connection, err := sql.Open("mysql", url)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("La base de datos se conecto correctamente")
		db = connection
	}

}

func Close() {
	db.Close()
}

func Ping() {
	if err := db.Ping(); err != nil {
		panic(err)
	}
}

func ExistTable(tableName string) bool {
	sqlQuery := fmt.Sprintf("SHOW TABLES LIKE '%s'", tableName)
	if rows, err := Query(sqlQuery); err != nil {
		fmt.Println(err)
		return false
	} else {
		return rows.Next()
	}
}

func CreateTable(schema string, name string) {
	if !ExistTable(name) {
		if sqlResult, err := Exec(schema); err == nil {
			println(sqlResult)
		} else {
			panic(err)
		}
	} else {
		fmt.Println("La tabla ya existe")
	}

}

//Polimorfizando EXEC
func Exec(query string, args ...interface{}) (sql.Result, error) {
	sqlResult, err := db.Exec(query, args...)
	if err != nil {
		fmt.Println(err)
	}
	return sqlResult, err
}

//Polimorfizando QUERY
func Query(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := db.Query(query, args...)
	if err != nil {
		fmt.Println(err)
	}
	return rows, err
}

//Reiniciar Registro de una tabla
func TruncateTable(tableName string) {
	sql := fmt.Sprintf("TRUNCATE %s", tableName)
	Exec(sql)
}
