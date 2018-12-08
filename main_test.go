package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/GeertJohan/go.rice"
	"github.com/foolin/echo-template/supports/gorice"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

var noErrorTests = []struct{ parameters, errorMsg string }{
	{"length=20", "Error"},
	{"length=32&digits=20", "Error"},
	{"length=32&symbols=20", "Error"},
	{"length=64&digits=20&symbols=20", "Error"},
	{"length=20&digits=8&denyrepeat=on", "Error"},
	{"length=64&symbols=20&denyrepeat=on", "Error"},
	{"length=64&symbols=20&denyrepeat=on&nopper=on", "Error"},
}
var errorTests = []struct{ parameters, errorMsg string }{
	{"", "Error: password can not have zero length"},
	{"length=0", "Error: password can not have zero length"},
	{"digits=2", "Error: password can not have zero length"},
	{"symbols=2", "Error: password can not have zero length"},
	{"denyrepeat=on", "Error: password can not have zero length"},
	{"noupper=on", "Error: password can not have zero length"},
	{"length=16&digits=20", "Error: number of digits and symbols must be less than total length"},
	{"length=16&symbols=20", "Error: number of digits and symbols must be less than total length"},
	{"length=32&digits=20&denyrepeat=on", "Error: number of digits exceeds available digits and repeats are not allowed"},
	{"length=64&symbols=40&denyrepeat=on", "Error: number of symbols exceeds available symbols and repeats are not allowed"},
}

func TestHealth(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, CheckHealth(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, `{"status":"OK"}`, rec.Body.String())
	}
}

func TestNoErrors(t *testing.T) {

	for _, s := range noErrorTests {
		// Setup
		e := echo.New()
		e.Renderer = gorice.New(rice.MustFindBox("views"))
		reader := strings.NewReader(s.parameters)
		req, _ := http.NewRequest(http.MethodPost, "/", reader)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		// Assertions
		if assert.NoError(t, GeneratePassword(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.NotContains(t, rec.Body.String(), s.errorMsg)
		}
	}
}
func TestErrors(t *testing.T) {

	for _, s := range errorTests {
		// Setup
		e := echo.New()
		e.Renderer = gorice.New(rice.MustFindBox("views"))
		reader := strings.NewReader(s.parameters)
		req, _ := http.NewRequest(http.MethodPost, "/", reader)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		// Assertions
		if assert.NoError(t, GeneratePassword(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Contains(t, rec.Body.String(), s.errorMsg)
		}
	}
}
