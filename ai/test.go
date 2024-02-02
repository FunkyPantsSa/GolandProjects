package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	url := "https://api.link-ai.chat/v1/chat/completions"

	payload := strings.NewReader("{\n  \"app_code\": \"default\",\n  \"messages\": [\n    {\n      \"role\": \"user\",\n      \"content\": \"你好\"\n    }\n  ]\n}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer undefined")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
