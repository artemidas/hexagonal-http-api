package main

import (
	"fmt"
	"log"
	"net/http"
)

const httpAddr = ":8080"

func main() {
	fmt.Println("Server running on", httpAddr)

	mux := http.NewServeMux()

	log.Fatal(http.ListenAndServe(httpAddr, mux))
}
