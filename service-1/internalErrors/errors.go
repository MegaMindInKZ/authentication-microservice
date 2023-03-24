package internalErrors

import "errors"

var (
	IncorrectPasswordOrUsernameError = errors.New("Password or username isn't correct")
	CannotCreateTokenError           = errors.New("Cannot create token")
	UnexpectedTokenSigningMethod     = errors.New("Unexpected token signing method")
)
