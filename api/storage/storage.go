package storage

import (
	"github.com/ariel17/xy/api/config"
	"github.com/ariel17/xy/api/models"
)

var (
	Instance Storage
)

// Storage TODO
type Storage interface {
	Connect() error
	Close() error
	Insert(m *models.Model) error
	Delete(m *models.Model) error
}

// CreateStorage TODO
func CreateStorage() error {
	if Instance == nil {
		Instance = &MongoDB{
			Auth: config.DatabaseAuth,
		}
		return Instance.Connect()
	}

	return nil
}
