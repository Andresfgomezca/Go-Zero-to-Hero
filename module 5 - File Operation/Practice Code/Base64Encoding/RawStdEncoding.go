package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	sample := "a@"
	encodedStringStdEncoding := base64.StdEncoding.EncodeToString([]byte(sample))
	fmt.Printf("STD Encoding: %s\n", encodedStringStdEncoding)

	encodedStringRawStdEncoding := base64.RawStdEncoding.EncodeToString([]byte(sample))
	fmt.Printf("Raw STD Encoding: %s\n", encodedStringRawStdEncoding)

	originalStringBytes, err := base64.RawStdEncoding.DecodeString(encodedStringRawStdEncoding)
	if err != nil {
		log.Fatalf("Some error occured during base64 decode. Error %s", err.Error())
	}
	fmt.Println(string(originalStringBytes))
}
