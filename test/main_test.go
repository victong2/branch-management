package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"example.com/tuto/controllers"
	"example.com/tuto/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {

	got := controllers.Add(4, 6)
	want := 10

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestGetAlbums(t *testing.T) {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	models.ConnectDatabase()

	// Router

	// Create a test router
	router := gin.Default()
	router.GET("/books", controllers.FindBooks)

	// router.GET("/", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	// })

	// Create a request to the API
	req, _ := http.NewRequest(http.MethodGet, "/books", nil)
	// req, _ := http.NewRequest(http.MethodGet, "/", nil)

	// Create a response recorder
	recorder := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(recorder, req)

	// Assert the response code is 200
	assert.Equal(t, http.StatusOK, recorder.Code)

	// Assert the response body contains expected data
	// expectedBody := `[{"id":"1","title":"Blue Train","artist":"John Coltrane","price":56.99},{"id":"2","title":"Jeru","artist":"Gerry Mulligan","price":17.99}]`
	// assert.JSONEq(t, expectedBody, recorder.Body.String())

	// Struct for the overall JSON response
	type BookResponse struct {
		Data []models.Book `json:"data"`
	}
	var bookResponse BookResponse

	// Convert response body to JSON
	err := json.Unmarshal(recorder.Body.Bytes(), &bookResponse)

	// Assert no error occurred during JSON unmarshalling
	assert.Nil(t, err)
}
