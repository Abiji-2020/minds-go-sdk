package minds

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func (md *Minds) Completion(model, message string, stream bool) (string, <-chan string, error) {

	requestBody := ChatCompletionRequest{
		Model: model,
		Messages: []Message{
			{
				Role:    "system",
				Content: "You are a helpful assistant",
			},
			{
				Role:    "user",
				Content: message,
			},
		},
		Stream: stream,
	}

	requestData, err := json.Marshal(requestBody)
	if err != nil {
		return "", nil, fmt.Errorf("failed to  marshal request body: %w", err)
	}

	baseURL := strings.TrimSuffix(md.config.BaseURL, "/api")
	chatURL := fmt.Sprintf("%s/chat/completions", baseURL)

	req, err := http.NewRequest("POST", chatURL, bytes.NewBuffer(requestData))
	if err != nil {
		return "", nil, fmt.Errorf("failed to create http request: %w", err)
	}
	req.Header = md.config.Headers()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", nil, fmt.Errorf("http request failed: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	if stream {
		streamChan := make(chan string)
		go func() {
			defer resp.Body.Close()
			defer close(streamChan)

			scanner := bufio.NewScanner(resp.Body)
			for scanner.Scan() {
				line := scanner.Text()
				if strings.HasPrefix(line, "data:") {
					rawData := strings.TrimPrefix(line, "data:")
					rawData = strings.TrimSpace(rawData)

					if rawData == "[DONE]" {
						break
					}

					var completionResponse ChatCompletionResponse
					err := json.Unmarshal([]byte(rawData), &completionResponse)
					if err == nil && len(completionResponse.Choices) > 0 {
						streamChan <- completionResponse.Choices[0].Delta.Content
					}
				}
			}
		}()

		return "", streamChan, nil
	}

	defer resp.Body.Close()
	var completionResponse ChatCompletionResponse
	err = json.NewDecoder(resp.Body).Decode(&completionResponse)
	if err != nil {
		return "", nil, fmt.Errorf("failed to decode response: %w", err)
	}

	if len(completionResponse.Choices) > 0 {
		return completionResponse.Choices[0].Delta.Content, nil, nil
	}

	return "", nil, fmt.Errorf("no choices returned from Minds")
}
