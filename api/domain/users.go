package domain

// User TODO
type User struct {
	Nick        string `json:"nick"`
	PendingPins []Pin  `json:"pending_pins"`
}

// NewUser TODO
func NewUser(nick string) *User {
	return &User{
		Nick:        nick,
		PendingPins: []Pin{},
	}
}
