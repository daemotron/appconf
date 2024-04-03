package appconf

import "errors"

// The ErrAllUsersProfileNotDefined custom error is raised when the %ALLUSERSPROFILE% environment is not defined (Windows only)
var ErrAllUsersProfileNotDefined = errors.New("ALLUSERSPROFILE environment not defined")

// The ErrOptionExists custom error is raised when an option with the same key already exists
var ErrOptionExists = errors.New("option with this key already exists")

// The ErrOptionDoesNotExist custom error is raised when a requested option key does not exist
var ErrOptionDoesNotExist = errors.New("option with this key does not exist")

// The ErrInvalidType custom error is raised when a datum cannot be cast
var ErrInvalidType = errors.New("invalid data type")

// The ErrFlagsAlreadyParsed custom error is raised when flag.Parse() has already been called
var ErrFlagsAlreadyParsed = errors.New("flags have already been parsed")
