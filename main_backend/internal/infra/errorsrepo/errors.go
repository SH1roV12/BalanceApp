package errorsrepo

import "errors"

var ErrCannotCreateUser error = errors.New("cannot create new user")