package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type User struct {
	Name string `json:"username"`
	Age  string `json:"age"`
}

func toJson(x User) string {
	byteArray, err := json.Marshal(x)
	if err != nil {
		fmt.Println(err)
	}
	return string(byteArray)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
}

func main() {
	err := godotenv.Load("development.env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	port := os.Getenv("PORT")
	newUser := User{Name: "Joseph", Age: "30"}

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		w.Header().Set("content-Type", "application/json")
		w.Write([]byte(toJson(newUser)))
	})
	http.ListenAndServe(port, nil)
}
