package summarizer

import (
    "bytes"
    "encoding/json"
    "io"
    "net/http"
)

type LLMRequest struct {
    Model  string `json:"model"`
    Prompt string `json:"prompt"`
    Stream bool   `json:"stream"`
}

type LLMResponse struct {
    Response string `json:"response"`
    Done     bool   `json:"done"`
}

const (
    LLM_URL = "http://host.docker.internal:11434/api/generate"
)

func GetSummary(prompt string) (string, error) {
    reqBody := LLMRequest{
        Model:  "llama3",
        Prompt: prompt,
        Stream: false,
    }

    jsonData, err := json.Marshal(reqBody)
    if err != nil {
        return "", err
    }

    resp, err := http.Post(LLM_URL, "application/json", bytes.NewBuffer(jsonData))
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }

    var llmResp LLMResponse
    if err := json.Unmarshal(body, &llmResp); err != nil {
        return "", err
    }

    return llmResp.Response, nil
}