package burner

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

//Send sends a message to phone number
//See more at: https://developer.burnerapp.com/api-documentation/api/
func Send(burnerID, toNumber, text, mediaURL string) error {
	if AuthToken == "" {
		return errors.New("Invalid AuthToken")
	}
	requestMsg := url.Values{}
	requestMsg.Set("burnerId", burnerID)
	requestMsg.Set("toNumber", toNumber)
	requestMsg.Set("text", text)
	requestMsg.Set("mediaURL", mediaURL)
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/messages", baseURL), strings.NewReader(requestMsg.Encode()))
	if err != nil {
		return fmt.Errorf("failed to create request. Error: %s", err.Error())
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

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
