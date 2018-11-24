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

// M materializes object invoking deferred execution
func (m *strMapNub) M() map[string]interface{} {
	return m.raw
}
