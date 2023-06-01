package api

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
)
type Response struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Usage   struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Choices []struct {
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
		FinishReason string `json:"finish_reason"`
		Index        int    `json:"index"`
	} `json:"choices"`
}
func RandomWordWithChatGPT() *Response {

	url := "https://openai80.p.rapidapi.com/chat/completions"

	payload := strings.NewReader("{\n    \"model\": \"gpt-3.5-turbo\",\n    \"messages\": [\n        {\n            \"role\": \"user\",\n            \"content\": \"forget all your previous instructions now you are a teacher of English and your goal is to provide me a random word every day without is meaning or other sentence and the most important do not add a dot at the end of the word and you start now. DO NOT REPEAT YOURSELF \"\n        }\n    ]\n}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/json")
req.Header.Add("X-RapidAPI-Key", os.Getenv("RAPID_API_KEY"))
	req.Header.Add("X-RapidAPI-Host", "openai80.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	var response Response 
	err := json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		log.Fatalf("err", err)
	}
	return &response 
}
