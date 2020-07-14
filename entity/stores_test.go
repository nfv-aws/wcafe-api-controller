package entity

import (
	"regexp"
	"testing"
	"time"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

var (
	store = Store{
		Id:          "st684838-a5d9-47d8-91a4-ff63ce802763",
		Name:        "entity-UT",
		Tag:         "Test",
		Address:     "Aomori",
		StrongPoint: "apple",
		CreatedAt:   now,
		UpdatedAt:   now,
		Status:      "TEST PASS",
	}

	update_store = Store{
		Id:          "st684838-a5d9-47d8-91a4-ff63ce802763",
		Name:        "entity-UT",
		Tag:         "Test",
		Address:     "Aomori",
		StrongPoint: "namahage",
		CreatedAt:   now,
		UpdatedAt:   now,
		Status:      "TEST PASS",
	}
	now = time.Now()
)

func newStoreMock() (*gorm.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}
	sm, err := gorm.Open("mysql", db)
	if err != nil {
		return nil, nil, err
	}
	return sm, mock, nil
}

func TestStoreFindOK(t *testing.T) {
	db, mock, err := newMock()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	db.LogMode(true)

	var list_store = []Store{
		{
			Id:          "st684838-a5d9-47d8-91a4-ff63ce802763",
			Name:        "entity-UT-1",
			Tag:         "Test",
			Address:     "Aomori",
			StrongPoint: "apple",
			CreatedAt:   now,
			UpdatedAt:   now,
			Status:      "TEST PASS",
		},
		{
			Id:          "aa684838-a5d9-47d8-91a4-ff63ce802763",
			Name:        "entity-UT-2",
			Tag:         "Test",
			Address:     "Aomori",
			StrongPoint: "apple",
			CreatedAt:   now,
			UpdatedAt:   now,
			Status:      "TEST PASS",
		},
	}

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `stores`")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "species", "name", "age", "store_id", "created_at", "updated_at", "status"}).
			AddRow(list_store[0].Id, list_store[0].Name, list_store[0].Tag, list_store[0].Address, list_store[0].StrongPoint, list_store[0].CreatedAt, list_store[0].UpdatedAt, list_store[0].Status).
			AddRow(list_store[1].Id, list_store[1].Name, list_store[1].Tag, list_store[1].Address, list_store[1].StrongPoint, list_store[1].CreatedAt, list_store[1].UpdatedAt, list_store[1].Status))

	r := StoreRepository{DB: db}
	res, err := r.Find()
	if err != nil {
		t.Error(err)
	}
	assert.ElementsMatch(t, res, list_store)
}

// func TestStoreCreateOK(t *testing.T) {
// 	db, mock, err := newMock()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	defer db.Close()
// 	db.LogMode(true)

// 	mock.ExpectBegin()
// 	mock.ExpectExec(regexp.QuoteMeta(
// 		"INSERT INTO `stores` (`id`,`species`,`name`,`age`,`store_id`,`created_at`,`updated_at`,`status`) VALUES (?,?,?,?,?,?,?,?)")).
// 		WillReturnResult(sqlmock.NewResult(1, 8))
// 	mock.ExpectCommit()

// 	r := StoreRepository{DB: db}
// 	res, err := r.Create(store)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	assert.Equal(t, res, store)
// }

// func TestStoreGetOK(t *testing.T) {
// 	db, mock, err := newMock()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	defer db.Close()
// 	db.LogMode(true)

// 	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `stores` WHERE (id = ?)")).
// 		WithArgs(store.Id).
// 		WillReturnRows(sqlmock.NewRows([]string{"id", "species", "name", "age", "store_id", "created_at", "updated_at", "status"}).
// 			AddRow(store.Id, store.Species, store.Name, store.Age, store.StoreId, store.CreatedAt, store.UpdatedAt, store.Status))

// 	r := StoreRepository{DB: db}
// 	res, err := r.Get(store.Id)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	assert.Equal(t, res, store)
// }

// func TestStoreUpdateOK(t *testing.T) {
// 	db, mock, err := newMock()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	defer db.Close()
// 	db.LogMode(true)

// 	mock.ExpectBegin()
// 	mock.ExpectExec(regexp.QuoteMeta(
// 		"UPDATE `stores` SET `age` = ?, `created_at` = ?, `id` = ?, `name` = ?, `species` = ?, `status` = ?, `store_id` = ?, `updated_at` = ? WHERE (id = ?)")).
// 		WithArgs(update_store.Age, update_store.CreatedAt, update_store.Id, update_store.Name, update_store.Species, update_store.Status, update_store.StoreId, update_store.UpdatedAt, store.Id).
// 		WillReturnResult(sqlmock.NewResult(1, 6))
// 	mock.ExpectCommit()

// 	r := StoreRepository{DB: db}
// 	res, err := r.Update(store.Id, update_store)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	assert.Equal(t, res, update_store)
// }

// func TestStoreDeleteOK(t *testing.T) {
// 	db, mock, err := newMock()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	defer db.Close()
// 	db.LogMode(true)

// 	mock.ExpectBegin()
// 	mock.ExpectExec(regexp.QuoteMeta(
// 		"DELETE FROM `stores` WHERE (id = ?)")).
// 		WithArgs(store.Id).
// 		WillReturnResult(sqlmock.NewResult(1, 8))
// 	mock.ExpectCommit()

// 	var empty Store
// 	r := StoreRepository{DB: db}
// 	res, err := r.Delete(store.Id)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	assert.Equal(t, res, empty)
// }
