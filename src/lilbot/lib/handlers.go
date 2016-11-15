package lib

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/hybridgroup/gobot"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!\n")
}

func UpdateColor(w http.ResponseWriter, r *http.Request) {
	var rgb []int
	vars := mux.Vars(r)
	var input = strings.Split(vars["rgb"], ",")

	for _, i := range input {
		v, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		rgb = append(rgb, v)
	}
	fmt.Printf("rgb value: %v\n", rgb)

	work := func() {

		gobot.Every(1*time.Second, func() {
			r := uint8(rgb[0])
			g := uint8(rgb[1])
			b := uint8(rgb[2])
			GRobot.driver.SetRGB(r, g, b)
		})

	}

	GRobot.Append(work)

	GRobot.Start()

	w.WriteHeader(http.StatusOK)
}
