package main

import (
	"fmt"
	"log"
	"net/http"
)

//handler

func HolaMundo(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("El metodo utilizado es: ", r.Method)
	fmt.Fprintln(rw, "Hola Mundo de Go Web")
}

func PageNF(rw http.ResponseWriter, r *http.Request) {
	http.NotFound(rw, r)
}

func Error(rw http.ResponseWriter, r *http.Request) {
	http.Error(rw, "Ocurrio un error", http.StatusNotFound)
}

func Saludar(rw http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	fmt.Println(r.URL.RawQuery)
	fmt.Println(r.URL.Query())

	name := r.URL.Query().Get("name")
	edad := r.URL.Query().Get("edad")
	fmt.Fprintf(rw, "Hola %s! actualmente tienes %s ages", name, edad)
}

func main() {
	//Mux
	mux := http.NewServeMux()

	mux.HandleFunc("/", HolaMundo)
	mux.HandleFunc("/page", PageNF)
	mux.HandleFunc("/error", Error)
	mux.HandleFunc("/saludar", Saludar)

	//Router
	/* http.HandleFunc("/", HolaMundo)
	http.HandleFunc("/page", PageNF)
	http.HandleFunc("/error", Error)
	http.HandleFunc("/saludar", Saludar) */

	//crear servidor
	server := &http.Server{
		Addr:    "localhost:3000",
		Handler: mux,
	}

	fmt.Println("El servidor esta corriendo en el puerto 3000")
	fmt.Println("Run server: http://localhost:3000")
	//log.Fatal(http.ListenAndServe("localhost:3000", mux))
	log.Fatal(server.ListenAndServe())
}
