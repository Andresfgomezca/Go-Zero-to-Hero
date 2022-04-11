package main

import "employee"

func main() {
	e := employee.Employee{
		FirstName:   "Luke",
		LastName:    "Sky",
		TotalLeaves: 30,
		LeavesTaken: 20,
	}
	e.LeavesRemaining()
}
