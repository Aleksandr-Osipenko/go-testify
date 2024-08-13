package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// тест на возврат кода 200 и не пустое тело
func TestMainHandlerWhenNoEmpty(t *testing.T) {

	req := httptest.NewRequest("GET", "/cafe?count=4&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	status := responseRecorder.Code
	assert.Equal(t, status, http.StatusOK)

	body := responseRecorder.Body.String()
	require.NotEmpty(t, body)
}

// тест на запрос города которого нет в списке
func TestMainHandlerWhenCity(t *testing.T) {

	req := httptest.NewRequest("GET", "/cafe?count=4&city=tula", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	status := responseRecorder.Code
	assert.Equal(t, status, http.StatusBadRequest)

	body := responseRecorder.Body.String()
	assert.Equal(t, body, "wrong city value")
}

// тест на запрос показ пяти кафе при имеющихся четырёх
func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4

	req := httptest.NewRequest("GET", "/cafe?count=5&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	body := responseRecorder.Body.String()
	list := strings.Split(body, ",")
	assert.Len(t, list, totalCount)
}
