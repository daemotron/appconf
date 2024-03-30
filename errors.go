package appconf

import "errors"

// The ErrAllUsersProfileNotDefined custom error is raised when the %ALLUSERSPROFILE% environment is not defined (Windows only)
var ErrAllUsersProfileNotDefined = errors.New("ALLUSERSPROFILE environment not defined")
