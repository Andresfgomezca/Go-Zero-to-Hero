package alumnos

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

//database constants
const (
	host = "localhost"
	//port     = 5432
	//user     = "postgres"
	//password = "admin"
	dbname = "postgres"
)

type Alumno struct {
	Id       int
	Nombre   string
	Apellido string
	Edad     string
}

var (
	Db       *sql.DB
	Err      error
	User     string
	Password string
	Port     string
)

func GetDBConnection() (*sql.DB, error) {
	psqlconn := fmt.Sprintf("host=%s port = %s user=%s password = %s dbname = %s sslmode=disable", host, Port, User, Password, dbname)
	fmt.Println("Database connected")
	return sql.Open("postgres", psqlconn)
}

//needs uppercase to be exported
func GetAlumnos() ([]Alumno, error) {
	fetchAllUsers := `select * from alumnos`
	rows, err := Db.Query(fetchAllUsers)
	if err != nil {
		log.Fatalf("Error while fetching all users: %s", err)
		return nil, err
	}
	var Alumnos []Alumno
	defer rows.Close()
	for rows.Next() {
		var Alumno Alumno
		if err := rows.Scan(&Alumno.Id, &Alumno.Nombre, &Alumno.Apellido, &Alumno.Edad); err != nil {
			log.Fatal(err)
		}
		Alumnos = append(Alumnos, Alumno)
	}
	return Alumnos, nil
}

func CargarAlumnos(a *Alumno) error {
	//query implemented in the db
	query := `insert into alumnos(nombre, apellido, edad) values ($1,$2,$3)`
	//execution of the query in the db
	_, e := Db.Exec(query, a.Nombre, a.Apellido, a.Edad)
	//verification of the erro that can be produced adding the new element
	if e != nil {
		log.Fatal("ERROR ADDING THE NEW USER", e)
		return e
	}
	fmt.Println("REGISTRATION COMPLETE")
	return nil
}
