package errors

import (
	"errors"
)

var (
	NotFoundError                 = errors.New("Not found!")
	BadRequestError               = errors.New("Bad request!")
	NilValidationFailedError      = errors.New("Validation failed! Mandatory fields missing!")
	SemanticValidationFailedError = errors.New("Validation failed! Semantic validation failed!")
	InternalServerError           = errors.New("Internal server error!")
	ConflictError                 = errors.New("Conflict!")
	UnAuthorizedError             = errors.New("Unathorized!")
)
