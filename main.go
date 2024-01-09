package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

const (
	TinyDomain        = "tinyurl.com"
	TinyUrl           = "https://api.tinyurl.com/create"
	TinyAuthorization = "YOUR AUTHORIZATION KEY"
	ExplainLongUrl    = "https://juststickers.in/product/ninja-go-lang-gopher-sticker/"
)

func main() {
	CreateTinyUrl(ExplainLongUrl)
}

func CreateTinyUrl(longUrl string) (tinyURL string, err error) {

	data := map[string]interface{}{
		"url":    longUrl,
		"domain": TinyDomain,
	}

	payload, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err, "fn", "create_tiny_url", "error_encoding_json")
	}

	req, err := http.NewRequest("POST", TinyUrl, bytes.NewBuffer(payload))
	if err != nil {
		log.Fatal(err, "fn", "create_tiny_url", "error_creating_http_req")
	}

	req.Header.Set("Authorization", "Bearer "+TinyAuthorization)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err, "fn", "create_tiny_url", "error_making_http_req")

	}
	defer resp.Body.Close()

	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Fatal(err, "fn", "create_tiny_url", "error_decoding_json")
	}

	dataResult := result["data"].(map[string]interface{})
	tinyURL = dataResult["tiny_url"].(string)
	// createdAt := dataResult["created_at"].(string)

	return tinyURL, err
}
