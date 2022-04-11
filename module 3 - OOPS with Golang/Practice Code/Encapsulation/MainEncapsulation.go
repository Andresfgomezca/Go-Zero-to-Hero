package main

import (
	"fmt"
)

//unexported , employee package level variable
const maxleaves = 30

//Exported struct, public -available to other package
type Employee struct {
	FirstName   string
	LastName    string
	TotalLeaves int
	LeavesTaken int
}

//Exported function
func (e Employee) LeavesRemaining() {
	fmt.Printf("%s %s has %d leaves remaining\n", e.FirstName, e.LastName, (e.TotalLeaves - e.LeavesTaken))
}

//unexported function-available within employee pkg
func leavesRemaining() {
}

func main() {
	e := Employee{
		FirstName:   "Luke",
		LastName:    "Sky",
		TotalLeaves: 30,
		LeavesTaken: 20,
	}
	e.LeavesRemaining()
}
