package service

import (
	"testing"

	"github.com/bogdanguranda/rest-service/util/logging"
	"github.com/stretchr/testify/assert"
)

type MockDB struct{}

func (m MockDB) GetInput() ([]int, error) {
	return []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}, nil
}

type MockLogger struct {
	LoggedMessages []string
}

func (ml *MockLogger) Log(level logging.LogLevel, message string) {
	ml.LoggedMessages = append(ml.LoggedMessages, message)
}

func TestSearchValueFound(t *testing.T) {
	mockDB := MockDB{}
	mockLogger := &MockLogger{}

	bs := NewBinarySearch(mockDB, mockLogger)
	index, value := bs.Search(50)

	assert.Equal(t, 4, index)
	assert.Equal(t, 50, value)
}

func TestSearchValueNotFoundButInRange(t *testing.T) {
	mockDB := MockDB{}
	mockLogger := &MockLogger{}

	bs := NewBinarySearch(mockDB, mockLogger)
	index, value := bs.Search(65)

	assert.Equal(t, 5, index)
	assert.Equal(t, 60, value)
}

func TestSearchValueFoundOutOfRange(t *testing.T) {
	mockDB := MockDB{}
	mockLogger := &MockLogger{}

	bs := NewBinarySearch(mockDB, mockLogger)
	index, value := bs.Search(105)

	assert.Equal(t, 9, index)
	assert.Equal(t, 100, value)
}

func TestSearchValueAtLowerBoundary(t *testing.T) {
	mockDB := MockDB{}
	mockLogger := &MockLogger{}

	bs := NewBinarySearch(mockDB, mockLogger)
	index, value := bs.Search(10)

	assert.Equal(t, 0, index)
	assert.Equal(t, 10, value)
}

func TestSearchValueAtUpperBoundary(t *testing.T) {
	mockDB := MockDB{}
	mockLogger := &MockLogger{}

	bs := NewBinarySearch(mockDB, mockLogger)
	index, value := bs.Search(100)

	assert.Equal(t, 9, index)
	assert.Equal(t, 100, value)
}

func TestSearchValueNotFound(t *testing.T) {
	mockDB := MockDB{}
	mockLogger := &MockLogger{}

	bs := NewBinarySearch(mockDB, mockLogger)
	index, value := bs.Search(200)

	assert.Equal(t, -1, index)
	assert.Equal(t, -1, value)
}
