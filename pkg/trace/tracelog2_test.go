package trace

func LogFunc3() {
	LogFunc4()
}

func LogFunc4() {
	if one {
		trace = []*Entry{CallerTraceOne(skipfiles...)}
	} else {
		trace = CallerTrace(skipfiles...)
	}
}
