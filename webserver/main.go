package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	catservice "github.com/kimpettersen/introduction-to-go/webserver/pkg/cat"
)

func main() {
	// Use the catservice package to create a new cat
	catHandler := func(w http.ResponseWriter, _ *http.Request) {
		cat := catservice.New()

		// Return the cat
		io.WriteString(w, fmt.Sprintf("%s is %d years old", cat.Name, cat.Age))
	}

	http.HandleFunc("/cat", catHandler)
	fmt.Println("Cat can be found on: http://localhost:8080/cat")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
