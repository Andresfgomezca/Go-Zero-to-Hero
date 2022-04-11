package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	sample := "a@"
	encodedStringURLEncoding := base64.URLEncoding.EncodeToString([]byte(sample))
	fmt.Printf("URL Encoding: %s\n", encodedStringURLEncoding)

	encodedStringRawURLEncoding := base64.RawURLEncoding.EncodeToString([]byte(sample))
	fmt.Printf("Raw URL Encoding: %s\n", encodedStringRawURLEncoding)

	originalStringBytes, err := base64.RawStdEncoding.DecodeString(encodedStringRawURLEncoding)
	if err != nil {
		log.Fatalf("Some error occured during base64 decode. Error %s", err.Error())
	}
	fmt.Println(string(originalStringBytes))
}
