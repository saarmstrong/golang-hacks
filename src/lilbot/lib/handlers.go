package lib

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!\n")

	work := func() {
        // green means go!
        GRobot.driver.SetRGB(uint8(0), uint8(255), uint8(0))
	}

	GRobot.Add(work)
	GRobot.Start()
}

func UpdateColor(w http.ResponseWriter, r *http.Request) {
	var rgb []int
	vars := mux.Vars(r)
	var input = strings.Split(vars["rgb"], ",")

    // TODO find a better way to get this param
	for _, i := range input {
		v, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		rgb = append(rgb, v)
	}

	fmt.Printf("rgb value: %v\n", rgb)

    GRobot.SetRGB(rgb)

	w.WriteHeader(http.StatusOK)
}
