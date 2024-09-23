package test

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"example.com/tuto/config"
	"example.com/tuto/handlers"
	"example.com/tuto/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetAlbums(t *testing.T) {
	config, err := config.LoadConfig("..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	models.ConnectDatabase(config)

	gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.GET("/books", handlers.FindBooks)

	req, _ := http.NewRequest(http.MethodGet, "/books", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)
	assert.Equal(t, http.StatusOK, recorder.Code)

	// Struct for the overall JSON response
	type BookResponse struct {
		Data []models.Book `json:"data"`
	}

	var bookResponse BookResponse
	err = json.Unmarshal(recorder.Body.Bytes(), &bookResponse)
	assert.Nil(t, err)

	// TODO: Add more asserts to check the response

	// Assert the response body contains expected data
	// expectedBody := `[{"id":"1","title":"Blue Train","artist":"John Coltrane","price":56.99},{"id":"2","title":"Jeru","artist":"Gerry Mulligan","price":17.99}]`
	// assert.JSONEq(t, expectedBody, recorder.Body.String())

	// test redirection with trailing slash
	// Create multiple root branch (without parent)
	// Name and parent_id can be changed with Patch.
}
