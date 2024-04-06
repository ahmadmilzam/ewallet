package server_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ahmadmilzam/ewallet/handler/server"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Parallel()

	srv := server.New()

	wr := httptest.NewRecorder()
	srv.Router.ServeHTTP(wr, httptest.NewRequest(http.MethodGet, "/path-not-found", nil))

	res := wr.Result()
	assert.Equal(t, res.StatusCode, http.StatusNotFound)
}
