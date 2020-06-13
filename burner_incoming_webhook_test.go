package burner

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendInWebhook(t *testing.T) {
	client := &Client{
		IncomingWebhookURL: "http://localhost:10196/incomingWebHook",
	}
	mux := http.NewServeMux()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.NotEmpty(t, http.MethodPost, r.Method)
		request := &incomingWebhookRequest{}
		bodyBytes, err := ioutil.ReadAll(r.Body)
		assert.Empty(t, err)
		assert.Empty(t, json.Unmarshal(bodyBytes, &request))
		assert.Equal(t, "message", request.Intent)
		assert.Equal(t, "2", request.Data.ToNumber)
		assert.Equal(t, "3", request.Data.Text)
	})
	mux.Handle("/incomingWebHook", handler)
	go http.ListenAndServe(":10196", mux)

	err := client.SendIncomingWebhook("1", "2", "3")
	assert.Empty(t, err)
}

func TestSendInWebhookNot200(t *testing.T) {
	client := &Client{
		IncomingWebhookURL: "http://localhost:197/incomingWebHook",
	}
	mux := http.NewServeMux()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	})
	mux.Handle("/incomingWebHook", handler)
	go http.ListenAndServe(":197", mux)

	err := client.SendIncomingWebhook("1", "2", "3")
	assert.NotEmpty(t, err)
}

func TestSendInWebhookInvalidURL(t *testing.T) {
	client := &Client{
		IncomingWebhookURL: "",
	}
	err := client.SendIncomingWebhook("1", "2", "3")
	assert.NotEmpty(t, err)
}
