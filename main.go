package main

import (
	"fmt"
	"net/http"
	"os"
)

func sayTheAnswer(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("42"))
}

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}

	http.HandleFunc("/", sayTheAnswer)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		panic(err)
	}
}
