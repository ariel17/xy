package storage

import (
	"github.com/ariel17/xy/api/config"
)

var (
	Instance Storage
)

// Storage TODO
type Storage interface {
	Connect() error
	Close() error
	Insert(m *domain.Model) error
	Delete(m *domain.Model) error
}

// CreateStorage TODO
func New() error {
	if Instance == nil {
		Instance = &MongoDB{
			Auth: config.DatabaseAuth,
		}
		return Instance.Connect()
	}

	return nil
}
