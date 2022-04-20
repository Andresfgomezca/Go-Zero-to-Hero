package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://www.globant.com/")
	if err != nil {
		log.Fatalf("Unable to fetch the data: %s", err)
	}

	defer resp.Body.Close()
	fmt.Println("Response Status:", resp.Status)

	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 10; i++ {
		fmt.Println(scanner.Text())
	}
}
