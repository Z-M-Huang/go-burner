//Package burner send sms over burner api
// Author: Z-M-Huang
// Repository: https://github.com/Z-M-Huang/go-burner
// License: MIT
// Note: Please do not abuse the Burner service!!!
package burner

import (
	"net/http"
	"time"
)

//Client Burner Client
type Client struct {
	//AuthToken Burner's auth token. If you have one already, you can set it directly
	AuthToken string
}

//HTTPClient http client to send burner request.
var HTTPClient *http.Client

func init() {
	HTTPClient = &http.Client{}
	HTTPClient.Timeout = 2 * time.Second
}

var baseURL string = "https://api.burnerapp.com/v1"

func (c *Client) setAuthHeader(req *http.Request) {
	req.Header.Add("Authorization", "Bearer "+c.AuthToken)
}
