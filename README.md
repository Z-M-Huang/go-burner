# go-burner
Unofficial Burner API at https://developer.burnerapp.com/api-documentation. 

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
    burners, err := burner.HandleAuthCallback(code, clientID, clientSecret, redirectURL)
    err = burner.Send(burners[0].ID, "+11234567890", "sample text", "")
  })
```

### Set access token directly
```go
  import "github.com/Z-M-Huang/go-burner"
  burner.AuthToken = "fakeAuthToken"
  burnerID := "fakeBurnerID"
  err := burner.Send(burnerID, "+11234567890", "sample text", "")
```