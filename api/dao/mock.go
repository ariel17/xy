package dao

import (
	"github.com/ariel17/xy/api/domain"
	"gopkg.in/mgo.v2/bson"
)

var (
	errors          map[string]error
	insertedUsers   map[string]domain.User
	insertedDevices map[string]domain.Device
)

// MockDB is the testing implementation for validation & verification.
type MockDB struct {
}

// Connect does nothing; it is a mock!
func (m *MockDB) Connect() error {
	return nil
}

// Users -----------------------------------------------------------------------

// InsertUser TODO
func (m *MockDB) InsertUser(u *domain.User) error {
	k := string(u.ID)
	if err := errors[k]; err != nil {
		return err
	}
	insertedUsers[k] = *u
	return nil
}

// DeleteUser TODO
func (m *MockDB) DeleteUser(id string) error {
	if err := errors[id]; err != nil {
		return err
	}
	delete(insertedUsers, id)
	return nil
}

// GetUser TODO
func (m *MockDB) GetUser(id string) (*domain.User, error) {
	if err := errors[id]; err != nil {
		return nil, err
	}
	u := insertedUsers[id]
	return &u, nil
}

// Devices + Users -------------------------------------------------------------

// GetUserDevices TODO
func (m *MockDB) GetUserDevices(id string) ([]domain.Device, error) {
	if err := errors[id]; err != nil {
		return nil, err
	}
	devices := []domain.Device{}
	for _, v := range insertedDevices {
		if v.UserID == bson.ObjectIdHex(id) {
			devices = append(devices, v)
		}
	}
	return devices, nil
}

// Devices ---------------------------------------------------------------------

// GetDevice TODO
func (m *MockDB) GetDevice(id string) (*domain.Device, error) {
	if err := errors[id]; err != nil {
		return nil, err
	}
	u := insertedDevices[id]
	return &u, nil
}

// Mock support ----------------------------------------------------------------

// AddMockError puts an error to be raised on a mock operation.
func AddMockError(id string, err error) {
	errors[id] = err
}

// CleanMocks removes all mock wires.
func CleanMocks() {
	errors = make(map[string]error)
	insertedUsers = make(map[string]domain.User)
	insertedDevices = make(map[string]domain.Device)
}

func init() {
	CleanMocks()
}
