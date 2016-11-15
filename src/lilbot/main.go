package main

import (
	"log"
	"net/http"
	"lilbot/lib"
)

var MyRobots *lib.Robots

func main() {

	MyRobots = lib.NewRobots()

	router := lib.NewRouter(MyRobots)

	log.Fatal(http.ListenAndServe(":8080", router))
}
