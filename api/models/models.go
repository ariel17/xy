package models

// Model TODO
type Model interface {
	GetModelName() string
	Save() error
	Delete() error
}
