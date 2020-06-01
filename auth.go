package burner

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

//GetAuthorizeEndpoint return url string to authorize Burner OAuth2
//This function should be used in a web server in order to receive the auth callback
//See more at: https://developer.burnerapp.com/api-documentation/authentication/
func GetAuthorizeEndpoint(clientID, redirectURL string) string {
	return fmt.Sprintf("https://app.burnerapp.com/oauth/authorize?client_id=%s&scope=messages:connect burners:read burners:write contacts:read contacts:write&redirect_uri=%s",
		clientID, redirectURL)
}

//AccessResponse response from https://api.burnerapp.com/oauth/access
type AccessResponse struct {
	AccessToken      string            `json:"access_token"`
	ExpiresIn        int64             `json:"expires_in"`
	Scope            string            `json:"scope"`
	TokenType        string            `json:"token_type"`
	ConnectedBurners []ConnectedBurner `json:"connected_burners"`
}

//ConnectedBurner contains a list of burner associated with the account
type ConnectedBurner struct {
	ID                 string `json:"id"`
	Name               string `json:"name"`
	Status             string `json:"status"`
	OutgoingWebhookURL string `json:"outgoing_webhook_url"`
}

//HandleAuthCallback process burner's callback request
//code will present in the query string from Burner's callback request.
//redirectURL will need to be exact match with the request.
//See more at: https://developer.burnerapp.com/api-documentation/authentication/
func HandleAuthCallback(code, clientID, clientSecret, redirectURL string) ([]ConnectedBurner, error) {
	requestMsg := url.Values{}
	requestMsg.Set("client_id", clientID)
	requestMsg.Set("client_secret", clientSecret)
	requestMsg.Set("code", code)
	requestMsg.Set("grant_type", "authorization_code")
	requestMsg.Set("redirect_uri", redirectURL)
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/oauth/access", baseURL), strings.NewReader(requestMsg.Encode()))
	if err != nil {
		return nil, fmt.Errorf("failed to create request. Error: %s", err.Error())
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json")

	resp, err := Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send auth request to Burner: %s", err.Error())
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get Burner auth token. Burner returned: %s", string(bodyBytes))
	}
	respBody := &AccessResponse{}
	err = json.Unmarshal(bodyBytes, &respBody)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal AccessResponse response. Error: %s", err.Error())
	}
	AuthToken = respBody.AccessToken
	return respBody.ConnectedBurners, nil
}
