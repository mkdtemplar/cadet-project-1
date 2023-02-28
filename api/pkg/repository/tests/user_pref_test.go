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

func TestUserPrefCreate(t *testing.T) {
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
	userPref := repository.NewUserPrefRepo()
	id := generate_id.GenerateID()
	userId, _ := uuid.Parse("26810825-80e2-46be-b72b-744c2de4a872")
	userPrefInput := models.UserPreferences{
		ID:          id,
		UserCountry: "USA",
		UserId:      userId,
	}

	userPrefObject := repository.NewUserPrefObject(id, userPrefInput.UserCountry, userPrefInput.UserId)

	mock.ExpectBegin()

	mock.ExpectExec(`INSERT INTO user_preferences VALUES ($1, $2, $2)`).
		WithArgs(userPrefObject.ID, userPrefObject.UserCountry, userPrefObject.UserId).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit()

	_, err = userPref.SaveUserPreferences(context.Background(), &userPrefObject)

	require.NoError(t, err)
}

func TestFindUserPreferences(t *testing.T) {
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
	userPref := repository.NewUserPrefRepo()

	id, _ := uuid.Parse("c7dd21c6-c26e-4702-a6d1-e9a3c714483e")
	userId, _ := uuid.Parse("7e0b6165-701e-478f-8282-4b6611a523cd")

	userPrefInput := &models.UserPreferences{
		ID:          id,
		UserCountry: "Argentina",
		UserId:      userId,
	}

	mock.ExpectQuery("SELECT(.*)").WithArgs(id).
		WillReturnRows(sqlmock.NewRows([]string{"id", "user_country", "user_id"}).
			AddRow(userPrefInput.ID, userPrefInput.UserCountry, userPrefInput.UserId))

	res, err := userPref.FindUserPreferences(context.Background(), id)

	require.NoError(t, err)
	require.Equal(t, res, userPrefInput)
}

func TestUpdateUserPref(t *testing.T) {
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
	userPref := repository.NewUserPrefRepo()

	id, err := uuid.Parse("36cfb416-780b-405a-bab1-b2e4a44bf226")
	userCountry := "Macedonia"
	const updateCountry = `UPDATE user_preferences SET "user_country" = $1 WHERE id = $2`

	mock.ExpectBegin()
	mock.ExpectExec(updateCountry).WithArgs(userCountry, id).
		WillReturnResult(sqlmock.NewResult(0, 1))

	mock.ExpectCommit()

	_, err = userPref.UpdateUserPref(context.Background(), id, userCountry)

	require.NoError(t, err)
}

func TestDeleteUserPreferences(t *testing.T) {
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
	userPref := repository.NewUserPrefRepo()

	id, err := uuid.Parse("2faeb522-e78d-4c6c-87b0-6bee3986f802")

	mock.ExpectBegin()
	mock.ExpectExec(`DELETE FROM user_preferences WHERE id = $1`).WithArgs(id).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	row, err := userPref.DeleteUserPreferences(context.Background(), id)

	require.NoError(t, err)
	require.Equal(t, row, int64(1))
}
