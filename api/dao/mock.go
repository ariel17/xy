package dao

import (
	"github.com/ariel17/xy/api/domain"
	"gopkg.in/mgo.v2/bson"
)

var (
	errors          map[bson.ObjectId]error
	insertedUsers   map[bson.ObjectId]domain.User
	insertedDevices map[bson.ObjectId]domain.Device
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
	if err := errors[u.ID]; err != nil {
		return err
	}
	insertedUsers[u.ID] = *u
	return nil
}

// DeleteUser TODO
func (m *MockDB) DeleteUser(id bson.ObjectId) error {
	if err := errors[id]; err != nil {
		return err
	}
	delete(insertedUsers, id)
	return nil
}

// GetUser TODO
func (m *MockDB) GetUser(id bson.ObjectId) (*domain.User, error) {
	if err := errors[id]; err != nil {
		return nil, err
	}
	u := insertedUsers[id]
	return &u, nil
}

// Devices + Users -------------------------------------------------------------

// GetUserDevices TODO
func (m *MockDB) GetUserDevices(id bson.ObjectId) ([]domain.Device, error) {
	if err := errors[id]; err != nil {
		return nil, err
	}
	devices := []domain.Device{}
	for _, v := range insertedDevices {
		if v.UserID == id {
			devices = append(devices, v)
		}
	}
	return devices, nil
}

// Devices ---------------------------------------------------------------------

// GetDevice TODO
func (m *MockDB) GetDevice(id bson.ObjectId) (*domain.Device, error) {
	if err := errors[id]; err != nil {
		return nil, err
	}
	u := insertedDevices[id]
	return &u, nil
}

// Mock support ----------------------------------------------------------------

// AddMockError puts an error to be raised on a mock operation.
func AddMockError(id bson.ObjectId, err error) {
	errors[id] = err
}

// CleanMocks removes all mock wires.
func CleanMocks() {
	errors = make(map[bson.ObjectId]error)
	insertedUsers = make(map[bson.ObjectId]domain.User)
	insertedDevices = make(map[bson.ObjectId]domain.Device)
}

func init() {
	CleanMocks()
}
