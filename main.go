package main

//AQUI IMPORTAMOS LO QUE NECESITAREMOS EN TODO EL PROYECTO COMO LIBRERIAS, PAQUETES
import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

// AQUI ESTABLECEMOS LAS RUTAS DE DONDE ESTAN NUESTROS ARCHIVOS DE TEMPLATE
var plantillas = template.Must(template.ParseGlob("Plantillas/*"))

// ESTA ES LA FUNCION PRINCIPAL, DONDE DEFINIMOS RUTAS
func main() {
	http.HandleFunc("/", Inicio)
	http.HandleFunc("/eliminar", Eliminar)
	http.HandleFunc("/crear", Crear)
	http.HandleFunc("/insertar", Insertar)
	http.HandleFunc("/editar", Editar)
	http.HandleFunc("/actualizar", Actualizar)
	log.Println("Servidor corriendo...")
	http.ListenAndServe(":8080", nil)
}

// AQUI TRAEMOS EL CONTENIDO DE LA BASE DE DATOS
type Empleado struct {
	Id       int
	Nombre   string
	Apellido string
	Edad     string
}

func Editar(w http.ResponseWriter, r *http.Request) {
	// Obtener el parámetro "id" de la URL
	idEmpleado := r.URL.Query().Get("id")
	fmt.Println(idEmpleado)

	// Establecer conexión a la base de datos
	conexionEstablecida := conexionBD()

	// Consultar el registro del empleado por su ID
	registro, err := conexionEstablecida.Query("SELECT id, nombre, apellido, edad, fecha FROM empleados WHERE id=?", idEmpleado)
	if err != nil {
		panic(err.Error())
	}
	defer registro.Close()

	empleado := Empleado{}

	// Leer el resultado de la consulta y asignar valores al struct Empleado
	for registro.Next() {
		var id int
		var nombre, apellido, edad, fecha string
		err = registro.Scan(&id, &nombre, &apellido, &edad, &fecha)
		if err != nil {
			panic(err.Error())
		}
		empleado.Id = id
		empleado.Nombre = nombre
		empleado.Apellido = apellido
		empleado.Edad = edad
		// Aquí puedes realizar cualquier acción necesaria con la variable 'fecha'
	}

	// Aquí puedes realizar cualquier acción necesaria con el struct Empleado
	// por ejemplo, mostrarlo en una plantilla HTML, modificar sus valores, etc.

	fmt.Println(empleado)
	plantillas.ExecuteTemplate(w, "editar", empleado)

}

// AQUI ES DONDE MOSTRAMOS LA VISTA INICIO, DONDE HACEMOS EL SELECT DE LOS DATOS
func Inicio(w http.ResponseWriter, r *http.Request) {
	conexionEstablecida := conexionBD()

	// Realiza la consulta para obtener los registros de la base de datos ordenados por fecha o ID descendente
	registros, err := conexionEstablecida.Query("SELECT id, nombre, apellido, edad FROM empleados ORDER BY fecha DESC")
	if err != nil {
		panic(err.Error())
	}

	defer registros.Close()
	// Crea una lista para almacenar los registros
	var empleados []Empleado
	// Itera sobre los registros y agrega cada uno a la lista
	for registros.Next() {
		var empleado Empleado
		err := registros.Scan(&empleado.Id, &empleado.Nombre, &empleado.Apellido, &empleado.Edad)
		if err != nil {
			panic(err.Error())
		}
		empleados = append(empleados, empleado)
	}
	// Pasa la lista de empleados a la plantilla
	plantillas.ExecuteTemplate(w, "inicio", empleados)
}
func Crear(w http.ResponseWriter, r *http.Request) {
	plantillas.ExecuteTemplate(w, "crear", nil)
}

// AQUI CREAMOS LA CONEXION A LABASE DE DATOS
func conexionBD() (conexion *sql.DB) {
	Driver := "mysql"
	Usuario := "root"
	Contrasena := ""
	Nombre := "golang"
	conexion, err := sql.Open(Driver, Usuario+":"+Contrasena+"@tcp(127.0.0.1)/"+Nombre)
	if err != nil {
		panic(err.Error())
	}
	return conexion
}

// ESTA FUNCION ES PARA INSERTAR LOS DATOS A LA BASE DE DATOS
func Insertar(w http.ResponseWriter, r *http.Request) {

	//AQUI VALIDAMOS LO QUE TRAMOS DEL FORMULARIO
	if r.Method == "POST" {
		nombre := r.FormValue("nombre")
		apellido := r.FormValue("apellido")
		edad := r.FormValue("edad")

		conexionEstablecida := conexionBD()

		//ESTE ES EL CODIGO SQL DONDE VAMOS A INSERTAR LOS DATOS TRAIDOS DEL FORMULARIO
		insertarRegistros, err := conexionEstablecida.Prepare("INSERT INTO empleados(nombre,apellido,edad) VALUES(?,?,?) ")
		if err != nil {
			panic(err.Error())
		}
		insertarRegistros.Exec(nombre, apellido, edad)
		http.Redirect(w, r, "/", 301)
	}
}

// ESTA ES LA FUNCION ELIMINAR DONDE TRAEMOS EL ID MEDIANTE UN LINK PARA DESPUES PROCESARLO Y ELIMINARLO
func Eliminar(w http.ResponseWriter, r *http.Request) {
	// Obtiene el ID del registro a eliminar desde la URL
	id := r.URL.Query().Get("id")

	conexionEstablecida := conexionBD()

	// Ejecuta la consulta de eliminación utilizando el ID del registro
	_, err := conexionEstablecida.Exec("DELETE FROM empleados WHERE id = ?", id)
	if err != nil {
		panic(err.Error())
	}

	// Redirige de vuelta a la página de inicio después de eliminar el registro
	http.Redirect(w, r, "/", http.StatusFound)
}

func Actualizar(w http.ResponseWriter, r *http.Request) {

	//AQUI VALIDAMOS LO QUE TRAMOS DEL FORMULARIO
	if r.Method == "POST" {
		id := r.FormValue("id")
		nombre := r.FormValue("nombre")
		apellido := r.FormValue("apellido")
		edad := r.FormValue("edad")

		conexionEstablecida := conexionBD()

		//ESTE ES EL CODIGO SQL DONDE VAMOS A INSERTAR LOS DATOS TRAIDOS DEL FORMULARIO
		modificarRegistros, err := conexionEstablecida.Prepare("UPDATE empleados SET nombre=?,apellido=?,edad=? WHERE id=? ")
		if err != nil {
			panic(err.Error())
		}
		modificarRegistros.Exec(nombre, apellido, edad, id)
		http.Redirect(w, r, "/", 301)
	}
}
