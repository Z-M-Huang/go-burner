package burner

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type incomingWebhookRequest struct {
	Intent string                      `json:"intent"`
	Data   *incomingWebhookRequestData `json:"data"`
}

type incomingWebhookRequestData struct {
	ToNumber string `json:"toNumber"`
	Text     string `json:"text"`
}

//SendIncomingWebhook sends a message to phone number using incoming webhook
//See more at: https://developer.burnerapp.com/api-documentation/incoming-webhooks/
func (c *Client) SendIncomingWebhook(burnerID, toNumber, text string) error {
	_, err := url.ParseRequestURI(c.IncomingWebhookURL)
	if err != nil {
		return errors.New("Empty Incoming Webhook URL")
	}
	requestBody, err := json.Marshal(&incomingWebhookRequest{
		Intent: "message",
		Data: &incomingWebhookRequestData{
			ToNumber: toNumber,
			Text:     text,
		},
	})
	if err != nil {
		return fmt.Errorf("failed to parse request. Error: %s", err.Error())
	}
	req, err := http.NewRequest("POST", c.IncomingWebhookURL, bytes.NewBuffer(requestBody))
	if err != nil {
		return fmt.Errorf("failed to create request. Error: %s", err.Error())
	}
	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	resp, err := HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request to Burner: %s", err.Error())
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("failed to send message via webhook. Burner returned: %s", string(bodyBytes))
	}
	return nil
}
