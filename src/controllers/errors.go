package controllers

import "errors"

var errInvalidUserData = errors.New("invalid username or password")
var errUserAlreadyExists = errors.New("this login or email is already used")
var errNoChanges = errors.New("no changes")
var errJsonDecode = errors.New("cannot decode json body")
var errNotFound = errors.New("item does not exist")

