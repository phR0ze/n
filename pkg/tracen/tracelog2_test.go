package tracen

func LogFunc3() {
	LogFunc4()
}

func LogFunc4() {
	if one {
		trace = []*Entry{CallerTraceOne(0, skipfiles...)}
	} else {
		trace = CallerTrace(0, skipfiles...)
	}
}
