package entity

import (
	"errors"
	"regexp"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

var (
	store = Store{
		Id:          "a103c7e0-b560-4b01-9628-24553f136a6f",
		Name:        "store_entity_test",
		Tag:         "BoadGame Shop",
		Address:     "Shinagawa",
		StrongPoint: "High Quality",
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	update_store = Store{
		Id:          "a103c7e0-b560-4b01-9628-24553f136a6f",
		Name:        "uodate_store_entity_test",
		Tag:         "Update",
		Address:     "Shinagawa",
		StrongPoint: "High Quality",
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	pet_list = []Pet{
		{
			Id:        "55b0b78e-7910-4e11-a2a6-9934660a0618",
			Species:   "TEST",
			Name:      "entity-UT",
			Age:       5,
			StoreId:   "a103c7e0-b560-4b01-9628-24553f136a6f",
			CreatedAt: now,
			UpdatedAt: now,
			Status:    "TEST PASS",
		},
		{
			Id:        "da596c97-c8c7-48c5-ab6c-2312216b6b70",
			Species:   "TEST2",
			Name:      "entity-UT2",
			Age:       1,
			StoreId:   "a103c7e0-b560-4b01-9628-24553f136a6f",
			CreatedAt: now,
			UpdatedAt: now,
			Status:    "TEST PASS",
		},
	}
)

func TestStoreFindOK(t *testing.T) {
	db, mock, err := newMock()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	db.LogMode(true)

	var list_store = []Store{
		{
			Id:          "74684838-a5d9-47d8-91a4-ff63ce802763",
			Name:        "store_entity_test",
			Tag:         "BoadGame Shop",
			Address:     "Shinagawa",
			StrongPoint: "High Quality",
			CreatedAt:   now,
			UpdatedAt:   now,
		},
		{
			Id:          "a103c7e0-b560-4b01-9628-24553f136a6f",
			Name:        "store_entity_test",
			Tag:         "BoadGame Shop",
			Address:     "Mitaka",
			StrongPoint: "High Quality",
			CreatedAt:   now,
			UpdatedAt:   now,
		},
	}
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `stores`")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "tag", "address", "strong_point", "created_at", "updated_at"}).
			AddRow(list_store[0].Id, list_store[0].Name, list_store[0].Tag, list_store[0].Address, list_store[0].StrongPoint, list_store[0].CreatedAt, list_store[0].UpdatedAt).
			AddRow(list_store[1].Id, list_store[1].Name, list_store[1].Tag, list_store[1].Address, list_store[1].StrongPoint, list_store[1].CreatedAt, list_store[1].UpdatedAt))

	r := StoreRepository{DB: db}
	res, err := r.Find()
	if err != nil {
		t.Error(err)
	}
	assert.ElementsMatch(t, list_store, res)
}

func TestStoreCreateOK(t *testing.T) {
	db, mock, err := newMock()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	db.LogMode(true)

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(
		"INSERT INTO `stores` (`id`,`name`,`tag`,`address`,`strong_point`,`created_at`,`updated_at`) VALUES (?,?,?,?,?,?,?)")).
		WillReturnResult(sqlmock.NewResult(1, 7))
	mock.ExpectCommit()

	r := StoreRepository{DB: db}
	res, err := r.Create(store)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, store, res)
}

func TestStoreCreateBadRequestvalidation(t *testing.T) {
	store.Name = ""
	db, mock, err := newMock()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	db.LogMode(true)

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(
		"INSERT INTO `stores` (`id`,`name`,`tag`,`address`,`strong_point`,`created_at`,`updated_at`) VALUES (?,?,?,?,?,?,?)")).
		WillReturnResult(sqlmock.NewResult(1, 7)).
		WillReturnError(errors.New(`'Store.Name' Error:Field validation for 'Name' failed on the 'required' tag","time":"2020-07-02T14:20:37+09:00"`))
	mock.ExpectCommit()

	r := StoreRepository{DB: db}
	res, err := r.Create(store)

	_ = res
	expect := `'Store.Name' Error:Field validation for 'Name' failed on the 'required' tag","time":"2020-07-02T14:20:37+09:00"`
	assert.EqualError(t, err, expect)
}

func TestStoreCreateBadRequestDuplicate(t *testing.T) {
	store.Name = ""
	db, mock, err := newMock()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	db.LogMode(true)

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(
		"INSERT INTO `stores` (`id`,`name`,`tag`,`address`,`strong_point`,`created_at`,`updated_at`) VALUES (?,?,?,?,?,?,?)")).
		WillReturnResult(sqlmock.NewResult(1, 7)).
		WillReturnError(errors.New(`Error 1062: Duplicate entry 'Shinagawa-Pet-Shop' for key 'name' `))
	mock.ExpectCommit()

	r := StoreRepository{DB: db}
	res, err := r.Create(store)

	_ = res
	expect := `Error 1062: Duplicate entry 'Shinagawa-Pet-Shop' for key 'name' `
	assert.EqualError(t, err, expect)
}

