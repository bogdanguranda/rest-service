package web

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bogdanguranda/rest-service/util/logging"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

type MockLogger struct {
	LogOutput string
}

func (ml *MockLogger) Log(level logging.LogLevel, message string) {
	ml.LogOutput = message
}

type MockSearcher struct{}

func (ms MockSearcher) Search(value int) (int, int) {
	if value == 100 {
		return 0, 100
	}
	return -1, -1
}

func TestHandleInvalidValueParam(t *testing.T) {
	mockLogger := &MockLogger{}
	req := httptest.NewRequest("GET", "/value/notANumber", nil)
	w := httptest.NewRecorder()

	handler := NewGetValueHandler(&MockSearcher{}, mockLogger)
	router := mux.NewRouter()
	router.HandleFunc("/value/{value}", handler.Handle)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code, "Expected status mismatch")
	assert.Contains(t, mockLogger.LogOutput, "Invalid 'value' path parameter, must be a number.", "Log output mismatch")
}

func TestHandleValueNotFound(t *testing.T) {
	req := httptest.NewRequest("GET", "/value/50", nil)
	w := httptest.NewRecorder()

	handler := NewGetValueHandler(&MockSearcher{}, &MockLogger{})
	router := mux.NewRouter()
	router.HandleFunc("/value/{value}", handler.Handle)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code, "Expected status mismatch")
	response := Response{}
	_ = json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, 0, response.Index, "Index mismatch")
	assert.Equal(t, 0, response.Value, "Value mismatch")
	assert.Contains(t, response.Message, "Value 50 not found", "Message mismatch")
}

func TestHandleValueFound(t *testing.T) {
	mockLogger := &MockLogger{}
	req := httptest.NewRequest("GET", "/value/100", nil)
	w := httptest.NewRecorder()

	handler := NewGetValueHandler(&MockSearcher{}, mockLogger)
	router := mux.NewRouter()
	router.HandleFunc("/value/{value}", handler.Handle)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Expected status mismatch")
	response := Response{}
	_ = json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, 0, response.Index, "Index mismatch")
	assert.Equal(t, 100, response.Value, "Value mismatch")
	assert.Empty(t, response.Message, "Message mismatch")
}
