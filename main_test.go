package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FetchX509SVID(t *testing.T) {
	rr := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/x509", nil)
	assert.NoError(t, err)

	s := Server{}
	s.NewRouter().ServeHTTP(rr, req)

	assert.NotNil(t, req)

	resp, err := ioutil.ReadAll(rr.Body)
	assert.NoError(t, err)

	assert.Equal(t, "FetchX509SVID", string(resp))
	assert.Equal(t, http.StatusNotImplemented, rr.Code)
}
