package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
)

func handle(w http.ResponseWriter, r *http.Request) {
	os := runtime.GOOS
	_, err := fmt.Fprintf(w, "Hello World from [%s] container!\n", os)
	if err != nil {
		return
	}
}

func main() {
	http.HandleFunc("/", handle)
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
