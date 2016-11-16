package main

import (
	"log"
	"net/http"
	"lilbot/lib"
)

var MyRobots *lib.LilRobots

func main() {

	MyRobots = lib.NewGRobots()

	router := lib.NewRouter(MyRobots)

	log.Fatal(http.ListenAndServe(":8080", router))
}