func TestStoreGetOK(t *testing.T) {
	db, mock, err := newMock()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	db.LogMode(true)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `stores` WHERE (id = ?)")).
		WithArgs(store.Id).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "tag", "address", "strong_point", "created_at", "updated_at"}).
			AddRow(store.Id, store.Name, store.Tag, store.Address, store.StrongPoint, store.CreatedAt, store.UpdatedAt))

	r := StoreRepository{DB: db}
	res, err := r.Get(store.Id)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, res, store)
}

func TestStoreGetNotFound(t *testing.T) {
	db, mock, err := newMock()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	db.LogMode(true)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `stores` WHERE (id = ?)")).
		WithArgs(store.Id).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "tag", "address", "strong_point", "created_at", "updated_at"}).
			AddRow(store.Id, store.Name, store.Tag, store.Address, store.StrongPoint, store.CreatedAt, store.UpdatedAt)).
		WillReturnError(gorm.ErrRecordNotFound)

	r := StoreRepository{DB: db}
	res, err := r.Get(store.Id)
	_ = res
	assert.EqualError(t, err, "record not found")
}

func TestStoreUpdateOK(t *testing.T) {
	db, mock, err := newMock()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	db.LogMode(true)

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(
		"UPDATE `stores` SET `address` = ?, `created_at` = ?, `id` = ?, `name` = ?, `strong_point` = ?, `tag` = ?, `updated_at` = ? WHERE (id = ?)")).
		WithArgs(update_store.Address, update_store.CreatedAt, update_store.Id, update_store.Name, update_store.StrongPoint, update_store.Tag, update_store.UpdatedAt, store.Id).
		WillReturnResult(sqlmock.NewResult(1, 8))
	mock.ExpectCommit()

	r := StoreRepository{DB: db}
	res, err := r.Update(store.Id, update_store)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, res, update_store)
}

func TestStoreUpdateNotFound(t *testing.T) {
	db, mock, err := newMock()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	db.LogMode(true)

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(
		"UPDATE `stores` SET `address` = ?, `created_at` = ?, `id` = ?, `name` = ?, `strong_point` = ?, `tag` = ?, `updated_at` = ? WHERE (id = ?)")).
		WithArgs(update_store.Address, update_store.CreatedAt, update_store.Id, update_store.Name, update_store.StrongPoint, update_store.Tag, update_store.UpdatedAt, "test-id").
		WillReturnResult(sqlmock.NewResult(1, 8)).
		WillReturnError(gorm.ErrRecordNotFound)
	mock.ExpectCommit()

	r := StoreRepository{DB: db}
	res, err := r.Update("test-id", update_store)
	_ = res
	assert.EqualError(t, err, "record not found")
}

func TestStoreDeleteOK(t *testing.T) {
	db, mock, err := newMock()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	db.LogMode(true)

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(
		"DELETE FROM `stores` WHERE (id = ?)")).
		WithArgs(store.Id).
		WillReturnResult(sqlmock.NewResult(1, 7))
	mock.ExpectCommit()

	var empty Store
	r := StoreRepository{DB: db}
	res, err := r.Delete(store.Id)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, res, empty)
}

func TestStoreDeleteNotFound(t *testing.T) {
	db, mock, err := newMock()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	db.LogMode(true)

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(
		"DELETE FROM `stores` WHERE (id = ?)")).
		WithArgs("test-id").
		WillReturnResult(sqlmock.NewResult(1, 7)).
		WillReturnError(gorm.ErrRecordNotFound)
	mock.ExpectCommit()

	r := StoreRepository{DB: db}
	res, err := r.Delete("test-id")
	_ = res
	assert.EqualError(t, err, "record not found")
}

func TestStorePetsListOK(t *testing.T) {
	db, mock, err := newMock()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	db.LogMode(true)
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `stores` WHERE (id = ?) ORDER BY `stores`.`id` ASC LIMIT 1")).
		WithArgs(store.Id).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "tag", "address", "strong_point", "created_at", "updated_at"}).
			AddRow(store.Id, store.Name, store.Tag, store.Address, store.StrongPoint, store.CreatedAt, store.UpdatedAt))

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `pets` WHERE (store_id = ?)")).
		WithArgs(store.Id).
		WillReturnRows(sqlmock.NewRows([]string{"id", "species", "name", "age", "store_id", "created_at", "updated_at", "status"}).
			AddRow(pet_list[0].Id, pet_list[0].Species, pet_list[0].Name, pet_list[0].Age, pet_list[0].StoreId, pet_list[0].CreatedAt, pet_list[0].UpdatedAt, pet_list[0].Status).
			AddRow(pet_list[1].Id, pet_list[1].Species, pet_list[1].Name, pet_list[1].Age, pet_list[1].StoreId, pet_list[1].CreatedAt, pet_list[1].UpdatedAt, pet_list[1].Status))

	r := StoreRepository{DB: db}
	res, err := r.PetsList(store.Id)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, res, pet_list)
}

func TestStorePetsListNotFound(t *testing.T) {
	db, mock, err := newMock()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	db.LogMode(true)
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `stores` WHERE (id = ?) ORDER BY `stores`.`id` ASC LIMIT 1")).
		WithArgs(store.Id).
		WillReturnError(gorm.ErrRecordNotFound)

	r := StoreRepository{DB: db}
	res, err := r.PetsList(store.Id)
	_ = res
	assert.EqualError(t, err, "record not found")
}
