package main

import (
	"alumnos/alumnos"
	"fmt"
	"log"
	"os"
)

type AlumnoRepository interface {
	getAlumnos()
	cargarAlumnos()
}
type AlumnoRepositoryPostgres struct {
	alumnos []alumnos.Alumno
	alumno  alumnos.Alumno
}

func checkError(e error) {
	if e != nil {
		log.Fatalf("An error happend: %v", e)
	}
}

func main() {
	os.Setenv("port", "5432")
	os.Setenv("user", "postgres")
	os.Setenv("password", "admin")
	//variables that needs to be read from the environment
	alumnos.Port = os.Getenv("port")         //5432
	alumnos.User = os.Getenv("user")         //"postgres"
	alumnos.Password = os.Getenv("password") //"admin"
	alumnos.Db, alumnos.Err = alumnos.GetDBConnection()

	defer alumnos.Db.Close()

	if alumnos.Err != nil {
		log.Fatalf("An error happend: %v", alumnos.Err)
	}
	fmt.Println("getting all elements")
	alumnosActuales := AlumnoRepositoryPostgres{}
	alumnosActuales.getAlumnos()
	fmt.Println(alumnosActuales.alumnos)

	fmt.Println("adding new element...")
	newStudent := alumnos.Alumno{Nombre: "pepito2", Apellido: "Perez2", Edad: "15"}

	alumnosActuales.alumno = newStudent
	cargarAlumnos(alumnosActuales)

	fmt.Println("getting all elements")
	alumnosActuales.getAlumnos()
	fmt.Println(alumnosActuales.alumnos)

	fmt.Println("finish...")

}

func (alumnosSlice *AlumnoRepositoryPostgres) getAlumnos() {
	var e error
	alumnosSlice.alumnos, e = alumnos.GetAlumnos()
	checkError(e)
}

func cargarAlumnos(a AlumnoRepositoryPostgres) {
	var e error
	e = alumnos.CargarAlumnos(&a.alumno)
	checkError(e)
}
