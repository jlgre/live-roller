package httpHelpers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func RespondWithJson(w http.ResponseWriter, response interface{}) {
	w.Header().Set("content-Type", "application/json")
	byteArray, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err)
	}
	w.Write([]byte(byteArray))
	return
}
