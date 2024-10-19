package code

// Common: basic errors
const (
	// Success - 200: OK.
	Success Code = iota + 100001

	// InternalServer - 500: Internal service error.
	InternalServer

	// Database - 500: Database error.
	Database

	// BadRequest - 400: Bad Request.
	BadRequest

	// PageNotFound - 404: Page not found.
	PageNotFound

	// Validation - 400: Validation failed.
	Validation

	// Bind - 400: Invalid request parameters.
	Bind
)
