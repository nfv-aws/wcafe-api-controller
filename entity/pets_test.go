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
	pet = Pet{
		Id:        "74684838-a5d9-47d8-91a4-ff63ce802763",
		Species:   "TEST",
		Name:      "entity-UT",
		Age:       5,
		StoreId:   "a103c7e0-b560-4b01-9628-24553f136a6f",
		CreatedAt: now,
		UpdatedAt: now,
		Status:    "TEST PASS",
	}

	updatePet = Pet{
		Id:        "74684838-a5d9-47d8-91a4-ff63ce802763",
		Species:   "UPDATE TEST",
		Name:      "entity-UT",
		Age:       6,
		StoreId:   "a103c7e0-b560-4b01-9628-24553f136a6f",
		CreatedAt: now,
		UpdatedAt: now,
		Status:    "TEST PASS",
	}

	badRequestPet = Pet{
		Id:        "74684838-a5d9-47d8-91a4-ff63ce802763",
		Age:       500,
		CreatedAt: now,
		UpdatedAt: now,
		Status:    "TEST PASS",
	}

	now = time.Now()
)

func newPetsMock() (*gorm.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}
	um, err := gorm.Open("mysql", db)
	if err != nil {
		return nil, nil, err
	}
	return um, mock, nil
}

func TestPetFindOK(t *testing.T) {
	db, mock, err := newPetsMock()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	db.LogMode(true)

	var listPet = []Pet{
		{
			Id:        "74684838-a5d9-47d8-91a4-ff63ce802763",
			Species:   "TEST",
			Name:      "entity-UT",
			Age:       1,
			StoreId:   "a103c7e0-b560-4b01-9628-24553f136a6f",
			CreatedAt: now,
			UpdatedAt: now,
			Status:    "TEST PASS",
		},
		{
			Id:        "84684838-a5d9-47d8-91a4-ff63ce802763",
			Species:   "TEST2",
			Name:      "entity-UT2",
			Age:       2,
			StoreId:   "a103c7e0-b560-4b01-9628-24553f136a6f",
			CreatedAt: now,
			UpdatedAt: now,
			Status:    "TEST PASS",
		},
	}

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `pets`")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "species", "name", "age", "store_id", "created_at", "updated_at", "status"}).
			AddRow(listPet[0].Id, listPet[0].Species, listPet[0].Name, listPet[0].Age, listPet[0].StoreId, listPet[0].CreatedAt, listPet[0].UpdatedAt, listPet[0].Status).
			AddRow(listPet[1].Id, listPet[1].Species, listPet[1].Name, listPet[1].Age, listPet[1].StoreId, listPet[1].CreatedAt, listPet[1].UpdatedAt, listPet[1].Status))

	r := PetRepository{DB: db}
	res, err := r.Find()
	if err != nil {
		t.Error(err)
	}
	assert.ElementsMatch(t, res, listPet)
}

func TestPetCreateOK(t *testing.T) {
	db, mock, err := newPetsMock()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	db.LogMode(true)

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(
		"INSERT INTO `pets` (`id`,`species`,`name`,`age`,`store_id`,`created_at`,`updated_at`,`status`) VALUES (?,?,?,?,?,?,?,?)")).
		WillReturnResult(sqlmock.NewResult(1, 8))
	mock.ExpectCommit()

	r := PetRepository{DB: db}
	res, err := r.Create(pet)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, res, pet)
}

func TestPetCreateBadRequest(t *testing.T) {
	db, mock, err := newPetsMock()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	db.LogMode(true)

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(
		"INSERT INTO `pets` (`id`,`age`,`created_at`,`updated_at`,`status`) VALUES (?,?,?,?,?)")).
		WillReturnResult(sqlmock.NewResult(1, 6))
	mock.ExpectCommit()

	r := PetRepository{DB: db}
	_, err = r.Create(badRequestPet)
	if err != nil {
		assert.Error(t, err)
	}
}

func TestPetGetOK(t *testing.T) {
	db, mock, err := newPetsMock()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	db.LogMode(true)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `pets` WHERE (id = ?)")).
		WithArgs(pet.Id).
		WillReturnRows(sqlmock.NewRows([]string{"id", "species", "name", "age", "store_id", "created_at", "updated_at", "status"}).
			AddRow(pet.Id, pet.Species, pet.Name, pet.Age, pet.StoreId, pet.CreatedAt, pet.UpdatedAt, pet.Status))

	r := PetRepository{DB: db}
	res, err := r.Get(pet.Id)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, res, pet)
}

