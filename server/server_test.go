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
		testPOSTMethod(t, "/api/v1/pets", `{"species": "Canine","name":"Shiba lnu", "age": 1, "store_id":"`+store[0].Id+`"}`)
		testPOSTMethod(t, "/api/v1/stores", `{"name": "`+random()+`","tag":"abc","address":"Tokyo"}`)
		testPOSTMethod(t, "/api/v1/users", `{"number":`+random_num()+`}`)
		testPOSTUserEmail(t, "/api/v1/users")
		testPOSTUserBadRequestEmail(t, "/api/v1/users")
		testPOSTMethod(t, "/api/v1/supplies", `{"name":"dog food", "price":500, "type": "food"}`)
		testPOSTBadRequest(t, "/api/v1/supplies", `{"price":500, "type": "food"}`)
		testPOSTBadRequest(t, "/api/v1/supplies", `{"name":"dog food", "type": "food"}`)
		testPOSTBadRequest(t, "/api/v1/supplies", `{"name":"dog food", "price":"500", "type": "food"}`)
	})

	t.Run("TEST PATCH Method", func(t *testing.T) {
		testPATCHMethod(t, "/api/v1/pets/"+pet[0].Id, `{"species":"`+pet[0].Species+`", "name":"`+pet[0].Name+`", "age":10, "store_id":"`+store[0].Id+`"}`)
		testPATCHNoneId(t, "/api/v1/pets/testpetid", `{"species":"`+pet[0].Species+`", "name":"`+pet[0].Name+`", "age":10`)
		testPATCHNoneStoreId(t, "/api/v1/pets/"+pet[0].Id, `{"species":"`+pet[0].Species+`", "name":"`+pet[0].Name+`", "age":10, "store_id":"teststoreid"}`)
		testPATCHBadRequest(t, "/api/v1/pets/"+pet[0].Id, `{"species":" 278493, "name":"`+pet[0].Name+`", "age":10, "store_id":"`+store[0].Id+`"}`)
		testPATCHBadRequest(t, "/api/v1/pets/"+pet[0].Id, `{"species":"`+pet[0].Species+`", "name":5674, "age":10, "store_id":"`+store[0].Id+`"}`)
		testPATCHBadRequest(t, "/api/v1/pets/"+pet[0].Id, `{"species":"`+pet[0].Species+`", "name":"`+pet[0].Name+`", "age":"123", "store_id":"`+store[0].Id+`"}`)
		testPATCHMethod(t, "/api/v1/stores/"+store[0].Id, `{"name":"`+store[0].Name+`", "tag": "`+store[0].Tag+`","address":"`+store[0].Address+`" }`)
		testPATCHNoneId(t, "/api/v1/stores/teststoreid", `{"name":"`+store[0].Name+`", "tag": "`+store[0].Tag+`","address":"`+store[0].Address+`" }`)
		testPATCHBadRequest(t, "/api/v1/stores/"+store[0].Id, `{"name":123, "tag": "`+store[0].Tag+`","address":"`+store[0].Address+`" }`)
		testPATCHBadRequest(t, "/api/v1/stores/"+store[0].Id, `{"name":"`+store[0].Name+`", "tag": 123,"address":"`+store[0].Address+`" }`)
		testPATCHBadRequest(t, "/api/v1/stores/"+store[0].Id, `{"name":"`+store[0].Name+`", "tag": "`+store[0].Tag+`","address":123 }`)
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

func random() string {
	var n uint64
	binary.Read(rand.Reader, binary.LittleEndian, &n)
	return strconv.FormatUint(n, 36)
}

func random_num() string {
	math_rand.Seed(time.Now().UnixNano())
	random_num := math_rand.Intn(10000)
	return strconv.Itoa(random_num)
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

func testPOSTMethod(t *testing.T, endpoint string, body string) {
	t.Helper()

	bodyReader := strings.NewReader(body)
	router := router()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", endpoint, bodyReader)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	router.ServeHTTP(w, req)
	assert.Equal(t, 201, w.Code)
}

func testPOSTBadRequest(t *testing.T, endpoint string, body string) {
	t.Helper()

	bodyReader := strings.NewReader(body)
	router := router()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", endpoint, bodyReader)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	router.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
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
