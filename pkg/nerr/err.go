// Package nerr provides a common set of error for N types
package nerr

// NErrorType is an enumeration of well known error types
type NErrorType string

const (
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

// IsTmplTagsInvalid returns true if the given err was created by NewInvalidTmplTag
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
