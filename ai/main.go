package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func ai(b string) string {
	url := ""
	data := map[string]interface{}{
		"prompt":  b,
		"history": []interface{}{},
	}

	payload, err := json.Marshal(data)
	if err != nil {
		fmt.Println("JSON marshaling error:", err)
		//return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Request creation error:", err)
		//return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Request execution error:", err)
		//return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Response reading error:", err)
		//return
	}

	// Print the response body
	var a string
	a = string(body)

	return a
	//fmt.Println(a)
}

func main() {
	var question string
	for {
		fmt.Scan(&question)
		println(ai(question))

	}

}
