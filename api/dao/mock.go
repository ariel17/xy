package dao

import (
	"github.com/ariel17/xy/api/domain"
)

var (
	errors   map[string]error
	inserted map[string]*domain.User
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
	k := string(u.ID)
	if err := errors[k]; err != nil {
		return err
	}
	inserted[k] = u
	return nil
}

// DeleteUser TODO
func (m *MockDB) DeleteUser(u *domain.User) error {
	k := string(u.ID)
	if err := errors[k]; err != nil {
		return err
	}
	delete(inserted, k)
	return nil
}

// GetUser TODO
func (m *MockDB) GetUser(id string) (*domain.User, error) {
	if err := errors[id]; err != nil {
		return nil, err
	}
	u := inserted[id]
	return u, nil
}

// AddMockError puts an error to be raised on a mock operation.
func AddMockError(id string, err error) {
	errors[id] = err
}

// CleanMocks removes all mock wires.
func CleanMocks() {
	errors = make(map[string]error)
	inserted = make(map[string]*domain.User)
}

func init() {
	CleanMocks()
}
