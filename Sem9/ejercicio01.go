package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

	
type Alumno struct {
    Name string `json:"nam"`
    Age  int `json:"ag"`
}

var lstAlumnos []Alumno

func cargarDatos(){
	lstAlumnos = []Alumno{
		{"Ppepito",12},
		{"Pancho",13},
		{"Renzo",14}}
}

func listarAlumnos(response http.ResponseWriter, request *http.Request){
	response.Header().Set("Content-Type","application/json")
	//Serializar
	jsonBytes,_ := json.MarshalIndent(lstAlumnos,""," ")
	io.WriteString(response,string(jsonBytes))

}



func buscarAlumno(response http.ResponseWriter, request *http.Request){
	log.Println("Ingreso a buscar alumno")

	// Forma 1 de recuperar parametros de entrada 
	code := request.FormValue("nombre")
	log.Println(code)

	// Forma 2 de recuperar
	values,_ := request.URL.Query()["nombre"]
	qcode := values[0]

	for _,a:= range lstAlumnos {
		if a.Name==code {
			jsonBytes,_ := json.MarshalIndent(a,""," ")
			io.WriteString(response,string(jsonBytes))
		}
	} 
}
func handleRequest(){
	http.HandleFunc("/listar_alumnos",listarAlumnos)
	http.HandleFunc("/buscar_alumno",buscarAlumno)
	log.Fatal(http.ListenAndServe(":9000",nil))
	
}

func main() {


	cargarDatos()
	handleRequest()

    
}