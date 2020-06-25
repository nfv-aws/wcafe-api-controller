package entity

import (
	"log"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"

	mocks "github.com/nfv-aws/wcafe-api-controller/sqlmocks"
)

var (
	p = Pet{
		Id:        "74684838-a5d9-47d8-91a4-ff63ce802763",
		Species:   "Canine",
		Name:      "Shiba-inu",
		Age:       1,
		StoreId:   "a103c7e0-b560-4b01-9628-24553f136a6f",
		CreatedAt: ct,
		UpdatedAt: ut,
		Status:    "PENDING_CREATE",
	}
	ct, ut = time.Now(), time.Now()
)

func TestPetFindOK(t *testing.T) {
	db, mock, err := mocks.NewMock()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	db.LogMode(true)

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "pets"`)).
		WillReturnRows(sqlmock.NewRows([]string{"id", "species", "name", "age", "store_id", "created_at", "updated_at", "status"}).
			AddRow("74684838-a5d9-47d8-91a4-ff63ce802763", "Canine", "Shiba-inu", 1, "a103c7e0-b560-4b01-9628-24553f136a6f", "2020-01-01 00:00:00", "2020-01-01 00:10:00", "PENDING_CREATE"))

	mock.ExpectCommit()

	r := PetRepository{DB: db}
	res, err := r.Find()
	log.Println(res)
	if err != nil {
		t.Error(err)
	}
}

// func TestPetCreateOK(t *testing.T) {
// 	db, mock, err := mocks.NewMock()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	defer db.Close()
// 	db.LogMode(true)

// 	mock.ExpectBegin()
// 	mock.ExpectQuery(regexp.QuoteMeta(
// 		`INSERT INTO "pets" VALUES ("Canine", "Shiba-inu", 1, "a103c7e0-b560-4b01-9628-24553f136a6f", "2020-01-01 00:00:00","2020-01-01 00:10:00","PENDING_CREATE")`)).
// 		WillReturnRows(
// 			sqlmock.NewRows([]string{"id", "species", "name", "age", "store_id", "created_at", "updated_at", "status"}).
// 				AddRow("74684838-a5d9-47d8-91a4-ff63ce802763", "Canine", "Shiba-inu", 1, "a103c7e0-b560-4b01-9628-24553f136a6f", "2020-01-01 00:00:00", "2020-01-01 00:10:00", "PENDING_CREATE"))
// 	mock.ExpectCommit()

// 	r := PetRepository{DB: db}
// 	var pet Pet
// 	_, err = r.Create(pet)
// 	// 	log.Println(_)
// 	if err != nil {
// 		t.Error(err)
// 	}
// }
