package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello from GO!\n")
}

func getJson(w http.ResponseWriter, r *http.Request) {
	p := R{
		Message: "Hello from Go!",
	}

	b, err := json.Marshal(p)
	if err != nil {
		io.WriteString(w, "Boomer")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

type R struct {
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/json", getJson)

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Listening on port 3000")
}
