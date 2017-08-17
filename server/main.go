package main

import (
	"log"
	"net/http"
)

func main() {
	http.ListenAndServe(":8081", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Header)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("hello!"))
	}))
}
