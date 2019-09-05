package test

import (
	"io"
	"os"
	"path/filepath"
	"reflect"
	"time"

	"github.com/bouk/monkey"
)

// ForceTimeNow patches time.Now to return the given value. Once it has been triggered it removes the patch and
// operates as per normal.
// This patch requires the -gcflags=-l to operate correctly e.g. go test -gcflags=-l ./pkg/sys
func ForceTimeNow(val time.Time) (patch *monkey.PatchGuard) {
	patch = monkey.Patch(time.Now, func() time.Time {
		return val
	})
	return
}

// OneShotForceFilePathAbsError patches filepath.Abs to return an error. Once it has been triggered it
// removes the patch and operates as per normal.
// This patch requires the -gcflags=-l to operate correctly e.g. go test -gcflags=-l ./pkg/sys
func OneShotForceFilePathAbsError() {
	var patch *monkey.PatchGuard
	patch = monkey.Patch(filepath.Abs, func(path string) (string, error) {
		patch.Unpatch()
		return "", os.ErrInvalid
	})
}

// OneShotForceFilePathGlobError patches filepath.Chmod to return an error. Once it has been triggered it
// removes the patch and operates as per normal.
// This patch requires the -gcflags=-l to operate correctly e.g. go test -gcflags=-l ./pkg/sys
func OneShotForceFilePathGlobError() {
	var patch *monkey.PatchGuard
	patch = monkey.Patch(filepath.Glob, func(pattern string) ([]string, error) {
		patch.Unpatch()
		return []string{}, os.ErrInvalid
	})
}

// OneShotForceIOCopyError patches io.Copy to return an error. Once it has been triggered it
// removes the patch and operates as per normal.
// This patch requires the -gcflags=-l to operate correctly e.g. go test -gcflags=-l ./pkg/sys
func OneShotForceIOCopyError() {
	var patch *monkey.PatchGuard
	patch = monkey.Patch(io.Copy, func(io.Writer, io.Reader) (int64, error) {
		patch.Unpatch()
		return 0, os.ErrInvalid
	})
}

// OneShotForceOSCloseError patches *os.File.Close to Close the file then return an error. Once it has been triggered it
// removes the patch and operates as per normal.
// This patch requires the -gcflags=-l to operate correctly e.g. go test -gcflags=-l ./pkg/sys
func OneShotForceOSCloseError() {
	var patch *monkey.PatchGuard
	patch = monkey.PatchInstanceMethod(reflect.TypeOf((*os.File)(nil)), "Close", func(f *os.File) error {
		patch.Unpatch()
		f.Close()
		return os.ErrInvalid
	})
}

// OneShotForceOSChmodError patches os.Chmod to return an error. Once it has been triggered it
// removes the patch and operates as per normal.
// This patch requires the -gcflags=-l to operate correctly e.g. go test -gcflags=-l ./pkg/sys
func OneShotForceOSChmodError() {
	var patch *monkey.PatchGuard
	patch = monkey.Patch(os.Chmod, func(dstPath string, mode os.FileMode) error {
		patch.Unpatch()
		return os.ErrInvalid
	})
}

// OneShotForceOSChtimesError patches os.Chtimesto return an error. Once it has been triggered it
// removes the patch and operates as per normal.
// This patch requires the -gcflags=-l to operate correctly e.g. go test -gcflags=-l ./pkg/sys
func OneShotForceOSChtimesError() {
	var patch *monkey.PatchGuard
	patch = monkey.Patch(os.Chtimes, func(name string, _, _ time.Time) error {
		patch.Unpatch()
		return os.ErrInvalid
	})
}

// OneShotForceOSCreateError patches os.Create to return an error. Once it has been triggered it
// removes the patch and operates as per normal.
// This patch requires the -gcflags=-l to operate correctly e.g. go test -gcflags=-l ./pkg/sys
func OneShotForceOSCreateError() {
	var patch *monkey.PatchGuard
	patch = monkey.Patch(os.Create, func(name string) (*os.File, error) {
		patch.Unpatch()
		return nil, os.ErrInvalid
	})
}

// OneShotForceOSReaddirnamesError patches *os.File.Readdirnames to return an error. Once it has been triggered it
// removes the patch and operates as per normal.
// This patch requires the -gcflags=-l to operate correctly e.g. go test -gcflags=-l ./pkg/sys
func OneShotForceOSReaddirnamesError() {
	var patch *monkey.PatchGuard
	patch = monkey.PatchInstanceMethod(reflect.TypeOf((*os.File)(nil)), "Readdirnames", func(*os.File, int) ([]string, error) {
		patch.Unpatch()
		return []string{}, os.ErrInvalid
	})
}

// OneShotForceOSReadlinkError patches os.Readlink to return an error. Once it has been triggered it
// removes the patch and operates as per normal.
// This patch requires the -gcflags=-l to operate correctly e.g. go test -gcflags=-l ./pkg/sys
func OneShotForceOSReadlinkError() {
	var patch *monkey.PatchGuard
	patch = monkey.Patch(os.Readlink, func(name string) (string, error) {
		patch.Unpatch()
		return "", os.ErrInvalid
	})
}

// OneShotForceOSSyncError patches *os.File.Sync to Sync the file then return an error. Once it has been triggered it
// removes the patch and operates as per normal.
// This patch requires the -gcflags=-l to operate correctly e.g. go test -gcflags=-l ./pkg/sys
func OneShotForceOSSyncError() {
	var patch *monkey.PatchGuard
	patch = monkey.PatchInstanceMethod(reflect.TypeOf((*os.File)(nil)), "Sync", func(f *os.File) error {
		patch.Unpatch()
		f.Sync()
		return os.ErrInvalid
	})
}
