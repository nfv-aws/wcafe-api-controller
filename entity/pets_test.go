package entity

import (
	"reflect"
	"regexp"
	"testing"
	"time"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
)

var (
	p = Pet{
		Id:        "74684838-a5d9-47d8-91a4-ff63ce802763",
		Species:   "TEST",
		Name:      "entity-UT",
		Age:       5,
		StoreId:   "a103c7e0-b560-4b01-9628-24553f136a6f",
		CreatedAt: ct,
		UpdatedAt: ut,
		Status:    "TEST PASS",
	}
	ct, ut = time.Now(), time.Now()
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

	pe := Pet{
		Id:        "84684838-a5d9-47d8-91a4-ff63ce802763",
		Species:   "TEST2",
		Name:      "entity-UT2",
		Age:       10,
		StoreId:   "a103c7e0-b560-4b01-9628-24553f136a6f",
		CreatedAt: ct,
		UpdatedAt: ut,
		Status:    "TEST PASS2",
	}

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `pets`")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "species", "name", "age", "store_id", "created_at", "updated_at", "status"}).
			AddRow(p.Id, p.Species, p.Name, p.Age, p.StoreId, p.CreatedAt, p.UpdatedAt, p.Status).
			AddRow(pe.Id, pe.Species, pe.Name, pe.Age, pe.StoreId, pe.CreatedAt, pe.UpdatedAt, pe.Status))

	r := PetRepository{DB: db}
	_, err = r.Find()
	if err != nil {
		t.Error(err)
	}
	pl := [...]Pet{p, pe}
	reflect.DeepEqual(r, pl)
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
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	r := PetRepository{DB: db}
	_, err = r.Create(p)
	if err != nil {
		t.Error(err)
	}
	reflect.DeepEqual(r, p)
}

func TestPetGetOK(t *testing.T) {
	db, mock, err := newMock()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	db.LogMode(true)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `pets` WHERE (id = ?)")).
		WithArgs(p.Id).
		WillReturnRows(sqlmock.NewRows([]string{"id", "species", "name", "age", "store_id", "created_at", "updated_at", "status"}).
			AddRow(p.Id, p.Species, p.Name, p.Age, p.StoreId, p.CreatedAt, p.UpdatedAt, p.Status))

	r := PetRepository{DB: db}
	_, err = r.Get(p.Id)
	if err != nil {
		t.Error(err)
	}
	reflect.DeepEqual(r, p)
}

func TestPetUpdateOK(t *testing.T) {
	db, mock, err := newMock()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	db.LogMode(true)

	pu := Pet{
		Id:        "74684838-a5d9-47d8-91a4-ff63ce802763",
		Species:   "UPDATE TEST",
		Name:      "entity-UT",
		Age:       6,
		StoreId:   "a103c7e0-b560-4b01-9628-24553f136a6f",
		CreatedAt: ct,
		UpdatedAt: ut,
		Status:    "TEST PASS",
	}

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(
		"UPDATE `pets` SET `age` = ?, `created_at` = ?, `id` = ?, `name` = ?, `species` = ?, `status` = ?, `store_id` = ?, `updated_at` = ? WHERE (id = ?)")).
		WithArgs(pu.Age, pu.CreatedAt, pu.Id, pu.Name, pu.Species, pu.Status, pu.StoreId, pu.UpdatedAt, p.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	r := PetRepository{DB: db}
	_, err = r.Update(p.Id, pu)
	if err != nil {
		t.Error(err)
	}
	reflect.DeepEqual(r, pu)
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
		WithArgs(p.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	r := PetRepository{DB: db}
	_, err = r.Delete(p.Id)
	if err != nil {
		t.Error(err)
	}
	reflect.DeepEqual(r, p)
}
