package api

import (
	"log"
	"fmt"
	"net/http"
	"github.com/julienschmidt/httprouter"
)


func Init() {
	// Setting up handlers
	go InitConn(GetRef("DirObjDB", "5432"))


	r := httprouter.New()
	r.GET("/", HandleIndex)

	
	// Handle User Requests
	r.GET("/users", HandleAllusers)
	r.GET("/users/:id", HandleUserId)


	fmt.Println("Microservice listening on port 8888")
	log.Fatal(http.ListenAndServe(":8888", r))
}