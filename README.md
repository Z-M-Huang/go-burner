# go-burner
Unofficial Burner API at https://developer.burnerapp.com/api-documentation. 

[![codecov](https://codecov.io/gh/Z-M-Huang/go-burner/branch/master/graph/badge.svg)](https://codecov.io/gh/Z-M-Huang/go-burner)
[![godoc](https://github.com/golang/gddo/blob/c782c79e0a3c3282dacdaaebeff9e6fd99cb2919/gddo-server/assets/status.svg)](https://pkg.go.dev/github.com/Z-M-Huang/go-burner?tab=doc)

Travis CI is failing 100%. Please run `go test -v ./...` locally to verify.

# Usage
### Run as a server
```go
  import "github.com/Z-M-Huang/go-burner"

  clientID := "fakeClientID"
  clientSecret := "fakeClientSecret"
  redirectURL := "https://example.com/auth_redirect"
  mux := http.NewServeMux()
  mux.Handle("/login", func(w http.ResponseWriter, r *http.Request){
    http.Redirect(w, r, burner.GetAuthorizeEndpoint(clientID, redirectURL), 301)
  })
  mux.Handle("/auth_redirect", func(w http.ResponseWriter, r *http.Request){
    code := r.URL.Query()["code"]
    client, burners, err := burner.HandleAuthCallback(code, clientID, clientSecret, redirectURL)
    err = client.Send(burners[0].ID, "+11234567890", "sample text", "")
  })
```

### Set access token directly
```go
  import "github.com/Z-M-Huang/go-burner"
  client := &burner.Client{
    AuthToken: "fakeAuthToken",
  }
  burnerID := "fakeBurnerID"
  err := client.Send(burnerID, "+11234567890", "sample text", "")
```

### Send message using Incoming Webhook
```go
  import "github.com/Z-M-Huang/go-burner"
  client := &burner.Client{
    IncomingWebhookURL: "https://api.burnerapp.com/webhooks/burner/fakeID?token=fakeToken",
  }
  err := client.SendIncomingWebhook("+11234567890", "sample text")
```