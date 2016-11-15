package main

import (
	"log"
	"net/http"
)

var MyRobots Robots

func Main() {

	MyRobots = NewRobots()

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
