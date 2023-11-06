package db

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetInput(t *testing.T) {
	tempFile, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tempFile.Name())

	data := []byte("10\n20\n30\n40\n50\n")
	if _, err := tempFile.Write(data); err != nil {
		t.Fatal(err)
	}

	if err := tempFile.Close(); err != nil {
		t.Fatal(err)
	}

	db := NewFileDB(tempFile.Name())

	input, err := db.GetInput()
	assert.NoError(t, err, "GetInput returned an error")

	expected := []int{10, 20, 30, 40, 50}

	assert.Len(t, input, len(expected), "Input length mismatch")

	for i := range input {
		assert.Equal(t, expected[i], input[i], "Value mismatch at index %d", i)
	}
}
