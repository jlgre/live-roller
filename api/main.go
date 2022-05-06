package main

import (
<<<<<<< HEAD
	"encoding/json"
	"fmt"
	"net/http"
=======
	"net/http"

	"github.com/josephtanderson/live-roller/api/httpHelpers"
>>>>>>> 63290c656a6068ee60a454db3a949324573883fe
)

type User struct {
	Name string `json:"username"`
	Age  string `json:"age"`
}
<<<<<<< HEAD

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
=======
>>>>>>> 63290c656a6068ee60a454db3a949324573883fe

func main() {

	newUser := User{Name: "Joseph", Age: "30"}

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
<<<<<<< HEAD
		enableCors(&w)
		w.Header().Set("content-Type", "application/json")
		w.Write([]byte(toJson(newUser)))
=======
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("content-Type", "application/json")
		w.Write([]byte(httpHelpers.RespondWithJson(newUser)))
>>>>>>> 63290c656a6068ee60a454db3a949324573883fe
	})
	http.ListenAndServe(":8080", nil)
}
