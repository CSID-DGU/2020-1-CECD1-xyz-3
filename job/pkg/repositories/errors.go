package repositories

import "errors"

var ErrNotFound = errors.New("not found")

func IsNotFoundErr(err error) bool {
	return errors.Is(err, ErrNotFound)
}

var ErrInvalidArgument = errors.New("invalid argument")

func IsInvalidArgumentErr(err error) bool {
	return errors.Is(err, ErrInvalidArgument)
}
