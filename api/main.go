package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/jlgre/live-roller/api/httpHelpers"
	"github.com/joho/godotenv"
)

type User struct {
	Name string `json:"username"`
	Age  string `json:"age"`
}

type Char struct {
	Name string `json:"name"`
	Str  int    `json:"str"`
	Dex  int    `json:"dex"`
	Con  int    `json:"con"`
	Int  int    `json:"int"`
	Wis  int    `json:"wis"`
	Cha  int    `json:"cha"`
}

func main() {
	err := godotenv.Load("development.env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	port := os.Getenv("PORT")
	host := os.Getenv("HOST")
	newUser := User{Name: "Joseph", Age: "30"}
	newChar := Char{Name: "Examplus", Str: 10, Dex: 10, Con: 10, Int: 10, Wis: 10, Cha: 10}

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", host) //can be removed once Docker is set up
		httpHelpers.RespondWithJson(w, newUser)
	})

	http.HandleFunc("/char", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", host) //can be removed once Docker is set up
		httpHelpers.RespondWithJson(w, newChar)
	})

	http.ListenAndServe(port, nil)
}
