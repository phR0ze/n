// Package errs provides a common set of error for the pkg n
package errs

import "fmt"

// ErrorType is an enumeration of well known error types
type ErrorType string

const (
	// ErrorTypeTmplEndTagNotFound indicates that a start tag was found but not an end tag or a template variable
	ErrorTypeTmplEndTagNotFound ErrorType = "TmplEndTagNotFound"

	// ErrorTypeTmplTagsInvalid indicates that the template start and/or end tags were invalid
	ErrorTypeTmplTagsInvalid ErrorType = "TmplTagsInvalid"

	// ErrorTypeTmplVarsNotFound indicates that template variables were not found
	ErrorTypeTmplVarsNotFound ErrorType = "TmplVarsNotFound"
)

// Error provides a common error type for N types
type Error struct {
	Message string
	Type    ErrorType
}

func (e Error) Error() string {
	return e.Message
}

// IsTmplEndTagNotFound returns true if the given err was created by NewTmplEndTagNotFound
func IsTmplEndTagNotFound(err error) bool {
	if nerr, ok := err.(Error); ok {
		return nerr.Type == ErrorTypeTmplEndTagNotFound
	}
	return false
}

// IsTmplTagsInvalid returns true if the given err was created by NewTmplTagsInvalid
func IsTmplTagsInvalid(err error) bool {
	if nerr, ok := err.(Error); ok {
		return nerr.Type == ErrorTypeTmplTagsInvalid
	}
	return false
}

// IsTmplVarsNotFound returns true if the given err was created by NewTmplVarsNotFound
func IsTmplVarsNotFound(err error) bool {
	if nerr, ok := err.(Error); ok {
		return nerr.Type == ErrorTypeTmplVarsNotFound
	}
	return false
}

// NewTmplEndTagNotFound indicates that a start tag was found but not an end tag or a template variable
func NewTmplEndTagNotFound(endTag string, starting []byte) Error {
	return Error{
		Message: fmt.Sprintf("cannot find end tag=%s starting from %s", endTag, starting),
		Type:    ErrorTypeTmplEndTagNotFound,
	}
}

// NewTmplTagsInvalid indicates that the template start and/or end tags were invalid
func NewTmplTagsInvalid() Error {
	return Error{
		Message: "start and/or end tag are invalid",
		Type:    ErrorTypeTmplTagsInvalid,
	}
}

// NewTmplVarsNotFound indicates that template variables were not found
func NewTmplVarsNotFound() Error {
	return Error{
		Message: "template variables were not found",
		Type:    ErrorTypeTmplVarsNotFound,
	}
}
