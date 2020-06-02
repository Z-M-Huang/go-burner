package burner

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

//GetBurners get a list of connected burners.
//See more at: https://developer.burnerapp.com/api-documentation/api/
func GetBurners() ([]ConnectedBurner, error) {
	if AuthToken == "" {
		return nil, errors.New("Invalid AuthToken")
	}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/burners/", baseURL), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create http request. Error: %s", err.Error())
	}
	setAuthHeader(req)
	resp, err := Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request to Burner: %s", err.Error())
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read body content. Error: %s", err.Error())
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get burners. Burner returned: %s", string(bodyBytes))
	}
	var respBody []ConnectedBurner
	err = json.Unmarshal(bodyBytes, &respBody)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body. Error: %s", err.Error())
	}
	return respBody, nil
}

//UpdateBurner update burner settings
//See more at: https://developer.burnerapp.com/api-documentation/api/
func UpdateBurner(burnerID, name string, ringer, notifications, autoReplyActive, autoReplyText, callerIDEnabled bool, color string) ([]ConnectedBurner, error) {
	if AuthToken == "" {
		return nil, errors.New("Invalid AuthToken")
	}
	baseURL, _ := url.Parse(fmt.Sprintf("%s/burners/%s/", baseURL, burnerID))
	params := url.Values{}
	params.Add("name", name)
	params.Add("ringer", strconv.FormatBool(ringer))
	params.Add("notifications", strconv.FormatBool(notifications))
	params.Add("autoReplyActive", strconv.FormatBool(autoReplyActive))
	params.Add("autoReplyText", strconv.FormatBool(autoReplyText))
	params.Add("callerIdEnabled", strconv.FormatBool(callerIDEnabled))
	params.Add("color", color)
	baseURL.RawQuery = params.Encode()
	req, err := http.NewRequest(http.MethodPut, baseURL.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create http request. Error: %s", err.Error())
	}
	setAuthHeader(req)
	resp, err := Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request to Burner: %s", err.Error())
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read body content. Error: %s", err.Error())
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to update burner. Burner returned: %s", string(bodyBytes))
	}
	var respBody []ConnectedBurner
	err = json.Unmarshal(bodyBytes, &respBody)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body. Error: %s", err.Error())
	}
	return respBody, nil
}
