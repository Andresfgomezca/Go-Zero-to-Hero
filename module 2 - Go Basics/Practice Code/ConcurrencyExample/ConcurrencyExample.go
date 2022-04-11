package main

import (
	"fmt"
	"time"
)

func process(i int) {
	fmt.Println("Procesando", i) // n segundos
	time.Sleep(3 * time.Second)
}

//el procesador ejecutara el programa en 3*n segundos
func main() {
	list := []int{1, 2, 3}
	for _, v := range list {
		//esta keyword will run the program in a different core
		//esta es una herramienta para aplicar paralelismo
		go process(v)
	}
	//proceso los 3 a la vez y los termino a la vez
	time.Sleep(4 * time.Second)
}