func TestPetGetNotFound(t *testing.T) {
	db, mock, err := newPetsMock()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	db.LogMode(true)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `pets` WHERE (id = ?)")).
		WithArgs(pet.Id).
		WillReturnRows(sqlmock.NewRows([]string{"id", "species", "name", "age", "store_id", "created_at", "updated_at", "status"}).
			AddRow(pet.Id, pet.Species, pet.Name, pet.Age, pet.StoreId, pet.CreatedAt, pet.UpdatedAt, pet.Status))

	r := PetRepository{DB: db}
	_, err = r.Get("111")
	if err != nil {
		assert.Error(t, err)
	}
}

func TestPetUpdateOK(t *testing.T) {
	db, mock, err := newPetsMock()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	db.LogMode(true)

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(
		"UPDATE `pets` SET `age` = ?, `created_at` = ?, `id` = ?, `name` = ?, `species` = ?, `status` = ?, `store_id` = ?, `updated_at` = ? WHERE (id = ?)")).
		WithArgs(updatePet.Age, updatePet.CreatedAt, updatePet.Id, updatePet.Name, updatePet.Species, updatePet.Status, updatePet.StoreId, updatePet.UpdatedAt, pet.Id).
		WillReturnResult(sqlmock.NewResult(1, 6))
	mock.ExpectCommit()

	r := PetRepository{DB: db}
	res, err := r.Update(pet.Id, updatePet)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, res, updatePet)
}

func TestPetUpdateBadRequest(t *testing.T) {
	db, mock, err := newPetsMock()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	db.LogMode(true)

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(
		"UPDATE `pets` SET `age` = ?, `created_at` = ?, `id` = ?,  `status` = ?,  `updated_at` = ? WHERE (id = ?)")).
		WithArgs(badRequestPet.Age, badRequestPet.CreatedAt, badRequestPet.Id, badRequestPet.Status, badRequestPet.UpdatedAt, pet.Id).
		WillReturnResult(sqlmock.NewResult(1, 6))
	mock.ExpectCommit()

	r := PetRepository{DB: db}
	_, err = r.Update(pet.Id, badRequestPet)
	if err != nil {
		assert.Error(t, err)
	}
}

func TestPetUpdateNotFound(t *testing.T) {
	db, mock, err := newPetsMock()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	db.LogMode(true)

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(
		"UPDATE `pets` SET `age` = ?, `created_at` = ?, `id` = ?, `name` = ?, `species` = ?, `status` = ?, `store_id` = ?, `updated_at` = ? WHERE (id = ?)")).
		WithArgs(updatePet.Age, updatePet.CreatedAt, updatePet.Id, updatePet.Name, updatePet.Species, updatePet.Status, updatePet.StoreId, updatePet.UpdatedAt, pet.Id).
		WillReturnResult(sqlmock.NewResult(1, 6))
	mock.ExpectCommit()

	r := PetRepository{DB: db}
	_, err = r.Update("111", updatePet)
	if err != nil {
		assert.Error(t, err)
	}
}

func TestPetDeleteOK(t *testing.T) {
	db, mock, err := newPetsMock()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	db.LogMode(true)

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(
		"DELETE FROM `pets` WHERE (id = ?)")).
		WithArgs(pet.Id).
		WillReturnResult(sqlmock.NewResult(1, 8))
	mock.ExpectCommit()

	var empty Pet
	r := PetRepository{DB: db}
	res, err := r.Delete(pet.Id)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, res, empty)
}

func TestPetDeleteNotFound(t *testing.T) {
	db, mock, err := newPetsMock()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	db.LogMode(true)

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(
		"DELETE FROM `pets` WHERE (id = ?)")).
		WithArgs(pet.Id).
		WillReturnResult(sqlmock.NewResult(1, 8))
	mock.ExpectCommit()

	r := PetRepository{DB: db}
	_, err = r.Delete("111")
	if err != nil {
		assert.Error(t, err)
	}
}
