package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	//crear servidor

	fmt.Println("El servidor esta corriendo en el puerto 3000")
	fmt.Println("Run server: http://localhost:3000/")
	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}
