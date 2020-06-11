package burner

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetContacts(t *testing.T) {
	baseURL = "http://localhost:89"
	client := &Client{
		AuthToken: "abcd",
	}
	mux := http.NewServeMux()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.NotEmpty(t, r.Header.Get("Authorization"))
		assert.Equal(t, "1", r.URL.Query()["pageSize"][0])
		assert.Equal(t, "2", r.URL.Query()["page"][0])
		assert.Equal(t, "true", r.URL.Query()["blocked"][0])
		var ret []Contacts
		ret = append(ret, Contacts{})
		bytes, err := json.Marshal(ret)
		assert.Empty(t, err)
		w.Write(bytes)
		w.Header().Add("Content-Type", "application/json")
	})
	mux.Handle("/contacts/", handler)
	go http.ListenAndServe(":89", mux)

	ret, err := client.GetContacts(1, 2, true)
	assert.Empty(t, err)
	assert.NotEmpty(t, ret)
}

func TestGetContactsInvalidResponse(t *testing.T) {
	baseURL = "http://localhost:90"
	client := &Client{
		AuthToken: "abcd",
	}
	mux := http.NewServeMux()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("abcd"))
		w.Header().Add("Content-Type", "application/json")
	})
	mux.Handle("/contacts/", handler)
	go http.ListenAndServe(":90", mux)

	ret, err := client.GetContacts(1, 2, false)
	assert.NotEmpty(t, err)
	assert.Empty(t, ret)
}

func TestGetContactsNot200(t *testing.T) {
	baseURL = "http://localhost:91"
	client := &Client{
		AuthToken: "abcd",
	}
	mux := http.NewServeMux()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	})
	mux.Handle("/burners/", handler)
	go http.ListenAndServe(":91", mux)

	ret, err := client.GetContacts(1, 2, false)
	assert.NotEmpty(t, err)
	assert.Empty(t, ret)
}

func TestGetContactsInvalidAuthToken(t *testing.T) {
	client := &Client{
		AuthToken: "",
	}
	ret, err := client.GetContacts(1, 2, false)
	assert.NotEmpty(t, err)
	assert.Empty(t, ret)
}

func TestUpdateContact(t *testing.T) {
	baseURL = "http://localhost:92"
	client := &Client{
		AuthToken: "abcd",
	}
	mux := http.NewServeMux()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.NotEmpty(t, r.Header.Get("Authorization"))
		assert.Equal(t, http.MethodPut, r.Method)
		assert.Equal(t, "1", r.URL.Query()["name"][0])
		assert.Equal(t, "2", r.URL.Query()["phoneNumber"][0])
		assert.Equal(t, "true", r.URL.Query()["blocked"][0])
	})
	mux.Handle("/contacts/id/", handler)
	go http.ListenAndServe(":92", mux)

	err := client.UpdateContact("id", "1", "2", true)
	assert.Empty(t, err)
}

func TestUpdateContactNot200(t *testing.T) {
	baseURL = "http://localhost:93"
	client := &Client{
		AuthToken: "abcd",
	}
	mux := http.NewServeMux()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	})
	mux.Handle("/contacts/id/", handler)
	go http.ListenAndServe(":93", mux)

	err := client.UpdateContact("id", "1", "2", true)
	assert.NotEmpty(t, err)
}

func TestUpdateContactInvalidAuthToken(t *testing.T) {
	client := &Client{
		AuthToken: "",
	}
	err := client.UpdateContact("id", "1", "2", true)
	assert.NotEmpty(t, err)
}

func TestCreateContact(t *testing.T) {
	baseURL = "http://localhost:94"
	client := &Client{
		AuthToken: "abcd",
	}
	mux := http.NewServeMux()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.NotEmpty(t, r.Header.Get("Authorization"))
		assert.Equal(t, http.MethodPost, r.Method)
		err := r.ParseForm()
		assert.Empty(t, err)
		assert.Equal(t, "1", r.FormValue("name"))
		assert.Equal(t, "2", r.FormValue("phoneNumber"))
		assert.NotEmpty(t, r.FormValue("burnerIds"))
		w.WriteHeader(http.StatusOK)
	})
	mux.Handle("/contacts/", handler)
	go http.ListenAndServe(":94", mux)

	err := client.CreateContact("1", "2", []string{"abcd"})
	assert.Empty(t, err)
}

func TestCreateContactNot200(t *testing.T) {
	baseURL = "http://localhost:95"
	client := &Client{
		AuthToken: "abcd",
	}
	mux := http.NewServeMux()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	})
	mux.Handle("/contacts/", handler)
	go http.ListenAndServe(":95", mux)

	err := client.CreateContact("1", "2", []string{"abcd"})
	assert.NotEmpty(t, err)
}

func TestCreateContactInvalidAuthToken(t *testing.T) {
	client := &Client{
		AuthToken: "",
	}
	err := client.CreateContact("1", "2", []string{"abcd"})
	assert.NotEmpty(t, err)
}
