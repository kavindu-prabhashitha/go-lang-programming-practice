package main

import "go-tutorial/300-oop-in-go/employee"

func main() {
	e := employee.Employee{
		FirstName:   "Sam",
		LastName:    "Adolf",
		TotalLeaves: 30,
		LeavesTaken: 20,
	}
	e.LeavesRemaining()
}
