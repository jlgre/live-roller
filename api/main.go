package main

import (
	"net/http"

	"github.com/josephtanderson/live-roller/api/httpHelpers"
)

type User struct {
	Name string `json:"username"`
	Age  string `json:"age"`
}

func main() {

	newUser := User{Name: "Joseph", Age: "30"}

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("content-Type", "application/json")
		w.Write([]byte(httpHelpers.RespondWithJson(newUser)))
	})
	http.ListenAndServe(":8080", nil)
}
