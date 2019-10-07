// Package errs provides a common set of error for the pkg n
package errs

import (
	"errors"
	"fmt"
)

var (

	// EOF indicates the end of a file equal to io.EOF
	EOF = errors.New("EOF")

	// BOF indicates the begining of a file
	BOF = errors.New("BOF")

	// TmplVarsNotFoundError indicates that template variables were not found
	TmplVarsNotFoundError = Error{Message: "template variables were not found", Type: ErrorTypeTmplVarsNotFound}

	// TmplTagsInvalidError indicates that the template start and/or end tags were invalid
	TmplTagsInvalidError = Error{Message: "start and/or end tag are invalid", Type: ErrorTypeTmplTagsInvalid}
)

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

// Tmpl Erros
//--------------------------------------------------------------------------------------------------

// TmplEndTagNotFoundError returns true if the given err was created by NewTmplEndTagNotFound
func TmplEndTagNotFoundError(err error) bool {
	if e, ok := err.(Error); ok {
		return e.Type == ErrorTypeTmplEndTagNotFound
	}
	return false
}

// NewTmplEndTagNotFoundError indicates that a start tag was found but not an end tag or a template variable
func NewTmplEndTagNotFoundError(endTag string, starting []byte) Error {
	return Error{
		Message: fmt.Sprintf("cannot find end tag=%s starting from %s", endTag, starting),
		Type:    ErrorTypeTmplEndTagNotFound,
	}
}
