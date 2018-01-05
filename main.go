package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {

	fmt.Println("Building router..")

	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)

	fmt.Println("Listening and Serving..")
	log.Fatal(http.ListenAndServe(":8081", router))
}