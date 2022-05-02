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
	pguser   = "postgres"
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
	db  *sql.DB
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
	psqlconn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", pguser, password, host, port, dbname)

	return sql.Open("postgres", psqlconn)
}

func InsertNewUser(user *User) error {
	db, _ = GetDBConnection()
	defer db.Close()
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

	_, e := db.Exec(query,
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

func DeleteUserById(userId int) (bool, error) {
	db, _ = GetDBConnection()
	defer db.Close()
	query := `delete from users where id = $1`

	_, e := db.Query(query, userId)

	if e != nil {
		log.Fatal("ERROR DELETING USER", e)

		return false, e
	}

	fmt.Println("RECORD DELETED SUCCESFULY")

	return true, nil
}

func GetUserByEmail(email string) (User, error) {
	db, _ = GetDBConnection()
	defer db.Close()
	query := `select * from users where email = $1`
	//verification of the email
	_, ErrorEmail := mail.ParseAddress(email)

	if ErrorEmail != nil {
		return User{}, ErrorEmail
	}

	rows, e := db.Query(query, email)

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

func GetUserById(id int) (User, error) {
	db, _ = GetDBConnection()
	defer db.Close()
	query := `select * from users where id = $1`

	rows, e := db.Query(query, id)

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
func GetAllUsers() ([]User, error) {
	db, _ = GetDBConnection()
	defer db.Close()
	fetchAllUsers := `select * from users`
	rows, err := db.Query(fetchAllUsers)
	if err != nil {
		log.Fatalf("Error while fetching all users: %s", err)
		return nil, err
	}
	var users []User
	defer rows.Close()
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id,
			&user.Age,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.ZipCode,
			&user.City,
			&user.AddressU,
			&user.StateU,
			&user.PasswordU); err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	return users, nil
}

func UpdateUser(user *User) (bool, error) {
	db, _ = GetDBConnection()
	defer db.Close()
	updateUserQuery := `update users set first_name=$2, last_name=$3, zipcode=$4, city=$5, addressu=$6, stateu=$7, age=$8 where id=$1`
	_, err := db.Exec(updateUserQuery,
		user.Id,
		user.FirstName,
		user.LastName,
		user.ZipCode,
		user.City,
		user.AddressU,
		user.StateU,
		user.Age)

	if err != nil {
		log.Fatalf("Error while updating the user: %s", err)
		return false, nil
	}

	return true, nil
}
