package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/KaustubhBuilds/acetlisto/stores"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func executeRequest(router *mux.Router, req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	return rr
}

func TestNew(t *testing.T) {
	store := stores.NewMemoryItemStore()
	router := New(store)
	assert.NotNil(t, router)
}

func TestHomeHandler(t *testing.T) {
	store := stores.NewMemoryItemStore()
	router := New(store)

	req, _ := http.NewRequest("GET", "/", nil)
	response := executeRequest(router, req)
	assert.Equal(t, http.StatusTemporaryRedirect, response.Code)

	assert.Equal(t, "/items/", response.Header().Get("Location"))
}
