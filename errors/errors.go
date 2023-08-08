package errors

import "errors"

var (
	InvalidID       = errors.New("Invalid ID")
	CollectionEmpty = errors.New("Collection is empty")
	NotExists       = errors.New("Not exists")
	MissingParams   = errors.New("Missing parameters")
	InvalidArg     = errors.New("Invalid argument")
)
