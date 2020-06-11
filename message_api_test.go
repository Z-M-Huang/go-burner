package burner

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSend(t *testing.T) {
	baseURL = "http://localhost:96"
	client := &Client{
		AuthToken: "abcd",
	}
	mux := http.NewServeMux()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.NotEmpty(t, r.Header.Get("Authorization"))
		assert.NotEmpty(t, http.MethodPost, r.Method)
		request := &sendMessageRequest{}
		bodyBytes, err := ioutil.ReadAll(r.Body)
		assert.Empty(t, err)
		assert.Empty(t, json.Unmarshal(bodyBytes, &request))
		assert.Equal(t, "1", request.BurnerID)
		assert.Equal(t, "2", request.ToNumber)
		assert.Equal(t, "3", request.Text)
		assert.Equal(t, "4", request.MediaURL)
	})
	mux.Handle("/messages/", handler)
	go http.ListenAndServe(":96", mux)

	err := client.Send("1", "2", "3", "4")
	assert.Empty(t, err)
}

func TestSendNot200(t *testing.T) {
	baseURL = "http://localhost:97"
	client := &Client{
		AuthToken: "abcd",
	}
	mux := http.NewServeMux()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	})
	mux.Handle("/messages/", handler)
	go http.ListenAndServe(":97", mux)

	err := client.Send("1", "2", "3", "4")
	assert.NotEmpty(t, err)
}

func TestSendInvalidAuthToken(t *testing.T) {
	client := &Client{
		AuthToken: "",
	}
	err := client.Send("1", "2", "3", "4")
	assert.NotEmpty(t, err)
}
