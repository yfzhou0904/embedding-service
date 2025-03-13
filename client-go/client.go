package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// TextEmbeddingRequest represents the request payload for text embedding.
type TextEmbeddingRequest struct {
	Text string `json:"text"`
}

// TextEmbeddingResponse represents the response from the server.
type TextEmbeddingResponse struct {
	Embedding []float32 `json:"embedding"`
}

// GetTextEmbedding sends a text to the server and returns the embedding.
func GetTextEmbedding(serverURL, text string) ([]float32, error) {
	requestBody, err := json.Marshal(TextEmbeddingRequest{Text: text})
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(serverURL+"/embedding/text", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("server returned non-200 status code")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response TextEmbeddingResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return response.Embedding, nil
}
