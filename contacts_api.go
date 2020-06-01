package burner

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

//Contacts Burner contacts
type Contacts struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	PhoneNumber string   `json:"phoneNumber"`
	Muted       bool     `json:"muted"`
	Blocked     bool     `json:"blocked"`
	BurnerIds   []string `json:"burnerIds"`
}

//GetContacts get a list of connected burners.
//See more at: https://developer.burnerapp.com/api-documentation/api/
func GetContacts(pageSize, page int, blocked bool) ([]Contacts, error) {
	baseURL, _ := url.Parse("https://app.burnerapp.com/contracts")
	params := url.Values{}
	params.Add("pageSize", strconv.Itoa(pageSize))
	params.Add("page", strconv.Itoa(page))
	params.Add("page", strconv.FormatBool(blocked))
	baseURL.RawQuery = params.Encode()
	req, err := http.NewRequest("GET", baseURL.String(), nil)
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
		return nil, fmt.Errorf("failed to get contacts. Burner returned: %s", string(bodyBytes))
	}
	var respBody []Contacts
	err = json.Unmarshal(bodyBytes, &respBody)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body. Error: %s", err.Error())
	}
	return respBody, nil
}

//UpdateContact update burner's contact
//See more at: https://developer.burnerapp.com/api-documentation/api/
func UpdateContact(contactPhoneNumber, name, newPhoneNumber string, blocked bool) error {
	baseURL, _ := url.Parse(fmt.Sprintf("https://app.burnerapp.com/contacts/%s", contactPhoneNumber))
	params := url.Values{}
	params.Add("name", name)
	params.Add("newPhoneNumber", newPhoneNumber)
	params.Add("blocked", strconv.FormatBool(blocked))
	baseURL.RawQuery = params.Encode()
	req, err := http.NewRequest(http.MethodPut, baseURL.String(), nil)
	if err != nil {
		return fmt.Errorf("failed to create http request. Error: %s", err.Error())
	}
	setAuthHeader(req)
	resp, err := Client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request to Burner: %s", err.Error())
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read body content. Error: %s", err.Error())
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to update contact. Burner returned: %s", string(bodyBytes))
	}
	return nil
}

//CreateContact create burner contact for multiple burnerIDs
//See more at: https://developer.burnerapp.com/api-documentation/api/
func CreateContact(name, phoneNumber string, burnerIds []string) error {
	idBytes, err := json.Marshal(burnerIds)
	if err != nil {
		return fmt.Errorf("failed to marshal burnerIds. Error: %s", err.Error())
	}

	requestMsg := url.Values{}
	requestMsg.Set("name", name)
	requestMsg.Set("phoneNumber", phoneNumber)
	requestMsg.Set("burnerIds", string(idBytes))
	req, err := http.NewRequest("POST", "https://api.burnerapp.com/contacts", strings.NewReader(requestMsg.Encode()))
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
		return fmt.Errorf("failed to get create contact. Burner returned: %s", string(bodyBytes))
	}
	return nil
}
