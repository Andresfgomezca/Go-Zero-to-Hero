package user

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/mail"
	"regexp"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "postgres"
)

type User struct {
	Id        int
	Age       int
	FirstName string
	LastName  string
	Email     string
	ZipCode   string
	City      string
	AddressU  string
	StateU    string
	PasswordU string
}

var (
	Db  *sql.DB
	Err error
)

//Custom Errors

type ZipCodeError struct {
	e            string
	zipCodeValue string
}

func (z *ZipCodeError) Error() string {
	return z.e
}

func GetDBConnection() (*sql.DB, error) {
	psqlconn := fmt.Sprintf("host=%s port = %d user=%s password = %s dbname = %s sslmode=disable", host, port, user, password, dbname)

	return sql.Open("postgres", psqlconn)
}

func InsertNewUser(user *User) error {

	matchZipCode, _ := regexp.MatchString("([0-9]{6})", user.ZipCode)

	if !(matchZipCode) {
		return &ZipCodeError{"Invalid Zipcode", user.ZipCode}
	}

	_, ErrorEmail := mail.ParseAddress(user.Email)

	if ErrorEmail != nil {
		return ErrorEmail
	}

	query := `insert into users(age, first_name, last_name,
		 email, zipcode, city, addressU, stateU, passwordU) values ($1,$2,$3,$4,$5,$6,$7,$8,$9)`

	_, e := Db.Exec(query,
		user.Age,
		user.FirstName,
		user.LastName,
		user.Email,
		user.ZipCode,
		user.City,
		user.AddressU,
		user.StateU,
		user.PasswordU)

	if e != nil {
		log.Fatal("ERROR ADDING THE NEW USER", e)
		return e
	}

	fmt.Println("USER REGISTERED SUCCESSFUL")
	return nil
}

func DeleteUserById(userId int) error {

	query := `delete from users where id = $1`

	_, e := Db.Query(query, Err)

	if e != nil {
		log.Fatal("ERROR DELETING USER", e)

		return e
	}

	fmt.Println("RECORD DELETED SUCCESFULY")

	return nil
}

func GetUserByEmail(email string) (User, error) {
	query := `select * from users where email = $1`
	//verification of the email
	_, ErrorEmail := mail.ParseAddress(email)

	if ErrorEmail != nil {
		return User{}, ErrorEmail
	}

	rows, e := Db.Query(query, email)

	if e != nil {
		log.Fatal("ERROR FINDING USER", e)
		return User{}, e
	}

	var u User

	if rows.Next() {
		if err := rows.Scan(&u.Id,
			&u.Age,
			&u.FirstName,
			&u.LastName,
			&u.Email,
			&u.ZipCode,
			&u.City,
			&u.AddressU,
			&u.StateU,
			&u.PasswordU); err != nil {
			log.Fatal(e)
		}

	}
	if u.Id == 0 {
		log.Fatal("ERROR FINDING USER", errors.New(" EMAIL DOESN'T EXISTS"))
		return User{}, errors.New("EMAIL DOESN'T EXISTS")
	}

	return u, nil
}
