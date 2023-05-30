package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	targetURL := "https://jsonplaceholder.typicode.com/users" // Hedef url


	resp, err := http.Get(targetURL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Cevabı okuma
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Cevabı yazdırma
	fmt.Println(string(body))
}
