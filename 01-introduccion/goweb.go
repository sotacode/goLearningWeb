package main

import (
	"fmt"
	"log"
	"net/http"
)

//handler

func HolaMundo(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("El metodo utilizado es: ", r.Method)
	fmt.Fprintln(rw, "Hola Mundo")
}

func PageNF(rw http.ResponseWriter, r *http.Request) {
	http.NotFound(rw, r)
}

func Error(rw http.ResponseWriter, r *http.Request) {
	http.Error(rw, "Ocurrio un error", http.StatusNotFound)
}

func main() {

	//Router
	http.HandleFunc("/", HolaMundo)
	http.HandleFunc("/page", PageNF)
	http.HandleFunc("/error", Error)

	//crear servidor

	fmt.Println("El servidor esta corriendo en el puerto 3000")
	fmt.Println("Run server: http://localhost:3000")
	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}
