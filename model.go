package n

// Opt provides a mechanism for passing in optional options to functions.
// Value can be anything the accepting function documents.
type Opt struct {
	Name  string      // name of the option
	Value interface{} // value of the option
}
