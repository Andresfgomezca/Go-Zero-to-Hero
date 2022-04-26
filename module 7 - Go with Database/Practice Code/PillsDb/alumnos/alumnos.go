package alumnos

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

//database constants
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "postgres"
)

type AlumnoRepository interface {
	getAlumnos()
}

type Alumnos struct {
	Id       int
	Nombre   string
	Apellido string
	Edad     string
}

var (
	Db  *sql.DB
	Err error
)

func GetDBConnection() (*sql.DB, error) {
	psqlconn := fmt.Sprintf("host=%s port = %d user=%s password = %s dbname = %s sslmode=disable", host, port, user, password, dbname)

	return sql.Open("postgres", psqlconn)
}
