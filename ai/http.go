package main

import (
	"encoding/json"
	"fmt"
	"github.com/yalp/jsonpath/linkai"
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
	log.Fatal(http.ListenAndServe(":8081", nil))
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

	//var token = `Link_RhEfGRxJi9Tvq1XKbEV8c2yvil4nbaDpwhII4FxecE`
	//var token = `api2d_token`
	var answer = requestBody.Variable
	var token = requestBody.token
	var response = linkai.LinkAiJsonToData(linkai.Linkai(token, answer))

	//var response = api2d.Api2d(token, answer)
	fmt.Println(response)
	byteData := []byte(response)

	//response, err := json.Marshal(responseBody)
	//if err != nil {
	//	http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	//	return
	//}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Write(byteData)

}
