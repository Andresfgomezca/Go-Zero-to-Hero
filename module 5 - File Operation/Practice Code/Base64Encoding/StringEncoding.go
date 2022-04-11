package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	sample := "Hello@"
	encodedString := base64.StdEncoding.EncodeToString([]byte(sample))
	fmt.Println("Encoded string = ", encodedString)

	originalStringBytes, err := base64.StdEncoding.DecodeString(encodedString)
	if err != nil {
		log.Fatalf("Some error occured during base64 decode. Error %s", err.Error())
	}
	fmt.Println("Decoded string = ", string(originalStringBytes))
}