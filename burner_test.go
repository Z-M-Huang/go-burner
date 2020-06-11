package burner

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetHeader(t *testing.T) {
	req, err := http.NewRequest("POST", baseURL, nil)
	assert.Empty(t, err)
	c := &Client{
		AuthToken: "TestSetHeader",
	}
	c.setAuthHeader(req)
	assert.Equal(t, "Bearer TestSetHeader", req.Header.Get("Authorization"))
}
