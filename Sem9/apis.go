package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

//HTTP GET LIST
//HTTP GET BY QUERY
//HTTP POST

type Alumno struct {
	Name string `json:"nombre"`
	Age  int    `json:"edad"`
	Dni  int    `json:"document"`
}

var lstAlumnos []Alumno

func cargarDatos() {
	lstAlumnos = []Alumno{
		{"Ppepito", 12, 772215},
		{"Pancho", 13, 123456},
		{"Renzo", 14, 65465}}
}

func listarAlumnos(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	//Serializar
	jsonBytes, _ := json.MarshalIndent(lstAlumnos, "", " ")
	io.WriteString(response, string(jsonBytes))
}

func buscarAlumno(response http.ResponseWriter, request *http.Request) {
	log.Println("Ingreso a buscar alumno")

	// Forma 1 de recuperar parametros de entrada
	code := request.FormValue("nombre")
	log.Println(code)

	// Forma 2 de recuperar
	values, _ := request.URL.Query()["nombre"]
	qcode := values[0]

	for _, a := range lstAlumnos {
		if a.Name == qcode {
			jsonBytes, _ := json.MarshalIndent(a, "", " ")
			io.WriteString(response, string(jsonBytes))
		}
	}
}

func agregar_alumno(response http.ResponseWriter, request *http.Request) {
	//validar POST
	if request.Method == "POST" {
		if request.Header.Get("Content-Type") == "application/json" {
			log.Println("PROCESAR post")
			jsonBytes, err := ioutil.ReadAll(request.Body)
			if err != nil {
				http.Error(response, "Problema con la lectura", http.StatusInternalServerError)
			} else {
				var oAlumno Alumno
				json.Unmarshal(jsonBytes, &oAlumno)
				lstAlumnos = append(lstAlumnos, oAlumno)

				response.Header().Set("Content-Type", "application/json")
				io.WriteString(response, `
					{
						"respuesta":"Registro Satisfactorio"
					}
				`)
				log.Println(lstAlumnos)
			}
		}
	} else {
		http.Error(response, "Metodo Invalido", http.StatusMethodNotAllowed)
	}
}

func handleRequest() {
	http.HandleFunc("/listar_alumnos", listarAlumnos)  //HTTP GET LIST
	http.HandleFunc("/buscar_alumno", buscarAlumno)    //HTTP GET BY QUERY
	http.HandleFunc("/agregar_alumno", agregar_alumno) //HTTP POST

	log.Fatal(http.ListenAndServe(":9000", nil))

}

func main() {

	cargarDatos()
	handleRequest()

}
