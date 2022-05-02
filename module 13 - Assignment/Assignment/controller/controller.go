package controller

import (
	"assignment/user"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func Controller() {
	r := mux.NewRouter()
	//main path
	usersR := r.PathPrefix("/users").Subrouter()
	//subroutes with its respective method defined as functions
	usersR.HandleFunc("/", GetAllUsersHandler).Methods(http.MethodGet)
	usersR.Path("/").Methods(http.MethodPost).HandlerFunc(CreateUserHandler)
	usersR.Path("/{id}").Methods(http.MethodGet).HandlerFunc(GetUserHandler)
	usersR.Path("/email/{email}").Methods(http.MethodGet).HandlerFunc(GetUserEmailHandler)
	usersR.Path("/{id}").Methods(http.MethodPut).HandlerFunc(UpdateUserHandler)
	usersR.Path("/{id}").Methods(http.MethodDelete).HandlerFunc(DeleteUserHandler)
	log.Println("Server started listening...")
	log.Println(http.ListenAndServe(":9090", usersR))

}

//HANDLERS
//get all users from db
func GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	users, err := user.GetAllUsers()
	if err != nil {
		log.Println("Error while fetching users:", err)
		http.Error(w, "Error While fetching the users", http.StatusInternalServerError)
		return
	}
	response, err := json.Marshal(users)
	if err != nil {
		log.Println("Error while marshling the response:", err)
		http.Error(w, "Error while marshling the response", http.StatusInternalServerError)
		return
	}
	w.Write(response)
}

//get user from db with id
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("Invalid user Id:", err)
		http.Error(w, "Invalid user Id", http.StatusBadRequest)
		return
	}
	user, err := user.GetUserById(userId)
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

//get user from db with email
func GetUserEmailHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(r)
	userEmail := vars["email"]
	user, err := user.GetUserByEmail(userEmail)
	if err != nil {
		log.Println("Error while fetching the user details:", err)
		http.Error(w, "Error while fetching the user details", http.StatusInternalServerError)
		return
	}

	if user.Email != userEmail {
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
func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("Invalid user Id:", err)
		http.Error(w, "Invalid user Id", http.StatusBadRequest)
		return
	}

	user1 := user.User{}
	if err := json.NewDecoder(r.Body).Decode(&user1); err != nil {
		log.Println(err)
		http.Error(w, "Error decoidng request object", http.StatusBadRequest)
		return
	}

	user1.Id = userId
	_, err = user.UpdateUser(&user1)
	if err != nil {
		log.Println("Error while updating the user:", err)
		http.Error(w, "Error while updating the user", http.StatusInternalServerError)
		return
	}

	res, _ := json.Marshal(user1)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//delete method to delete user by id from db
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("Invalid user Id:", err)
		http.Error(w, "Invalid user Id", http.StatusBadRequest)
		return
	}
	isDeleted, err := user.DeleteUserById(userId)
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
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	user1 := user.User{}
	if err := json.NewDecoder(r.Body).Decode(&user1); err != nil {
		log.Println(err)
		http.Error(w, "Error decoding request object", http.StatusBadRequest)
		return
	}
	err := user.InsertNewUser(&user1)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error while inserting the user", http.StatusInternalServerError)
		return
	}
	res, _ := json.Marshal(user1)
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}
