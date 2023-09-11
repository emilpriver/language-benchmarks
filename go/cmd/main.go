package main

import (
	"fmt"
	"io"
	"net/http"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello from GO!\n")
}

func main() {
	http.HandleFunc("/", getRoot)

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Listening on port 3000")
}
