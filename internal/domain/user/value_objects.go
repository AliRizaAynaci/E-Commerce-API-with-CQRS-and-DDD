package user

import (
	"net/mail"
	"regexp"
	"strings"
)

// ID represents a user identifier
type ID string

// NewID creates a new ID
func NewID(id string) (ID, error) {
	if id == "" {
		return "", ErrInvalidID
	}
	return ID(id), nil
}

// String returns the string representation of the ID
func (id ID) String() string {
	return string(id)
}

// Email represents a user email
type Email string

// NewEmail creates a new Email
func NewEmail(email string) (Email, error) {
	if email == "" {
		return "", ErrInvalidEmail
	}

	// Validate email format
	_, err := mail.ParseAddress(email)
	if err != nil {
		return "", ErrInvalidEmail
	}

	return Email(email), nil
}

// String returns the string representation of the Email
func (e Email) String() string {
	return string(e)
}

// Password represents a user password
type Password string

// NewPassword creates a new Password
func NewPassword(password string) (Password, error) {
	if len(password) < 8 {
		return "", ErrInvalidPassword
	}

	// Check for at least one uppercase letter, one lowercase letter, and one number
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)

	if !hasUpper || !hasLower || !hasNumber {
		return "", ErrInvalidPassword
	}

	return Password(password), nil
}

// String returns the string representation of the Password
func (p Password) String() string {
	return string(p)
}

// Name represents a user name
type Name string

// NewName creates a new Name
func NewName(name string) (Name, error) {
	trimmedName := strings.TrimSpace(name)
	if trimmedName == "" || len(trimmedName) < 2 {
		return "", ErrInvalidName
	}
	return Name(trimmedName), nil
}

// String returns the string representation of the Name
func (n Name) String() string {
	return string(n)
}
