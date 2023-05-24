package main

import (
	"database/sql"
	"log"
	"net/http"
	"text/template"
	_"github.com/go-sql-driver/mysql"
)
var plantillas= template.Must(template.ParseGlob("Plantillas/*"))
func main() {
	http.HandleFunc("/", Inicio)
	http.HandleFunc("/crear", Crear)
	http.HandleFunc("/insertar", Insertar)
	log.Println("Servidor corriendo...")
	http.ListenAndServe(":8080", nil)
}
type Empleado struct {
	Nombre   string
	Apellido string
	Edad     int
}


func Inicio(w http.ResponseWriter, r *http.Request) {
	conexionEstablecida := conexionBD()

	// Realiza la consulta para obtener los registros de la base de datos
// Realiza la consulta para obtener los registros de la base de datos ordenados por fecha o ID descendente
registros, err := conexionEstablecida.Query("SELECT nombre, apellido, edad FROM empleados ORDER BY fecha DESC")
if err != nil {
    panic(err.Error())
}
defer registros.Close()


	// Crea una lista para almacenar los registros
	var empleados []Empleado

	// Itera sobre los registros y agrega cada uno a la lista
	for registros.Next() {
		var empleado Empleado
		err := registros.Scan(&empleado.Nombre, &empleado.Apellido, &empleado.Edad)
		if err != nil {
			panic(err.Error())
		}
		empleados = append(empleados, empleado)
	}

	// Pasa la lista de empleados a la plantilla
	plantillas.ExecuteTemplate(w, "inicio", empleados)
}

func Crear(w http.ResponseWriter, r *http.Request) {
	plantillas.ExecuteTemplate(w, "crear",nil)
}

//CREAR CONEXION A LA BASE DE DATOS
func conexionBD()(conexion *sql.DB){
	Driver:="mysql"
	Usuario:="root"
	Contrasena:=""
	Nombre:="golang"

	conexion,err:= sql.Open(Driver, Usuario+":"+Contrasena+"@tcp(127.0.0.1)/"+Nombre)
	if err!=nil{
		panic(err.Error())
	}
	return conexion
}

func Insertar(w http.ResponseWriter, r *http.Request){
	if r.Method=="POST"{
	nombre:= r.FormValue("nombre")
	apellido:= r.FormValue("apellido")
	edad:= r.FormValue("edad")
	
	conexionEstablecida:= conexionBD ()
	insertarRegistros,err:= conexionEstablecida.Prepare("INSERT INTO empleados(nombre,apellido,edad) VALUES(?,?,?) ")
	if err!=nil{
		panic(err.Error())
	}
	insertarRegistros.Exec(nombre,apellido,edad)
	http.Redirect(w,r, "/",301)
	}
}
