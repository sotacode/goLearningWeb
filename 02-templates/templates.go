package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type User struct {
	Name   string
	Edad   int
	Status bool
	Admin  bool
}

func Index(rw http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(rw, "Hola mundos")
	template, err := template.ParseFiles("index.html")

	usuario := User{"Sota", 24, true, false}

	if err != nil {
		panic(err)
	} else {
		template.Execute(rw, usuario)
	}
}

func main() {

	//Mux
	mux := http.NewServeMux()

	mux.HandleFunc("/", Index)

	//servidor
	server := &http.Server{
		Addr:    "localhost:3000",
		Handler: mux,
	}

	fmt.Println("El servidor esta corriendo en el puerto 3000")
	fmt.Println("Run server: http://localhost:3000")
	//log.Fatal(http.ListenAndServe("localhost:3000", mux))
	log.Fatal(server.ListenAndServe())
}
