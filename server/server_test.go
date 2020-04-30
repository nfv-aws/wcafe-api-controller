package server

import (
	"crypto/rand"
	"encoding/binary"
	"github.com/nfv-aws/wcafe-api-controller/db"
	"github.com/nfv-aws/wcafe-api-controller/entity"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strconv"
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
		testGETMethod(t, "/api/v1/users")
	})
	t.Run("TEST POST Method", func(t *testing.T) {
		testPOSTPetMethod(t, "/api/v1/pets")
		testPOSTStoreMethod(t, "/api/v1/stores")
		testPOSTUserMethod(t, "/api/v1/users")
	})
	t.Run("TEST PATCH Method", func(t *testing.T) {
		testPATCHMethod(t, "/api/v1/pets/:id")
		testPATCHMethod(t, "/api/v1/stores/:id")
		testPATCHMethod(t, "/api/v1/users/:id")
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

func testPOSTPetMethod(t *testing.T, endpoint string) {
	t.Helper()
	var store []entity.Store
	db := db.GetDB()
	db.Find(&store)
	bodyReader := strings.NewReader(`{"species": "Canine","name":"Shiba lnu", "age": 0, "store_id":"` + store[0].Id + `"}`)
	router := router()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", endpoint, bodyReader)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	router.ServeHTTP(w, req)
	assert.Equal(t, 201, w.Code)
}

func testPOSTStoreMethod(t *testing.T, endpoint string) {
	t.Helper()
	bodyReader := strings.NewReader(`{"name": "` + random() + `","tag":"abc","address":"Tokyo"}`)
	router := router()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", endpoint, bodyReader)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	router.ServeHTTP(w, req)
	assert.Equal(t, 201, w.Code)
}

func random() string {
	var n uint64
	binary.Read(rand.Reader, binary.LittleEndian, &n)
	return strconv.FormatUint(n, 36)
}

func testPOSTUserMethod(t *testing.T, endpoint string) {
	t.Helper()
	bodyReader := strings.NewReader(`{"name": "test man"}`)
	router := router()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", endpoint, bodyReader)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	router.ServeHTTP(w, req)
	assert.Equal(t, 201, w.Code)
}

func testPATCHMethod(t *testing.T, endpoint string) {
	t.Helper()
	bodyReader := strings.NewReader(`{"body": "test"}`)
	router := router()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PATCH", endpoint, bodyReader)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}
