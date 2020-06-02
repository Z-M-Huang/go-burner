package burner

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBurners(t *testing.T) {
	baseURL = "http://localhost:83"
	AuthToken = "abcd"
	mux := http.NewServeMux()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.NotEmpty(t, r.Header.Get("Authorization"))
		var ret []ConnectedBurner
		ret = append(ret, ConnectedBurner{})
		bytes, err := json.Marshal(ret)
		assert.Empty(t, err)
		w.Write(bytes)
		w.Header().Add("Content-Type", "application/json")
	})
	mux.Handle("/burners/", handler)
	go http.ListenAndServe(":83", mux)

	ret, err := GetBurners()
	assert.Empty(t, err)
	assert.NotEmpty(t, ret)
}

func TestGetBurnersInvalidResponse(t *testing.T) {
	baseURL = "http://localhost:84"
	AuthToken = "abcd"
	mux := http.NewServeMux()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("abcd"))
		w.Header().Add("Content-Type", "application/json")
	})
	mux.Handle("/burners/", handler)
	go http.ListenAndServe(":84", mux)

	ret, err := GetBurners()
	assert.NotEmpty(t, err)
	assert.Empty(t, ret)
}

func TestGetBurnersNot200(t *testing.T) {
	baseURL = "http://localhost:85"
	AuthToken = "abcd"
	mux := http.NewServeMux()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	})
	mux.Handle("/burners/", handler)
	go http.ListenAndServe(":85", mux)

	ret, err := GetBurners()
	assert.NotEmpty(t, err)
	assert.Empty(t, ret)
}

func TestGetBurnersInvalidAuthToken(t *testing.T) {
	AuthToken = ""
	ret, err := GetBurners()
	assert.NotEmpty(t, err)
	assert.Empty(t, ret)
}

func TestUpdateBurner(t *testing.T) {
	baseURL = "http://localhost:86"
	AuthToken = "abcd"
	burnerID := "1234"
	name := "testname"
	ringer := true
	notification := true
	autoReplyActive := true
	autoReplyText := true
	callerIDEnabled := true
	color := "testcolor"
	mux := http.NewServeMux()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method)
		assert.NotEmpty(t, r.Header.Get("Authorization"))
		err := r.ParseForm()
		assert.Empty(t, err)
		assert.Equal(t, name, r.FormValue("name"))
		assert.Equal(t, "true", r.FormValue("ringer"))
		assert.Equal(t, "true", r.FormValue("notifications"))
		assert.Equal(t, "true", r.FormValue("autoReplyActive"))
		assert.Equal(t, "true", r.FormValue("autoReplyText"))
		assert.Equal(t, "true", r.FormValue("callerIdEnabled"))
		assert.Equal(t, color, r.FormValue("color"))

		var ret []ConnectedBurner
		ret = append(ret, ConnectedBurner{})
		bytes, err := json.Marshal(ret)
		assert.Empty(t, err)
		w.Write(bytes)
		w.Header().Add("Content-Type", "application/json")
	})
	mux.Handle("/burners/"+burnerID+"/", handler)
	go http.ListenAndServe(":86", mux)

	ret, err := UpdateBurner(burnerID, name, ringer, notification, autoReplyActive, autoReplyText, callerIDEnabled, color)
	assert.Empty(t, err)
	assert.NotEmpty(t, ret)
}

func TestUpdateBurnersInvalidResponse(t *testing.T) {
	baseURL = "http://localhost:87"
	AuthToken = "abcd"
	burnerID := "1234"
	name := "testname"
	ringer := true
	notification := true
	autoReplyActive := true
	autoReplyText := true
	callerIDEnabled := true
	color := "testcolor"
	mux := http.NewServeMux()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("abcd"))
		w.Header().Add("Content-Type", "application/json")
	})
	mux.Handle("/burners/"+burnerID+"/", handler)
	go http.ListenAndServe(":87", mux)

	ret, err := UpdateBurner(burnerID, name, ringer, notification, autoReplyActive, autoReplyText, callerIDEnabled, color)
	assert.NotEmpty(t, err)
	assert.Empty(t, ret)
}

func TestUpdateBurnersNot200(t *testing.T) {
	baseURL = "http://localhost:88"
	AuthToken = "abcd"
	burnerID := "1234"
	name := "testname"
	ringer := true
	notification := true
	autoReplyActive := true
	autoReplyText := true
	callerIDEnabled := true
	color := "testcolor"
	mux := http.NewServeMux()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	})
	mux.Handle("/burners/"+burnerID+"/", handler)
	go http.ListenAndServe(":88", mux)

	ret, err := UpdateBurner(burnerID, name, ringer, notification, autoReplyActive, autoReplyText, callerIDEnabled, color)
	assert.NotEmpty(t, err)
	assert.Empty(t, ret)
}

func TestUpdateBurnersInvalidAuthToken(t *testing.T) {
	AuthToken = ""
	burnerID := "1234"
	name := "testname"
	ringer := true
	notification := true
	autoReplyActive := true
	autoReplyText := true
	callerIDEnabled := true
	color := "testcolor"
	ret, err := UpdateBurner(burnerID, name, ringer, notification, autoReplyActive, autoReplyText, callerIDEnabled, color)
	assert.NotEmpty(t, err)
	assert.Empty(t, ret)
}
