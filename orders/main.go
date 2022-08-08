package main

import (
	"fmt"
	"pkg/utils"
	"net/http"
)

func hello(w http.ResponseWriter, _ *http.Request) {

	_, err := fmt.Fprintf(w, "Orders %d \n", utils.Trigger())
	if err != nil {
		panic(err)
	}
}

func main() {

	http.HandleFunc("/", hello)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
