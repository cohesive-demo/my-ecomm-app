package main

import (
	"fmt"
	"log"
	"github.com/testOrgHimali/app/pkg/utils"
	"net/http"
)

func hello(w http.ResponseWriter, _ *http.Request) {
	log.Printf("Logging")
log.Printf("Second line\n")
	_, err := fmt.Fprintf(w, "Hello Orders %d \n", utils.Trigger())
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
