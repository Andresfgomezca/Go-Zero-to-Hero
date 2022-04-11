package main
import "fmt"
func main(){
	a:=[]string{}
	//filling slice
	for i:=0; i <= 9; i++{
		a = append(a,"item")
	}
	//iterating over slice a to print each values
	for i,v := range a{
		fmt.Println("index ",i, " ",v)
	}
}