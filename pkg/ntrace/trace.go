package ntrace

import (
	"path"
	"runtime"

	"github.com/phR0ze/n"
)

// Entry provides a simple struct for trace data
type Entry struct {
	Line int
	Func string
	File string
}

// CallerTrace skips frames until skipfiles[0] is reached e.g. skipping logrus code.
// If no target is given you'll get the full stack.
func CallerTrace(skipframes int, skipfiles ...string) (result []*Entry) {
	more := false
	var frame runtime.Frame
	callers := make([]uintptr, 20)
	runtime.Callers(2, callers)
	frames := runtime.CallersFrames(callers)
	methods := n.S("ntrace.CallerTrace", "ntrace.CallerTraceOne")

	for {
		frame, more = frames.Next()

		// Move up the stack skipping files or methods called out
		if more && (methods.Contains(path.Base(frame.Function)) ||
			skipfiles != nil && n.A(frame.File).ContainsAny(skipfiles...)) {
			continue
		}

		if frame.Function != "" {
			f := path.Base(frame.Function)
			file := n.A(frame.File).Split("/").TakeLastCnt(3).Join("/").A()
			result = append(result, &Entry{Func: f, Line: frame.Line, File: file})
		}
		if !more {
			break
		}

	}

	if skipframes > 0 && len(result) > 0 {
		if skipframes < len(result) {
			result = result[skipframes:]
		} else {
			result = []*Entry{}
		}
	}

	return
}

// CallerTraceOne simply returns the first entry from CallerTrace
func CallerTraceOne(skipframes int, skipfiles ...string) *Entry {
	entries := CallerTrace(skipframes, skipfiles...)
	if entries != nil && len(entries) > 0 {
		return entries[0]
	}
	return &Entry{Func: "<???>", File: "<???>", Line: 0}
}
