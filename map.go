package nub

//--------------------------------------------------------------------------------------------------
// StrMap Nub
//--------------------------------------------------------------------------------------------------
type strMapNub struct {
	Raw map[string]interface{}
}

// StrMap creates a new nub from the given string map
func StrMap(other map[string]interface{}) *strMapNub {
	return &strMapNub{Raw: other}
}
