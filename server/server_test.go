package server

import (
	"github.com/nfv-aws/wcafe-api-controller/db"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func setup() {
	// UT共通初期設定
	db.Init()
}

func tearDown() {
	// UT共通終了処理
}

func TestServer(t *testing.T) {
	setup()
	t.Run("TEST GET Method", func(t *testing.T) {
		testGETMethod(t, "/api/v1/pets")
		testGETMethod(t, "/api/v1/stores")
		// testGETMethod(t, "/api/v1/users")
	})
	t.Run("TEST POST Method", func(t *testing.T) {
		testPOSTMethod(t, "/api/v1/pets")
		testPOSTMethod(t, "/api/v1/stores")
		// testPOSTMethod(t, "/api/v1/users")
	})
	tearDown()
}

func testGETMethod(t *testing.T, endpoint string) {
	t.Helper()
	router := router()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", endpoint, nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func testPOSTMethod(t *testing.T, endpoint string) {
	t.Helper()
	bodyReader := strings.NewReader(`{"body": "test"}`)
	router := router()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", endpoint, bodyReader)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	router.ServeHTTP(w, req)
	assert.Equal(t, 201, w.Code)
}
