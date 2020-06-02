package burner

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSend(t *testing.T) {
	baseURL = "http://localhost:96"
	AuthToken = "abcd"
	mux := http.NewServeMux()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.NotEmpty(t, r.Header.Get("Authorization"))
		assert.NotEmpty(t, http.MethodPost, r.Method)
		err := r.ParseForm()
		assert.Empty(t, err)
		assert.Equal(t, "1", r.FormValue("burnerId"))
		assert.Equal(t, "2", r.FormValue("toNumber"))
		assert.Equal(t, "3", r.FormValue("text"))
		assert.Equal(t, "4", r.FormValue("mediaUrl"))
	})
	mux.Handle("/messages", handler)
	go http.ListenAndServe(":96", mux)

	err := Send("1", "2", "3", "4")
	assert.Empty(t, err)
}

func TestSendNot200(t *testing.T) {
	baseURL = "http://localhost:97"
	AuthToken = "abcd"
	mux := http.NewServeMux()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	})
	mux.Handle("/messages", handler)
	go http.ListenAndServe(":97", mux)

	err := Send("1", "2", "3", "4")
	assert.NotEmpty(t, err)
}

func TestSendInvalidAuthToken(t *testing.T) {
	AuthToken = ""
	err := Send("1", "2", "3", "4")
	assert.NotEmpty(t, err)
}
