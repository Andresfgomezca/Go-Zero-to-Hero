package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 54320
	user     = "postgres"
	password = "admin"
	dbname   = "postgres"
)

type User struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Age       int    `json:"age"`
}

func main() {
	r := mux.NewRouter()
	//main path
	usersR := r.PathPrefix("/users").Subrouter()
	//subroutes with its respective method defined as functions
	usersR.Path("/").Methods(http.MethodGet).HandlerFunc(getAllUsersHandler)
	usersR.Path("/").Methods(http.MethodPost).HandlerFunc(createUserHandler)
	usersR.Path("/{id}").Methods(http.MethodGet).HandlerFunc(getUserHandler)
	usersR.Path("/{id}").Methods(http.MethodPut).HandlerFunc(updateUserHandler)
	usersR.Path("/{id}").Methods(http.MethodDelete).HandlerFunc(deleteUserHandler)
	log.Println("Server started listening...")
	log.Println(http.ListenAndServe(":8080", r))
}

//get all users from db
func getAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	users, err := getAllUsers()
	if err != nil {
		log.Println("Error while fetching users:", err)
		http.Error(w, "Error While fetching the users", http.StatusInternalServerError)
		return
	}
	response, err := json.Marshal(&users)
	if err != nil {
		log.Println("Error while marshling the response:", err)
		http.Error(w, "Error while marshling the response", http.StatusInternalServerError)
		return
	}
	w.Write(response)
}

//get user from db with id
func getUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("Invalid user Id:", err)
		http.Error(w, "Invalid user Id", http.StatusBadRequest)
		return
	}
	user, err := getUser(userId)
	if err != nil {
		log.Println("Error while fetching the user details:", err)
		http.Error(w, "Error while fetching the user details", http.StatusInternalServerError)
		return
	}

	if user.Id != userId {
		log.Println("Requested user not found")
		http.Error(w, "Requested user not found", http.StatusNotFound)
		return
	}

	response, err := json.Marshal(user)
	if err != nil {
		log.Println("Error while marshling the user object:", err)
		http.Error(w, "Error while marshling the user object", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

//put method to update an user by id
func updateUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("Invalid user Id:", err)
		http.Error(w, "Invalid user Id", http.StatusBadRequest)
		return
	}

	user := User{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Println(err)
		http.Error(w, "Error decoidng request object", http.StatusBadRequest)
		return
	}

	user.Id = userId
	_, err = updateUser(&user)
	if err != nil {
		log.Println("Error while updating the user:", err)
		http.Error(w, "Error while updating the user", http.StatusInternalServerError)
		return
	}

	res, _ := json.Marshal(user)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//delete method to delete user by id from db
func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("Invalid user Id:", err)
		http.Error(w, "Invalid user Id", http.StatusBadRequest)
		return
	}
	isDeleted, err := deleteUser(userId)
	if err != nil {
		log.Println("Error while deleting the user:", err)
		http.Error(w, "Error while deleting the user", http.StatusInternalServerError)
		return
	}
	if !isDeleted {
		http.Error(w, "UserId not found in the system to delete", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte{})
}

//post method to create user
func createUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	user := User{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Println(err)
		http.Error(w, "Error decoding request object", http.StatusBadRequest)
		return
	}
	u, err := insertUser(&user)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error while inserting the user", http.StatusInternalServerError)
		return
	}
	res, _ := json.Marshal(u)
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

//method to conect to the db
func getDBConnection() (*sql.DB, error) {
	psqlconn := fmt.Sprintf("host=%s port = %d user=%s password = %s dbname = %s sslmode=disable", host, port, user, password, dbname)

	return sql.Open("postgres", psqlconn)
}

func insertUser(user *User) (User, error) {
	db, err := getDBConnection()
	if err != nil {
		log.Printf("Error while connecting to the Database: %s", err)
		return User{}, err
	}
	defer db.Close()

	insertStmt := `insert into users(first_name, last_name, email, age) values($1, $2, $3, $4) returning id`
	lastInsertedId := 0
	rows := db.QueryRow(insertStmt, user.FirstName, user.LastName, user.Email, user.Age).Scan(&lastInsertedId)
	if lastInsertedId == 0 {
		log.Printf("Error while inserting the user: %s", rows.Error())
		return User{}, err
	}

	log.Println("Record Inserted successfully")
	user.Id = lastInsertedId

	return *user, nil
}

//function that gets a slice with the users from the database
func getAllUsers() ([]User, error) {
	db, err := getDBConnection()
	if err != nil {
		log.Printf("Error while connecting to the Database: %s", err)
		return nil, err
	}
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
		if err := rows.Scan(&user.Id, &user.Age, &user.FirstName, &user.LastName, &user.Email); err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	return users, nil
}

//function that getsan user from the db by submiting the id
func getUser(userId int) (User, error) {
	db, err := getDBConnection()
	if err != nil {
		log.Printf("Error while connecting to the Database: %s", err)
		return User{}, err
	}
	defer db.Close()

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

func deleteUser(userId int) (bool, error) {
	db, err := getDBConnection()
	if err != nil {
		log.Printf("Error while connecting to the Database: %s", err)
		return false, err
	}
	defer db.Close()

	deleteUser := `delete from users where id = $1`
	rows, err := db.Exec(deleteUser, userId)
	if err != nil {
		log.Fatalf("Error while deleting user: %s", err)
		return false, err
	}
	cnt, _ := rows.RowsAffected()

	return (cnt > 0), nil
}

//function that updates an user from the database
func updateUser(user *User) (bool, error) {
	db, err := getDBConnection()
	if err != nil {
		log.Printf("Error while connecting to the Database: %s", err)
		return false, err
	}
	defer db.Close()

	updateUserQuery := `update users set first_name=$1, last_name=$2, email=$3, age=$4 where id=$5`
	_, err = db.Exec(updateUserQuery, user.FirstName, user.LastName, user.Email, user.Age, user.Id)
	if err != nil {
		log.Fatalf("Error while updating the user: %s", err)
		return false, nil
	}

	return true, nil
}
