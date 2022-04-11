package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	sample := "�hello"

	encodedStringURL := base64.URLEncoding.EncodeToString([]byte(sample))
	fmt.Printf("URL Encoding: %s\n", encodedStringURL)
	encodedStringSTD := base64.StdEncoding.EncodeToString([]byte(sample))
	fmt.Printf("STD Encoding: %s\n", encodedStringSTD)

	originalStringBytes, err := base64.URLEncoding.DecodeString(encodedStringURL)
	if err != nil {
		log.Fatalf("Some error occured during base64 decode. Error %s", err.Error())
	}
	fmt.Println(string(originalStringBytes))
}