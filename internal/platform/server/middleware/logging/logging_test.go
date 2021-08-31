package logging

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)
	engine := gin.New()
	engine.Use(Middleware())

	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	httpRecorder := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/test-middleware", nil)
	require.NoError(t, err)

	engine.ServeHTTP(httpRecorder, req)

	require.NoError(t, w.Close())
	got, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	assert.Contains(t, string(got), "GET")
	assert.Contains(t, string(got), "/test-middleware")
	assert.Contains(t, string(got), "404")
}
