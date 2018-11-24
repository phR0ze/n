package nub

//--------------------------------------------------------------------------------------------------
// StrMap Nub
//--------------------------------------------------------------------------------------------------
type strMapNub struct {
	raw map[string]interface{}
}

// StrMap creates a new nub from the given string map
func StrMap(other map[string]interface{}) *strMapNub {
	return &strMapNub{raw: other}
}

// Ex processes any deferred execution and returns the underlying map
func (m *strMapNub) Ex() map[string]interface{} {
	return m.raw
}
