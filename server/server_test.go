package server

import (
	"crypto/rand"
	"encoding/binary"
	"github.com/nfv-aws/wcafe-api-controller/db"
	"github.com/nfv-aws/wcafe-api-controller/entity"
	"github.com/stretchr/testify/assert"
	math_rand "math/rand"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"
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

	var pet []entity.Pet
	var store []entity.Store
	var user []entity.User
	db := db.GetDB()
	db.Find(&pet)
	db.Find(&store)
	db.Find(&user)

	t.Run("TEST PATCH Method", func(t *testing.T) {
		testPATCHMethod(t, "/api/v1/pets/"+pet[0].Id, `{"species":"`+pet[0].Species+`", "name":"`+pet[0].Name+`", "age":10, "store_id":"`+store[0].Id+`"}`)
		testPATCHMethod(t, "/api/v1/stores/"+store[0].Id, `{"name":"`+store[0].Name+`", "tag": "`+store[0].Tag+`","address":"`+store[0].Address+`" }`)
		testPATCHMethod(t, "/api/v1/users/"+user[0].Id, `{"number":334}`)
	})

	t.Run("TEST DELETE Method", func(t *testing.T) {
		testDELETEMethod(t, "/api/v1/pets/"+pet[0].Id)
		testDELETEMethod(t, "/api/v1/stores/"+store[5].Id)
		// testDELETEMethod(t, "/api/v1/users/"+user[0].Id)
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
	bodyReader := strings.NewReader(`{"species": "Canine","name":"Shiba lnu", "age": 1, "store_id":"` + store[0].Id + `"}`)
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
	math_rand.Seed(time.Now().UnixNano())
	random_num := math_rand.Intn(10000)
	bodyReader := strings.NewReader(`{"number":` + strconv.Itoa(random_num) + `}`)
	router := router()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", endpoint, bodyReader)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	router.ServeHTTP(w, req)
	assert.Equal(t, 201, w.Code)
}

func testPATCHMethod(t *testing.T, endpoint string, body string) {
	t.Helper()

	bodyReader := strings.NewReader(body)
	router := router()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PATCH", endpoint, bodyReader)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func testDELETEMethod(t *testing.T, endpoint string) {
	t.Helper()
	router := router()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", endpoint, nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 204, w.Code)
}
