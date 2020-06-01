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

//Client http client to call burner's service. If you want to call behind a proxy, you can create a new one and set
var Client *http.Client

//AuthToken Burner's auth token. If you have one already, you can set it directly
var AuthToken string

func init() {
	Client = &http.Client{}
	Client.Timeout = 2 * time.Second
}

func setAuthHeader(req *http.Request) {
	req.Header.Add("Authorization", "Bearer "+AuthToken)
}
