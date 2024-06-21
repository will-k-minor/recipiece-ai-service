package clients

import (
	"os"
	"net/http"
	"bytes"
	"encoding/json"
)

type ChatGPTClient struct {
    apiKey string
    client *http.Client
	headers map[string]string
	baseUrl string
}

func NewChatGPTClient() *ChatGPTClient {
    apiKey := os.Getenv("CHATGPT_API_KEY")
    if apiKey == "" {
        panic("CHATGPT_API_KEY environment variable is not set")
    }

	baseUrl := "https://api.openai.com/v1/"

	headers := map[string]string{
        "Content-Type": "application/json",
        "Authorization": "Bearer " + apiKey,
		"OpenAI-Beta": "assistants=v2", // think about how to dynamically set this
    }

    return &ChatGPTClient{
        apiKey: apiKey,
        client: &http.Client{},
		headers: headers,
		baseUrl: baseUrl,
    }
}

func (c *ChatGPTClient) newRequest(method, endpoint string, body interface{}) (*http.Request, error) {
    var requestBody []byte
    var err error

    if body != nil {
        requestBody, err = json.Marshal(body)
        if err != nil {
            return nil, err
        }
    }

	url := c.baseUrl + endpoint

    req, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
    if err != nil {
        return nil, err
    }

    for key, value := range c.headers {
        req.Header.Set(key, value)
    }

    return req, nil
}

func (c *ChatGPTClient) CreateThread(message string) (map[string]interface{}, error) {
    requestBody := map[string]interface{}{
        // "model": "gpt-4", // or your preferred model
        "messages": []map[string]string{
            {"role": "user", "content": message},
        },
    }

    req, err := c.newRequest("POST", "/threads", requestBody)
    if err != nil {
        return nil, err
    }

    resp, err := c.client.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var apiResponse map[string]interface{}
    err = json.NewDecoder(resp.Body).Decode(&apiResponse)
    if err != nil {
        return nil, err
    }

    return apiResponse, nil
}