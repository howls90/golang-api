package test

import (
	"api/routes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
func TestHelloWorld(t *testing.T) {
	// Build our expected body
	body := gin.H{
		"hello": "world",
	}
	routes.Init()
	w := performRequest(routes.Router, "GET", "/api/v1/hello")

	var response map[string]string
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	value, exists := response["hello"]

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, body["hello"], value)
}
