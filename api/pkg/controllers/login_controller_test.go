package controllers

import (
	"cadet-project/pkg/models"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
)

func TestServeHTTP(t *testing.T) {
	// Initialize test variables
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/", nil)

	// Create a new LoginController instance using mock dependencies
	lc := &LoginController{
		IUserRepository:      &mockUserRepository{},
		IShipPortsRepository: &mockShipPortsRepository{},
	}

	// Call the ServeHTTP function
	lc.ServeHTTP(w, r)

	// Check if proper HTTP status code and response body is returned
	if w.Code != http.StatusOK {
		t.Errorf("Expected status OK; got %v", w.Code)
	}

	expectedResponseBody := `{
            "id": "7e0b6165-701e-478f-8282-4b6611a523cd",
            "email": "test@test.com",
            "name": "name surname"
        }`
	if w.Body.String() != expectedResponseBody {
		t.Errorf("Expected response body %s; got %s", expectedResponseBody, w.Body.String())
	}
}

// Mock UserRepository implementation used in test cases
type mockUserRepository struct{}

func (m *mockUserRepository) Delete(ctx context.Context, uid uuid.UUID) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (m *mockUserRepository) GetById(ctx context.Context, id uuid.UUID) (*models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (m *mockUserRepository) GetUserEmail(ctx context.Context, email string) (*models.User, error) {
	id, _ := uuid.Parse("7e0b6165-701e-478f-8282-4b6611a523cd")
	return &models.User{
		ID:    id,
		Email: "test@test.com",
		Name:  "name surname",
	}, nil
}

func (m *mockUserRepository) Create(ctx context.Context, u *models.User) (*models.User, error) {
	return u, nil
}

// Mock ShipPortsRepository implementation used in test cases
type mockShipPortsRepository struct{}

func (m *mockShipPortsRepository) FindUserPrefPorts(ctx context.Context, usrpref *models.UserPreferences) (*models.UserPreferences, error) {
	//TODO implement me
	panic("implement me")
}

func (m *mockShipPortsRepository) FindUserPorts(ctx context.Context, id uuid.UUID) (*models.User, error) {
	id, _ = uuid.Parse("7e0b6165-701e-478f-8282-4b6611a523cd")
	return &models.User{
		ID:    id,
		Email: "test@test.com",
		Name:  "name surname",
	}, nil
}

func (m *mockShipPortsRepository) GetCityByName(ctx context.Context, name string) (string, error) {
	//TODO implement me
	panic("implement me")
}
