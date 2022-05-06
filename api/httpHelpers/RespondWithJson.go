package httpHelpers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HttpRespond(endpoint string, response interface{}, host string) {
	http.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", host) //can be removed once Docker is set up
		w.Header().Set("content-Type", "application/json")
		w.Write([]byte(RespondWithJson(response)))
	})
}

func RespondWithJson(response interface{}) string {
	byteArray, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err)
	}
	return string(byteArray)
}
