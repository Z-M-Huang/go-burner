package burner

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthEndpoint(t *testing.T) {
	clientID := "123456"
	redirectURL := "https://tools.zh-code.com"
	assert.Equal(t, "https://app.burnerapp.com/oauth/authorize?client_id=123456&scope=messages:connect burners:read burners:write contacts:read contacts:write&redirect_uri=https://tools.zh-code.com", GetAuthorizeEndpoint(clientID, redirectURL))
}

func TestHandleAuthCallback(t *testing.T) {
	code := "fakeCode"
	clientID := "fakeClientID"
	clientSecret := "fakeClientSecret"
	redirectURL := "fakeURL"
	baseURL = "http://localhost"
	mux := http.NewServeMux()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Empty(t, r.ParseForm())
		assert.Equal(t, code, r.FormValue("code"))
		assert.Equal(t, clientID, r.FormValue("client_id"))
		assert.Equal(t, clientSecret, r.FormValue("client_secret"))
		assert.Equal(t, "authorization_code", r.FormValue("grant_type"))
		assert.Equal(t, redirectURL, r.FormValue("redirect_uri"))
		var ret []ConnectedBurner
		ret = append(ret, ConnectedBurner{})
		bytes, err := json.Marshal(&AccessResponse{
			AccessToken:      "abcd",
			ConnectedBurners: ret,
		})
		assert.Empty(t, err)
		w.Write(bytes)
		w.Header().Add("Content-Type", "application/json")
	})
	mux.Handle("/oauth/access", handler)
	go http.ListenAndServe(":80", mux)

	b, err := HandleAuthCallback(code, clientID, clientSecret, redirectURL)
	assert.Empty(t, err)
	assert.NotEmpty(t, b)
	assert.Equal(t, "abcd", AuthToken)
}

func TestHandleAuthCallbackFailNot200(t *testing.T) {
	code := "fakeCode"
	clientID := "fakeClientID"
	clientSecret := "fakeClientSecret"
	redirectURL := "fakeURL"
	baseURL = "http://localhost:81"
	mux := http.NewServeMux()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})
	mux.Handle("/oauth/access", handler)
	go http.ListenAndServe(":81", mux)

	_, err := HandleAuthCallback(code, clientID, clientSecret, redirectURL)
	assert.NotEmpty(t, err)
}

func TestHandleAuthCallbackFailInvalidResponse(t *testing.T) {
	code := "fakeCode"
	clientID := "fakeClientID"
	clientSecret := "fakeClientSecret"
	redirectURL := "fakeURL"
	baseURL = "http://localhost:82"
	mux := http.NewServeMux()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("absdafsd"))
		w.Header().Add("Content-Type", "application/json")
	})
	mux.Handle("/oauth/access", handler)
	go http.ListenAndServe(":82", mux)

	_, err := HandleAuthCallback(code, clientID, clientSecret, redirectURL)
	assert.NotEmpty(t, err)
}
