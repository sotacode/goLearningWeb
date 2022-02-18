package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type User struct {
	Name string
	Edad int
	/* Status bool
	Admin  bool
	Cursos []Curso */
}

/* type Curso struct {
	Name string
} */

//Funciones
func Saludar(nombre string) string {
	return "Hola " + nombre + " desde funcion"
}

func Index(rw http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(rw, "Hola mundos")
	/* c1 := Curso{"GoWeb"}
	c2 := Curso{"React"}
	c3 := Curso{"Node"}
	c4 := Curso{"Flutter"} */
	//template, err := template.ParseFiles("index.html")

	/* cursos := []Curso{c1, c2, c3, c4} */
	/* usuario := User{"Sota", 24, true, false, cursos}
	 */

	funciones := template.FuncMap{
		"saludar": Saludar,
	}
	usuario := User{"Sota", 24}
	template := template.Must(template.New("index.html").Funcs(funciones).
		ParseFiles("index.html", "base.html"))

	template.Execute(rw, usuario)
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
