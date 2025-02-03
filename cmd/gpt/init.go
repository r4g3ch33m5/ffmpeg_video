package gpt

import (
	"encoding/json"
	"os"
)

var apiKey string

func init() {
	val, err := os.ReadFile("./credential/chatgpt.json")
	if err != nil {
		panic(err)
	}
	value := make(map[string]string)
	json.Unmarshal(val, &value)
	apiKey = value["api_key"]
}
