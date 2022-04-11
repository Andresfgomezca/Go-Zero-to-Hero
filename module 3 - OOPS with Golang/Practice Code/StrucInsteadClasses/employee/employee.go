package employee
 
import (
   "fmt"
)
 
 
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