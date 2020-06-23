package server

import (
	"crypto/rand"
	"encoding/binary"
	math_rand "math/rand"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/nfv-aws/wcafe-api-controller/db"
	"github.com/nfv-aws/wcafe-api-controller/entity"
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

	var pet []entity.Pet
	var store []entity.Store
	var user []entity.User
	db := db.GetDB()
	db.Find(&pet)
	db.Find(&store)
	db.Find(&user)

	t.Run("TEST GET Method", func(t *testing.T) {
		testGETMethod(t, "/api/v1/pets")
		testGETMethod(t, "/api/v1/stores")
		testGETMethod(t, "/api/v1/users")
		testGETStorePetsMethod(t, "/api/v1/stores/"+store[0].Id+"/pets")
		testGETMethod(t, "/api/v1/clerks")
		testGETMethod(t, "/api/v1/supplies")
	})

	t.Run("TEST POST Method", func(t *testing.T) {
		testPOSTPetMethod(t, "/api/v1/pets")
		testPOSTStoreMethod(t, "/api/v1/stores")
		testPOSTUserMethod(t, "/api/v1/users")
		testPOSTUserEmail(t, "/api/v1/users")
		testPOSTUserBadRequestEmail(t, "/api/v1/users")
		testPOSTClerkMethod(t, "/api/v1/clerks")
	})

	t.Run("TEST PATCH Method", func(t *testing.T) {
		//pets
		testPATCHMethod(t, "/api/v1/pets/"+pet[0].Id, `{"species":"`+pet[0].Species+`", "name":"`+pet[0].Name+`", "age":10, "store_id":"`+store[0].Id+`"}`)
		testPATCHNoneId(t, "/api/v1/pets/testpetid", `{"species":"`+pet[0].Species+`", "name":"`+pet[0].Name+`", "age":10`)
		testPATCHNoneStoreId(t, "/api/v1/pets/"+pet[0].Id, `{"species":"`+pet[0].Species+`", "name":"`+pet[0].Name+`", "age":10, "store_id":"teststoreid"}`)
		testPATCHBadRequest(t, "/api/v1/pets/"+pet[0].Id, `{"species":" 278493, "name":"`+pet[0].Name+`", "age":10, "store_id":"`+store[0].Id+`"}`)
		testPATCHBadRequest(t, "/api/v1/pets/"+pet[0].Id, `{"species":"`+pet[0].Species+`", "name":5674, "age":10, "store_id":"`+store[0].Id+`"}`)
		testPATCHBadRequest(t, "/api/v1/pets/"+pet[0].Id, `{"species":"`+pet[0].Species+`", "name":"`+pet[0].Name+`", "age":"123", "store_id":"`+store[0].Id+`"}`)
		//stores
		testPATCHMethod(t, "/api/v1/stores/"+store[0].Id, `{"name":"`+store[0].Name+`", "tag": "`+store[0].Tag+`","address":"`+store[0].Address+`" }`)
		testPATCHNoneId(t, "/api/v1/stores/teststoreid", `{"name":"`+store[0].Name+`", "tag": "`+store[0].Tag+`","address":"`+store[0].Address+`" }`)
		testPATCHBadRequest(t, "/api/v1/stores/"+store[0].Id, `{"name":123, "tag": "`+store[0].Tag+`","address":"`+store[0].Address+`" }`)
		testPATCHBadRequest(t, "/api/v1/stores/"+store[0].Id, `{"name":"`+store[0].Name+`", "tag": 123,"address":"`+store[0].Address+`" }`)
		testPATCHBadRequest(t, "/api/v1/stores/"+store[0].Id, `{"name":"`+store[0].Name+`", "tag": "`+store[0].Tag+`","address":123 }`)
		//users
		testPATCHMethod(t, "/api/v1/users/"+user[0].Id, `{"number":3345,"email":"test@test.com"}`)
		testPATCHNoneId(t, "/api/v1/users/testuserid", `{"number":3345,"email":"test@test.com"}`)
		testPATCHBadRequest(t, "/api/v1/users/"+user[0].Id, `{"number":"3345","email":"test@test.com"}`)
		testPATCHBadRequest(t, "/api/v1/users/"+user[0].Id, `{"name":9473,"email":"test@test.com"}`)
		testPATCHBadRequest(t, "/api/v1/users/"+user[0].Id, `{"address":9473,"email":"test@test.com"}`)
		testPATCHBadRequest(t, "/api/v1/users/"+user[0].Id, `{"number":3345,"email":"test"}`)
	})

	t.Run("TEST DELETE Method", func(t *testing.T) {
		testDELETEMethod(t, "/api/v1/pets/"+pet[0].Id)
		testDELETEStoreMethod(t, "/api/v1/stores/"+store[0].Id)
		testDELETEMethod(t, "/api/v1/users/"+user[0].Id)
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

func testGETStorePetsMethod(t *testing.T, endpoint string) {
	t.Helper()
	router := router()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", endpoint, nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
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

func testPOSTUserEmail(t *testing.T, endpoint string) {
	t.Helper()
	math_rand.Seed(time.Now().UnixNano())
	random_num := math_rand.Intn(10000)
	bodyReader := strings.NewReader(`{"number":` + strconv.Itoa(random_num) + `,"email":"test@test.com"}`)
	router := router()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", endpoint, bodyReader)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	router.ServeHTTP(w, req)
	assert.Equal(t, 201, w.Code)
}

func testPOSTUserBadRequestEmail(t *testing.T, endpoint string) {
	t.Helper()
	math_rand.Seed(time.Now().UnixNano())
	random_num := math_rand.Intn(10000)
	bodyReader := strings.NewReader(`{"number":` + strconv.Itoa(random_num) + `,"email":"test"}`)
	router := router()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", endpoint, bodyReader)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	router.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
}

func testPOSTClerkMethod(t *testing.T, endpoint string) {
	t.Helper()
	bodyReader := strings.NewReader(`{"name": "testman"}`)
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

func testPATCHNoneId(t *testing.T, endpoint string, body string) {
	t.Helper()

	bodyReader := strings.NewReader(body)
	router := router()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PATCH", endpoint, bodyReader)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	router.ServeHTTP(w, req)
	assert.Equal(t, 404, w.Code)
}

func testPATCHNoneStoreId(t *testing.T, endpoint string, body string) {
	t.Helper()

	bodyReader := strings.NewReader(body)
	router := router()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PATCH", endpoint, bodyReader)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	router.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
}

func testPATCHBadRequest(t *testing.T, endpoint string, body string) {
	t.Helper()

	bodyReader := strings.NewReader(body)
	router := router()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PATCH", endpoint, bodyReader)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	router.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
}

func testDELETEMethod(t *testing.T, endpoint string) {
	t.Helper()
	router := router()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", endpoint, nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 204, w.Code)
}

func testDELETEStoreMethod(t *testing.T, endpoint string) {
	t.Helper()
	var store []entity.Store
	var pet []entity.Pet
	db := db.GetDB()

	db.Find(&store)
	db.Where("store_id = ?", store[0].Id).Delete(&pet)

	router := router()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", endpoint, nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 204, w.Code)
}
