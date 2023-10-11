package main

import (
	"encoding/json"
	"github.com/yalp/jsonpath/api2d"
	"log"
	"net/http"
)

type RequestBody struct {
	Variable string `json:"variable"`
}
type ResponseBody struct {
	Result string `json:"result"`
}

func main() {
	http.HandleFunc("/", handleRequest)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func handleRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	decoder := json.NewDecoder(r.Body)
	//fmt.Println(decoder)
	var requestBody RequestBody
	err := decoder.Decode(&requestBody)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	//responseBody := ResponseBody{
	//	Result: requestBody.Variable,
	//}
	var token = `fk196722`
	var answer = requestBody.Variable
	var response = api2d.Api2d(token, answer)
	byteData := []byte(response)

	//response, err := json.Marshal(responseBody)
	//if err != nil {
	//	http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	//	return
	//}

	w.Header().Set("Content-Type", "application/json")
	w.Write(byteData)

}
