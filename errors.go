package authzen

import "errors"

var (
	ErrInvalidPassword      = errors.New("password is invalid")
	ErrWrongPassword        = errors.New("password is wrong")
	ErrInvalidEmail         = errors.New("email is invalid")
	ErrUserNotFound         = errors.New("user not found")
	ErrUserNotLoggedIn      = errors.New("user is not logged in")
	ErrEmailNotConfirmed    = errors.New("email has not been confirmed")
	ErrLoginSessionNotFound = errors.New("a valid login session wasn't found")
	ErrSessionValNotFound   = errors.New("session val not found")
	ErrInternal             = errors.New("internal server error")
	ErrUserExists           = errors.New("email is already linked to a user")
	ErrSendingEmail         = errors.New("problem sending the email")
)
