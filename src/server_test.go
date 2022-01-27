package main

import (
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestUserIndexHandler(t *testing.T) {
	router := NewRouter()

	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, `[{"name": "Taro", "email": "taro@example.com"}, {"name": "Jiro", "email": "jiro@example.com"}]`, rec.Body.String())
}

func TestUserShowHandler(t *testing.T) {
	router := NewRouter()

	req := httptest.NewRequest(http.MethodGet, "/users/jiro", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, `{"name": "Jiro", "email": "jiro@example.com"}`, rec.Body.String())
}

func TestUserCreateHandler(t *testing.T) {
	router := NewRouter()

	form := make(url.Values)
	form.Set("name", "Saburo")
	form.Set("email", "saburo@example.com")
	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(form.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusCreated, rec.Code)
	assert.JSONEq(t, `{"name": "Saburo", "email": "saburo@example.com"}`, rec.Body.String())
}
