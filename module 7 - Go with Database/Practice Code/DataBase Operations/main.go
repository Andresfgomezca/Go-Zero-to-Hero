package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 54320
	user     = "postgres"
	password = "my_password"
	dbname   = "postgres"
)

type User struct {
	Id        int
	FirstName string
	LastName  string
	Email     string
	Age       int
}

var db *sql.DB

func main() {
	var err error
	db, err = getDBConnection()
	if err != nil {
		log.Fatalf("DB connection error: %s", err)
	}
	defer db.Close()

	user := &User{
		FirstName: "Felicity",
		LastName:  "Smoke",
		Email:     "felicity@example.com",
		Age:       20,
	}
	e := InsertUser(user)
	if e != nil {
		log.Fatalf("Error while inserting the record.")
	}

	users, err := getAllUsers()
	if err != nil {
		log.Fatalf("Error while fetching users data: %s", err)
	}
	fmt.Println("Users Data", users)

	user2, err2 := getUser(3)
	if err2 != nil {
		log.Fatalf("Error while fetching user: %s", err)
	}
	fmt.Println("User Details: ", user2)

	err = deleteUser(5)
	if err != nil {
		log.Fatalf("Error while deleting the user: %s", err)
	}
	fmt.Println("User Deleted Successfully.")

	user1 := &User{
		FirstName: "Felicityy",
		LastName:  "Smokee",
		Email:     "smoke@example.com",
		Age:       23,
	}

	isUpdated, err := updateUser(user1)
	if err != nil {
		log.Fatalf("Error while updating the user: %s", err)
	}
	if isUpdated {
		fmt.Println("User updated successfully. User Details:", user1)
	}

}

func getDBConnection() (*sql.DB, error) {
	psqlconn := fmt.Sprintf("host=%s port = %d user=%s password = %s dbname = %s sslmode=disable", host, port, user, password, dbname)

	return sql.Open("postgres", psqlconn)
}

func InsertUser(user *User) error {
	insertStmt := `insert into users(first_name, last_name, email, age) values($1, $2, $3, $4)`

	_, e := db.Exec(insertStmt, user.FirstName, user.LastName, user.Email, user.Age)
	if e != nil {
		log.Fatalf("Error while inserting into Emp table: %s", e)
		return e
	}
	fmt.Println("Record Inserted successfully")

	return nil
}

func getAllUsers() ([]User, error) {
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
		if err := rows.Scan(&user.Id, &user.Age, &user.FirstName, &user.LastName, &user.Email); err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	return users, nil
}

func getUser(userId int) (User, error) {
	getUserSql := `select * from users where id = $1`
	rows, err := db.Query(getUserSql, userId)
	if err != nil {
		return User{}, err
	}
	var user User
	if rows.Next() {
		if err = rows.Scan(&user.Id, &user.Age, &user.FirstName, &user.LastName, &user.Email); err != nil {
			log.Fatal(err)
		}
	}
	return user, nil
}

func deleteUser(userId int) error {
	deleteUser := `delete from users where id = $1`
	_, err := db.Query(deleteUser, userId)
	if err != nil {
		log.Fatalf("Error while deleting user: %s", err)
		return err
	}

	return nil
}

func updateUser(user *User) (bool, error) {
	updateUserQuery := `update users set first_name=$1, last_name=$2, email=$3, age=$4 where id=$5`
	_, err := db.Exec(updateUserQuery, user.FirstName, user.LastName, user.Email, user.Age, user.Id)
	if err != nil {
		log.Fatalf("Error while updating the user: %s", err)
		return false, nil
	}

	return true, nil
}