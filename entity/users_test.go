package entity

import (
	"regexp"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

var (
	user = User{
		Id:        "74684838-a5d9-47d8-91a4-ff63ce802763",
		Number:    123,
		Name:      "entity-UT",
		Address:   "Test City",
		Email:     "test@mail.com",
		CreatedAt: now,
		UpdatedAt: now,
		Status:    "TEST PASS",
	}

	update_user = User{
		Id:        "74684838-a5d9-47d8-91a4-ff63ce802763",
		Number:    321,
		Name:      "entity2-UT",
		Address:   "Test2 City",
		Email:     "test@mail.com",
		CreatedAt: now,
		UpdatedAt: now,
		Status:    "TEST PASS",
	}
)

func newUsersMock() (*gorm.DB, sqlmock.Sqlmock, error) {
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

//正常系
func TestUserValidator(t *testing.T) {
	var cases = []User{
		{
			Id:        "74684838-a5d9-47d8-91a4-ff63ce802763",
			Number:    123,
			Name:      "entity-UT",
			Address:   "Test City",
			Email:     "test@mail.com",
			CreatedAt: now,
			UpdatedAt: now,
			Status:    "TEST PASS",
		},
		{
			Id:        "74684838-a5d9-47d8-91a4-ff63ce802763",
			Number:    123,
			Name:      "entity-UT",
			Address:   "Test City",
			CreatedAt: now,
			UpdatedAt: now,
			Status:    "TEST PASS",
		},
		{
			Id:        "74684838-a5d9-47d8-91a4-ff63ce802763",
			Number:    123,
			Name:      "entity-UT",
			CreatedAt: now,
			UpdatedAt: now,
			Status:    "TEST PASS",
		},
		{
			Id:        "74684838-a5d9-47d8-91a4-ff63ce802763",
			Number:    123,
			CreatedAt: now,
			UpdatedAt: now,
			Status:    "TEST PASS",
		},
	}
	for _, tc := range cases {
		assert.Equal(t, nil, UserValidator(tc))
	}
}

func TestUserFindOK(t *testing.T) {
	db, mock, err := newUsersMock()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	db.LogMode(true)

	var list_user = []User{
		{
			Id:        "74684838-a5d9-47d8-91a4-ff63ce802763",
			Number:    123,
			Name:      "entity-UT",
			Address:   "Test City",
			Email:     "test@mail.com",
			CreatedAt: now,
			UpdatedAt: now,
			Status:    "TEST PASS",
		},
		{
			Id:        "74669088-s5d9-47n8-41c4-ff63ce808456",
			Number:    456,
			Name:      "entity3-UT",
			Address:   "Test3 City",
			Email:     "test3@mail.com",
			CreatedAt: now,
			UpdatedAt: now,
			Status:    "TEST PASS",
		},
	}

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users`")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "number", "name", "address", "email", "created_at", "updated_at", "status"}).
			AddRow(list_user[0].Id, list_user[0].Number, list_user[0].Name, list_user[0].Address, list_user[0].Email, list_user[0].CreatedAt, list_user[0].UpdatedAt, list_user[0].Status).
			AddRow(list_user[1].Id, list_user[1].Number, list_user[1].Name, list_user[1].Address, list_user[1].Email, list_user[1].CreatedAt, list_user[1].UpdatedAt, list_user[1].Status))

	r := UserRepository{DB: db}
	res, err := r.Find()
	if err != nil {
		t.Error(err)
	}
	assert.ElementsMatch(t, list_user, res)
}

func TestUserCreateOK(t *testing.T) {
	db, mock, err := newUsersMock()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	db.LogMode(true)

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(
		"INSERT INTO `users` (`id`,`number`,`name`,`address`,`email`,`created_at`,`updated_at`,`status`) VALUES (?,?,?,?,?,?,?,?)")).
		WillReturnResult(sqlmock.NewResult(1, 8))
	mock.ExpectCommit()

	r := UserRepository{DB: db}
	res, err := r.Create(user)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, user, res)
}

func TestUserGetOK(t *testing.T) {
	db, mock, err := newUsersMock()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	db.LogMode(true)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE (id = ?)")).
		WithArgs(user.Id).
		WillReturnRows(sqlmock.NewRows([]string{"id", "number", "name", "address", "email", "created_at", "updated_at", "status"}).
			AddRow(user.Id, user.Number, user.Name, user.Address, user.Email, user.CreatedAt, user.UpdatedAt, user.Status))

	r := UserRepository{DB: db}
	res, err := r.Get(user.Id)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, user, res)
}

func TestUserUpdateOK(t *testing.T) {
	db, mock, err := newUsersMock()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	db.LogMode(true)

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(
		"UPDATE `users` SET `address` = ?, `created_at` = ?, `email` = ?, `id` = ?, `name` = ?, `number` = ?, `status` = ?, `updated_at` = ? WHERE (id = ?)")).
		WithArgs(update_user.Address, update_user.CreatedAt, update_user.Email, update_user.Id, update_user.Name, update_user.Number, update_user.Status, update_user.UpdatedAt, user.Id).
		WillReturnResult(sqlmock.NewResult(1, 8))
	mock.ExpectCommit()

	r := UserRepository{DB: db}
	res, err := r.Update(user.Id, update_user)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, update_user, res)
}

func TestUserDeleteOK(t *testing.T) {
	db, mock, err := newUsersMock()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	db.LogMode(true)

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(
		"DELETE FROM `users` WHERE (id = ?)")).
		WithArgs(user.Id).
		WillReturnResult(sqlmock.NewResult(1, 8))
	mock.ExpectCommit()

	var empty User
	r := UserRepository{DB: db}
	res, err := r.Delete(user.Id)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, empty, res)
}

//準正常系
func TestUserValidatorNG(t *testing.T) {
	var cases = []User{
		{
			Id:        "74684838-a5d9-47d8-91a4-ff63ce802763",
			Name:      "entity-UT",
			Address:   "Test City",
			Email:     "test@mail.com",
			CreatedAt: now,
			UpdatedAt: now,
			Status:    "TEST PASS",
		},
		{
			Id:        "74684838-a5d9-47d8-91a4-ff63ce802763",
			Number:    123,
			Name:      "entity-UT",
			Address:   "Test City",
			Email:     "test",
			CreatedAt: now,
			UpdatedAt: now,
			Status:    "TEST PASS",
		},
	}
	for _, tc := range cases {
		assert.Error(t, UserValidator(tc))
	}
}
