package trace

import (
	"path"
	"runtime"

	"github.com/phR0ze/nub"
)

// Entry provides a simple struct for trace data
type Entry struct {
	Line int
	Func string
	File string
}

// CallerTrace skips frames until skipfiles[0] is reached e.g. skipping logrus code.
// If no target is given you'll get the full stack.
func CallerTrace(skipfiles ...string) (result []*Entry) {
	more := false
	var frame runtime.Frame
	callers := make([]uintptr, 20)
	runtime.Callers(2, callers)
	frames := runtime.CallersFrames(callers)
	methods := nub.StrSlice([]string{"trace.CallerTrace", "trace.CallerTraceOne"})

	for {
		frame, more = frames.Next()

		// Move up the stack skipping files or methods called out
		if more && (methods.Contains(path.Base(frame.Function)) ||
			skipfiles != nil && nub.Str(frame.File).ContainsAny(skipfiles)) {
			continue
		}

		if frame.Function != "" {
			f := path.Base(frame.Function)
			file := nub.Str(frame.File).Split("/").TakeLastCnt(3).Join("/").Ex()
			result = append(result, &Entry{Func: f, Line: frame.Line, File: file})
		}
		if !more {
			break
		}

	}
	return
}

// CallerTraceOne simply returns the first entry from CallerTrace
func CallerTraceOne(skipfiles ...string) *Entry {
	entries := CallerTrace(skipfiles...)
	if entries != nil && len(entries) > 0 {
		return entries[0]
	}
	return &Entry{Func: "<???>", File: "<???>", Line: 0}
}
