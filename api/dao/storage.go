package dao

import (
	"github.com/ariel17/xy/api/domain"
)

// Storage TODO
type Storage interface {
	Connect() error
	Close() error
	InsertUser(u *domain.User) error
	DeleteUser(u *domain.User) error
	// TODO GetUser
}

// New TODO
func New() Storage {
	return &MongoDB{}
}
