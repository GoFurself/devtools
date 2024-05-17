package userops

import (
	"testing"
)

// TestUserService tests the user service functions that are not dependent on the database
func TestUserService(t *testing.T) {
	repo := NewMockUserRepository()
	hasher := NewMockHasher()
	service := newUseropsService(repo, hasher)

	// Create user
	// Check that the service creates a user with correct values and sets the created time and enabled to true
	user := NewUser("einar.nysedt@centria.fi", "password", WithRole(1), WithFirstName("Einar"), WithLastName("Nystedt"))
	err := service.CreateUser(user)
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}

	if user.Created.IsZero() {
		t.Errorf("Expected created time to be set, got zero time")
	}

	if !user.Enabled {
		t.Errorf("Expected enabled to be true, got false")
	}

	if user.Metadata != "{}" {
		t.Errorf("Expected metadata to be empty, got %s", user.Metadata)
	}

	if user.Role != 1 {
		t.Errorf("Expected role 1, got %d", user.Role)
	}

	if user.FirstName != "Einar" {
		t.Errorf("Expected first name Einar, got %s", user.FirstName)
	}

	if user.LastName != "Nystedt" {
		t.Errorf("Expected last name Nystedt, got %s", user.LastName)
	}

	if user.Password != "hashedpassword" {
		t.Errorf("Expected password hashedpassword")
	}

	service.AuthenticateUser(user, "password")
}

func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{}
}

type MockUserRepository struct {
}

func (m *MockUserRepository) CreateUser(user *user) error {
	return nil
}

func (m *MockUserRepository) GetUserByID(id uint64) (*user, error) {
	return nil, nil
}

func (m *MockUserRepository) GetUserByEmail(email string) (*user, error) {
	return nil, nil
}

func (m *MockUserRepository) GetUsers() ([]*user, error) {
	return nil, nil
}

func (m *MockUserRepository) UpdateUser(user *user) error {
	return nil
}

func (m *MockUserRepository) DeleteUser(user *user) error {
	return nil
}

func (m *MockUserRepository) EnableUser(user *user) error {
	return nil
}

func (m *MockUserRepository) DisableUser(user *user) error {
	return nil
}

func NewMockHasher() *MockHasher {
	return &MockHasher{}
}

type MockHasher struct {
}

func (m *MockHasher) HashPassword(password string) (string, error) {
	hashed := "hashed" + password
	return hashed, nil
}

func (m *MockHasher) CheckPassword(password, hash string) (bool, error) {
	return password == hash, nil
}
