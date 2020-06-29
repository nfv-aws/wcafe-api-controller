package entity

import (
	"regexp"
	"testing"
	"time"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
<<<<<<< HEAD
)

var (
	p = Pet{
=======
	"github.com/stretchr/testify/assert"
)

var (
	pet = Pet{
>>>>>>> master
		Id:        "74684838-a5d9-47d8-91a4-ff63ce802763",
		Species:   "TEST",
		Name:      "entity-UT",
		Age:       5,
		StoreId:   "a103c7e0-b560-4b01-9628-24553f136a6f",
<<<<<<< HEAD
		CreatedAt: ct,
		UpdatedAt: ut,
		Status:    "TEST PASS",
	}
	ct, ut = time.Now(), time.Now()
=======
		CreatedAt: now,
		UpdatedAt: now,
		Status:    "TEST PASS",
	}

	update_pet = Pet{
		Id:        "74684838-a5d9-47d8-91a4-ff63ce802763",
		Species:   "UPDATE TEST",
		Name:      "entity-UT",
		Age:       6,
		StoreId:   "a103c7e0-b560-4b01-9628-24553f136a6f",
		CreatedAt: now,
		UpdatedAt: now,
		Status:    "TEST PASS",
	}
	now = time.Now()
>>>>>>> master
)

func newMock() (*gorm.DB, sqlmock.Sqlmock, error) {
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
	db, mock, err := newMock()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	db.LogMode(true)

<<<<<<< HEAD
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `pets`")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "species", "name", "age", "store_id", "created_at", "updated_at", "status"}).
			AddRow(p.Id, p.Species, p.Name, p.Age, p.StoreId, ct, ut, p.Status).
			AddRow("84684838-a5d9-47d8-91a4-ff63ce802763", "Canine", "Shiba-inu", 1, "a103c7e0-b560-4b01-9628-24553f136a6f", ct, ut, "PENDING_CREATE"))

	r := PetRepository{DB: db}
	_, err = r.Find()
	if err != nil {
		t.Error(err)
	}
=======
	var list_pet = []Pet{
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
			AddRow(list_pet[0].Id, list_pet[0].Species, list_pet[0].Name, list_pet[0].Age, list_pet[0].StoreId, list_pet[0].CreatedAt, list_pet[0].UpdatedAt, list_pet[0].Status).
			AddRow(list_pet[1].Id, list_pet[1].Species, list_pet[1].Name, list_pet[1].Age, list_pet[1].StoreId, list_pet[1].CreatedAt, list_pet[1].UpdatedAt, list_pet[1].Status))

	r := PetRepository{DB: db}
	res, err := r.Find()
	if err != nil {
		t.Error(err)
	}
	assert.ElementsMatch(t, res, list_pet)
>>>>>>> master
}

func TestPetCreateOK(t *testing.T) {
	db, mock, err := newMock()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	db.LogMode(true)

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(
		"INSERT INTO `pets` (`id`,`species`,`name`,`age`,`store_id`,`created_at`,`updated_at`,`status`) VALUES (?,?,?,?,?,?,?,?)")).
<<<<<<< HEAD
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	r := PetRepository{DB: db}
	_, err = r.Create(p)
	if err != nil {
		t.Error(err)
	}
=======
		WillReturnResult(sqlmock.NewResult(1, 8))
	mock.ExpectCommit()

	r := PetRepository{DB: db}
	res, err := r.Create(pet)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, res, pet)
>>>>>>> master
}

func TestPetGetOK(t *testing.T) {
	db, mock, err := newMock()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	db.LogMode(true)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `pets` WHERE (id = ?)")).
<<<<<<< HEAD
		WithArgs(p.Id).
		WillReturnRows(sqlmock.NewRows([]string{"id", "species", "name", "age", "store_id", "created_at", "updated_at", "status"}).
			AddRow(p.Id, p.Species, p.Name, p.Age, p.StoreId, ct, ut, p.Status))

	r := PetRepository{DB: db}
	_, err = r.Get(p.Id)
	if err != nil {
		t.Error(err)
	}
=======
		WithArgs(pet.Id).
		WillReturnRows(sqlmock.NewRows([]string{"id", "species", "name", "age", "store_id", "created_at", "updated_at", "status"}).
			AddRow(pet.Id, pet.Species, pet.Name, pet.Age, pet.StoreId, pet.CreatedAt, pet.UpdatedAt, pet.Status))

	r := PetRepository{DB: db}
	res, err := r.Get(pet.Id)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, res, pet)
>>>>>>> master
}

func TestPetUpdateOK(t *testing.T) {
	db, mock, err := newMock()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	db.LogMode(true)

<<<<<<< HEAD
	var (
		pe = Pet{
			Id:        "74684838-a5d9-47d8-91a4-ff63ce802763",
			Species:   "TEST",
			Name:      "entity-UT",
			Age:       6,
			StoreId:   "a103c7e0-b560-4b01-9628-24553f136a6f",
			CreatedAt: ct,
			UpdatedAt: ut,
			Status:    "TEST PASS",
		}
	)

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(
		"UPDATE `pets` SET `age` = ?, `created_at` = ?, `id` = ?, `name` = ?, `species` = ?, `status` = ?, `store_id` = ?, `updated_at` = ? WHERE (id = ?)")).
		WithArgs(pe.Age, ct, pe.Id, pe.Name, pe.Species, pe.Status, pe.StoreId, ut, p.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	r := PetRepository{DB: db}
	_, err = r.Update(p.Id, pe)
	if err != nil {
		t.Error(err)
	}
=======
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(
		"UPDATE `pets` SET `age` = ?, `created_at` = ?, `id` = ?, `name` = ?, `species` = ?, `status` = ?, `store_id` = ?, `updated_at` = ? WHERE (id = ?)")).
		WithArgs(update_pet.Age, update_pet.CreatedAt, update_pet.Id, update_pet.Name, update_pet.Species, update_pet.Status, update_pet.StoreId, update_pet.UpdatedAt, pet.Id).
		WillReturnResult(sqlmock.NewResult(1, 6))
	mock.ExpectCommit()

	r := PetRepository{DB: db}
	res, err := r.Update(pet.Id, update_pet)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, res, update_pet)
>>>>>>> master
}

func TestPetDeleteOK(t *testing.T) {
	db, mock, err := newMock()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	db.LogMode(true)

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(
		"DELETE FROM `pets` WHERE (id = ?)")).
<<<<<<< HEAD
		WithArgs(p.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	r := PetRepository{DB: db}
	_, err = r.Delete(p.Id)
	if err != nil {
		t.Error(err)
	}
=======
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
>>>>>>> master
}
