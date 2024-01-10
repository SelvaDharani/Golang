package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock database implementation
// MockDB: A struct that embeds testify's mock.Mock, enabling it to act as a mock implementation of a database.
type MockDB struct {
	mock.Mock
}

// GetBookByID: A method on MockDB that simulates fetching a book by its ID.
// It uses m.Called to record the call and manage return values, which are set up in the test.
func (m *MockDB) GetBookByID(bookID string) (Book, bool) {
	args := m.Called(bookID)
	return args.Get(0).(Book), args.Bool(1)
}

// This function is a test case for testing the GetBookByID functionality.
func TestGetBookByID(t *testing.T) {
	// Create a mock object for the database

	// Setting Up the Mock and Test Expectations
	mockDB := new(MockDB)
	expectedBook := Book{
		ID:    "123",
		Title: "Sample",
	}
	mockDB.On("GetBookByID", "123").Return(expectedBook, true)

	// Create a request with a dummy response writer
	req, err := http.NewRequest("GET", "/books/123", nil)
	assert.NoError(t, err)

	//httptest.NewRecorder is used to create a response recorder, which will capture the HTTP response from the handler.
	w := httptest.NewRecorder()

	// Call the GetBookByID function with the mock database
	GetBookByID(w, req, mockDB)

	// Check the response status code
	assert.Equal(t, http.StatusOK, w.Code)
	// Define the expected JSON response
	expectedJSON := `{"id":"123","title":"Sample"}`

	// Compare the JSON response with the expected JSON
	assert.JSONEq(t, expectedJSON, w.Body.String())

	//checks that all expectations set on the mock (mockDB) were met.
	mockDB.AssertExpectations(t)
}
