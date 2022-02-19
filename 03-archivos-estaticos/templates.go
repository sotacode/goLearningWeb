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

var funciones = template.FuncMap{
	"saludar": Saludar,
}
var templates = template.Must(template.New("t").Funcs(funciones).
	ParseGlob("templates/**/*.html"))

var errorTemplate = template.Must(template.ParseFiles("templates/error/error.html"))

func handlerError(rw http.ResponseWriter, status int) {
	rw.WriteHeader(status)
	errorTemplate.Execute(rw, nil)
}

func renderTemplate(rw http.ResponseWriter, name string, data interface{}) {
	err := templates.ExecuteTemplate(rw, name, data)
	if err != nil {
		//http.Error(rw, "No es posible retornar el template", http.StatusInternalServerError)
		handlerError(rw, http.StatusInternalServerError)
	}
}

func Index(rw http.ResponseWriter, r *http.Request) {
	usuario := User{"Sota", 24}

	renderTemplate(rw, "index.html", usuario)

	//fmt.Fprintln(rw, "Hola mundos")
	/* c1 := Curso{"GoWeb"}
	c2 := Curso{"React"}
	c3 := Curso{"Node"}
	c4 := Curso{"Flutter"} */
	//template, err := template.ParseFiles("index.html")

	/* cursos := []Curso{c1, c2, c3, c4} */
	/* usuario := User{"Sota", 24, true, false, cursos}
	 */

	/* template := template.Must(template.New("index.html").Funcs(funciones).
	ParseFiles("index.html", "base.html"))
	*/
	//template.Execute(rw, usuario)

}

func Registro(rw http.ResponseWriter, r *http.Request) {
	renderTemplate(rw, "registro.html", nil)
}

func main() {
	//Archivos estaticos
	staticFile := http.FileServer(http.Dir("static"))

	//Mux
	mux := http.NewServeMux()

	mux.HandleFunc("/", Index)
	mux.HandleFunc("/registro", Registro)

	//Mux DE STATIC FILE
	mux.Handle("/static/", http.StripPrefix("/static/", staticFile))

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
