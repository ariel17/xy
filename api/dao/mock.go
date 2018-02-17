package dao

import (
	"github.com/ariel17/xy/api/domain"
)

// MockDB is the testing implementation for validation & verification.
type MockDB struct {
}

// Connect does nothing; it is a mock!
func (m *MockDB) Connect() error {
	return nil
}

// InsertUser TODO
func (m *MockDB) InsertUser(u *domain.User) error {
	return nil
}

// DeleteUser TODO
func (m *MockDB) DeleteUser(u *domain.User) error {
	return nil
}

// GetUser TODO
func (m *MockDB) GetUser(id string) (*domain.User, error) {
	return &domain.User{}, nil
}
