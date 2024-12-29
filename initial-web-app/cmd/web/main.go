package main

import (
	"fmt"
	"log"
	"net/http"
)

// servidor web q escucha peticion
const port = ":4000"

type application struct{}

func main() {
	// app := application{}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})

	fmt.Println("Listening on ", port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Panicln(err)
	}
}
