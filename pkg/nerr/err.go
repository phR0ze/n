// Package nerr provides a common set of error for N types
package nerr

import "fmt"

// NErrorType is an enumeration of well known error types
type NErrorType string

const (
	// NErrorTypeTmplEndTagNotFound indicates that a start tag was found but not an end tag or a template variable
	NErrorTypeTmplEndTagNotFound NErrorType = "TmplEndTagNotFound"

	// NErrorTypeTmplTagsInvalid indicates that the template start and/or end tags were invalid
	NErrorTypeTmplTagsInvalid NErrorType = "TmplTagsInvalid"

	// NErrorTypeTmplVarsNotFound indicates that template variables were not found
	NErrorTypeTmplVarsNotFound NErrorType = "TmplVarsNotFound"
)

// NError provides a common error type for N types
type NError struct {
	Message string
	Type    NErrorType
}

func (e NError) Error() string {
	return e.Message
}

// IsTmplEndTagNotFound returns true if the given err was created by NewTmplEndTagNotFound
func IsTmplEndTagNotFound(err error) bool {
	if nerr, ok := err.(NError); ok {
		return nerr.Type == NErrorTypeTmplEndTagNotFound
	}
	return false
}

// IsTmplTagsInvalid returns true if the given err was created by NewTmplTagsInvalid
func IsTmplTagsInvalid(err error) bool {
	if nerr, ok := err.(NError); ok {
		return nerr.Type == NErrorTypeTmplTagsInvalid
	}
	return false
}

// IsTmplVarsNotFound returns true if the given err was created by NewTmplVarsNotFound
func IsTmplVarsNotFound(err error) bool {
	if nerr, ok := err.(NError); ok {
		return nerr.Type == NErrorTypeTmplVarsNotFound
	}
	return false
}

// NewTmplEndTagNotFound indicates that a start tag was found but not an end tag or a template variable
func NewTmplEndTagNotFound(endTag string, starting []byte) NError {
	return NError{
		Message: fmt.Sprintf("cannot find end tag=%s starting from %s", endTag, starting),
		Type:    NErrorTypeTmplEndTagNotFound,
	}
}

// NewTmplTagsInvalid indicates that the template start and/or end tags were invalid
func NewTmplTagsInvalid() NError {
	return NError{
		Message: "start and/or end tag are invalid",
		Type:    NErrorTypeTmplTagsInvalid,
	}
}

// NewTmplVarsNotFound indicates that template variables were not found
func NewTmplVarsNotFound() NError {
	return NError{
		Message: "template variables were not found",
		Type:    NErrorTypeTmplVarsNotFound,
	}
}
