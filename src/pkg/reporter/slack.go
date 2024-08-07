package reporter

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/ethereum/go-ethereum/log"
	"github.com/spf13/viper"
)

type SlackMessage struct {
	Attachments []Attachment `json:"attachments"`
}

type Attachment struct {
	Color    string `json:"color"`
	Title    string `json:"title"`
	Text     string `json:"text"`
	Fallback string `json:"fallback"`
}

var webhookURL string

func init() {
	webhookURL = viper.GetString("slack.webhook_url")
}

func Report(title string, text string) {
	attachment := Attachment{
		Color:    "good",
		Title:    title,
		Text:     text,
		Fallback: text,
	}

	message := SlackMessage{
		Attachments: []Attachment{attachment},
	}

	jsonBytes, err := json.Marshal(message)
	if err != nil {
		log.Error("Failed to marshal JSON", "error", err)
		return
	}

	req, err := http.NewRequest("POST", webhookURL, bytes.NewBuffer(jsonBytes))
	if err != nil {
		log.Error("Failed to create request", "error", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Error("Failed to send request", "error", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error("Failed to read response body", "error", err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		log.Error("slack API error: %s", string(body))
		return
	}
}
