package main

import (
	"fmt"
	"log"
	"postgres-demo/user"
)

func checkError(e error) {
	if e != nil {
		log.Fatalf("An error happend: %v", e)
	}
}

func main() {

	user.Db, user.Err = user.GetDBConnection()

	defer user.Db.Close()

	if user.Err != nil {
		log.Fatalf("An error happend: %v", user.Err)
	}

	var requestUser user.User

	requestUser, user.Err = user.GetUserByEmail("Jbond13@globant.com")
	checkError(user.Err)
	fmt.Println(requestUser)

}
