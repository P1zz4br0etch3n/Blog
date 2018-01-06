/*
    Matrikelnummern: 5836402, 2416160
*/

package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
)

func TestIndexHandler(t *testing.T) {
	testHandler(t, http.HandlerFunc(indexHandler))
}

func testHandler(t *testing.T, fn http.HandlerFunc) {
	server := httptest.NewServer(http.HandlerFunc(fn))
	defer server.Close()
	resp, err := http.Get(server.URL)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode, "wrong status")
}