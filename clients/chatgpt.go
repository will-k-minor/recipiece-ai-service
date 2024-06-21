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

func (c *ChatGPTClient) makeRequest(method, endpoint string, body interface{}) (map[string]interface{}, error) {
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

func (c *ChatGPTClient) CreateThread(message string) (map[string]interface{}, error) {
	requestBody := map[string]interface{}{
		"messages": []map[string]string{
			{"role": "user", "content": message},
		},
	}
	return c.makeRequest("POST", "/threads", requestBody)
}

func (c *ChatGPTClient) CreateMessage(threadId, message string) (map[string]interface{}, error) {
	requestBody := map[string]interface{}{
		"role":    "user",
		"content": message,
	}
	endpoint := "/threads/" + threadId + "/messages"
	return c.makeRequest("POST", endpoint, requestBody)
}

func (c *ChatGPTClient) RunAssistant(threadId string) (map[string]interface{}, error) {
	assistantId := os.Getenv("CLAUDE_ASSISTANT_ID")
	requestBody := map[string]interface{}{
		"assistant_id": assistantId,
		"response_format": map[string]string{"type": "json_object"},
		"additional_messages": []map[string]interface{}{
			{
				"role":    "user",
				"content": "Give me a response in JSON",
			},
		},
	}
	endpoint := "/threads/" + threadId + "/runs"
	return c.makeRequest("POST", endpoint, requestBody)
}

func (c *ChatGPTClient) GetMessagesFromThread(threadId string) (map[string]interface{}, error) {
	endpoint := "/threads/" + threadId + "/messages"
	return c.makeRequest("GET", endpoint, nil)
}

func (c *ChatGPTClient) CancelRun(threadId, runId string) (map[string]interface{}, error) {
	endpoint := "/threads/" + threadId + "/runs/" + runId + "/cancel"
	return c.makeRequest("POST", endpoint, nil)
}

func (c *ChatGPTClient) GetRunDetails(threadId, runId string) (map[string]interface{}, error) {
	endpoint := "/threads/" + threadId + "/runs/" + runId
	return c.makeRequest("GET", endpoint, nil)
}

func (c *ChatGPTClient) SubmitToolOutput(threadId, runId, id, output string) (map[string]interface{}, error) {
	requestBody := map[string]interface{}{
		"tool_outputs": []map[string]string{
			{"tool_call_id": id, "output": output},
		},
	}
	endpoint := "/threads/" + threadId + "/runs/" + runId + "/submit_tool_outputs"
	return c.makeRequest("POST", endpoint, requestBody)
}