package burner

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type sendMessageRequest struct {
	BurnerID string `json:"burnerId"`
	ToNumber string `json:"toNumber"`
	Text     string `json:"text"`
	MediaURL string `json:"mediaUrl,omitempty"`
}

//Send sends a message to phone number
//See more at: https://developer.burnerapp.com/api-documentation/api/
func Send(burnerID, toNumber, text, mediaURL string) error {
	if AuthToken == "" {
		return errors.New("Invalid AuthToken")
	}
	requestBody, err := json.Marshal(&sendMessageRequest{
		BurnerID: burnerID,
		ToNumber: toNumber,
		Text:     text,
		MediaURL: mediaURL,
	})
	if err != nil {
		return fmt.Errorf("failed to parse request. Error: %s", err.Error())
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/messages/", baseURL), bytes.NewBuffer(requestBody))
	if err != nil {
		return fmt.Errorf("failed to create request. Error: %s", err.Error())
	}
	req.Header.Add("Content-Type", "application/json")
	setAuthHeader(req)

	resp, err := Client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request to Burner: %s", err.Error())
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("failed to send message. Burner returned: %s", string(bodyBytes))
	}
	return nil
}
