package main

import (
	"fmt"
	"net/http"
)

func HandleGeo(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.RawQuery)
}

func main() {
	http.HandleFunc("/geo", HandleGeo)
	fmt.Println("Starting...")
	log.Fatal(http.ListenAndServe(":5050", nil))
}
