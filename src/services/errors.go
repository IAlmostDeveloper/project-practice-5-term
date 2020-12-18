package services

import "errors"

var errInvalidUserData = errors.New("invalid username or password")
var errUserAlreadyExists = errors.New("this login or email is already used")