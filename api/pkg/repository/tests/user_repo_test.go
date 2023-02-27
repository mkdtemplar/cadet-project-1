package tests

import (
	"cadet-project/pkg/models"
	"cadet-project/pkg/repository"
	"cadet-project/pkg/repository/generate_id"
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestUserRepoGetById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)
	repository.InitDb()
	userRepo := repository.NewUserRepo()
	id, _ := uuid.Parse("7e0b6165-701e-478f-8282-4b6611a523cd")
	var users = &models.User{
		ID:    id,
		Email: "test@test.com",
		Name:  "name surname",
	}
	mock.ExpectQuery("SELECT(.*)").WithArgs(id).
		WillReturnRows(sqlmock.NewRows([]string{"id", "email", "name"}).AddRow(users.ID, users.Email, users.Name))

	res, err := userRepo.GetById(context.Background(), users.ID)

	require.NoError(t, err)
	require.Equal(t, res, users)
}

func TestUserRepoCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)
	repository.InitDb()
	userRepo := repository.NewUserRepo()
	id := generate_id.GenerateID()
	var userCreate = models.User{
		ID:    id,
		Email: "testcreate@testcreate.com",
		Name:  "test create",
	}

	mock.ExpectBegin()
	mock.ExpectExec(`INSERT INTO users(id, email, name) VALUES ($1, $2, $3)`).
		WithArgs(userCreate.ID, userCreate.Email, userCreate.Name).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit()

	_, err = userRepo.Create(context.Background(), &userCreate)

	require.NoError(t, err)
}

func TestUserRepoDelete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)
	repository.InitDb()
	userRepo := repository.NewUserRepo()

	id, _ := uuid.Parse("bc692689-09aa-4f78-a6d8-c95a68ba4eac")

	mock.ExpectBegin()
	mock.ExpectExec(`DELETE FROM users WHERE id = $1`).WithArgs(id).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	_, err = userRepo.Delete(context.Background(), id)
	require.NoError(t, err)

}

func TestUserRepoGetUserEmail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)
	repository.InitDb()
	userRepo := repository.NewUserRepo()

	id, _ := uuid.Parse("7e0b6165-701e-478f-8282-4b6611a523cd")
	var users = &models.User{
		ID:    id,
		Email: "test@test.com",
		Name:  "name surname",
	}

	mock.ExpectBegin()

	mock.ExpectQuery(`SELECT * FROM users WHERE email = $1`).WithArgs(users.Email).
		WillReturnRows(sqlmock.NewRows([]string{"id", "email", "name"}).AddRow(users.ID, users.Email, users.Name))

	res, err := userRepo.GetUserEmail(context.Background(), users.Email)

	require.NoError(t, err)
	require.Equal(t, res, users)
}
