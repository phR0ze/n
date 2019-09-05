package test

import (
	"io"
	"os"
	"path/filepath"
	"reflect"
	"time"

	"github.com/bouk/monkey"
)

// ForceTimeNow patches time.Now to return the given value. Removing the patch is up to the caller.
// This patch requires the -gcflags=-l to operate correctly e.g. go test -gcflags=-l ./pkg/sys
func ForceTimeNow(val time.Time) (patch *monkey.PatchGuard) {
	patch = monkey.Patch(time.Now, func() time.Time {
		return val
	})
	return
}

// ForceOSCloseError patches *os.File.Close to Close the file then return an error. Removing the patch is up to the caller.
// This patch requires the -gcflags=-l to operate correctly e.g. go test -gcflags=-l ./pkg/sys
func ForceOSCloseError(errs ...error) (patch *monkey.PatchGuard) {
	err := os.ErrInvalid
	if len(errs) > 0 {
		err = errs[0]
	}
	patch = monkey.PatchInstanceMethod(reflect.TypeOf((*os.File)(nil)), "Close", func(f *os.File) error {
		patch.Unpatch()
		f.Close()
		patch.Restore()
		return err
	})
	return
}

// OneShotForceFilePathAbsError patches filepath.Abs to return an error. Once it has been triggered it
// removes the patch and operates as per normal.
// This patch requires the -gcflags=-l to operate correctly e.g. go test -gcflags=-l ./pkg/sys
func OneShotForceFilePathAbsError(errs ...error) {
	err := os.ErrInvalid
	if len(errs) > 0 {
		err = errs[0]
	}
	var patch *monkey.PatchGuard
	patch = monkey.Patch(filepath.Abs, func(path string) (string, error) {
		patch.Unpatch()
		return "", err
	})
}

// OneShotForceFilePathGlobError patches filepath.Chmod to return an error. Once it has been triggered it
// removes the patch and operates as per normal.
// This patch requires the -gcflags=-l to operate correctly e.g. go test -gcflags=-l ./pkg/sys
func OneShotForceFilePathGlobError(errs ...error) {
	err := os.ErrInvalid
	if len(errs) > 0 {
		err = errs[0]
	}
	var patch *monkey.PatchGuard
	patch = monkey.Patch(filepath.Glob, func(pattern string) ([]string, error) {
		patch.Unpatch()
		return []string{}, err
	})
}

// OneShotForceIOCopyError patches io.Copy to return an error. Once it has been triggered it
// removes the patch and operates as per normal.
// This patch requires the -gcflags=-l to operate correctly e.g. go test -gcflags=-l ./pkg/sys
func OneShotForceIOCopyError(errs ...error) {
	err := os.ErrInvalid
	if len(errs) > 0 {
		err = errs[0]
	}
	var patch *monkey.PatchGuard
	patch = monkey.Patch(io.Copy, func(io.Writer, io.Reader) (int64, error) {
		patch.Unpatch()
		return 0, err
	})
}

// OneShotForceOSCloseError patches *os.File.Close to Close the file then return an error. Once it has been triggered it
// removes the patch and operates as per normal.
// This patch requires the -gcflags=-l to operate correctly e.g. go test -gcflags=-l ./pkg/sys
func OneShotForceOSCloseError(errs ...error) {
	err := os.ErrInvalid
	if len(errs) > 0 {
		err = errs[0]
	}
	var patch *monkey.PatchGuard
	patch = monkey.PatchInstanceMethod(reflect.TypeOf((*os.File)(nil)), "Close", func(f *os.File) error {
		patch.Unpatch()
		f.Close()
		return err
	})
}

// OneShotForceOSChmodError patches os.Chmod to return an error. Once it has been triggered it
// removes the patch and operates as per normal.
// This patch requires the -gcflags=-l to operate correctly e.g. go test -gcflags=-l ./pkg/sys
func OneShotForceOSChmodError(errs ...error) {
	err := os.ErrInvalid
	if len(errs) > 0 {
		err = errs[0]
	}
	var patch *monkey.PatchGuard
	patch = monkey.Patch(os.Chmod, func(dstPath string, mode os.FileMode) error {
		patch.Unpatch()
		return err
	})
}

// OneShotForceOSChtimesError patches os.Chtimesto return an error. Once it has been triggered it
// removes the patch and operates as per normal.
// This patch requires the -gcflags=-l to operate correctly e.g. go test -gcflags=-l ./pkg/sys
func OneShotForceOSChtimesError(errs ...error) {
	err := os.ErrInvalid
	if len(errs) > 0 {
		err = errs[0]
	}
	var patch *monkey.PatchGuard
	patch = monkey.Patch(os.Chtimes, func(name string, _, _ time.Time) error {
		patch.Unpatch()
		return err
	})
}

// OneShotForceOSCreateError patches os.Create to return an error. Once it has been triggered it
// removes the patch and operates as per normal.
// This patch requires the -gcflags=-l to operate correctly e.g. go test -gcflags=-l ./pkg/sys
func OneShotForceOSCreateError(errs ...error) {
	err := os.ErrInvalid
	if len(errs) > 0 {
		err = errs[0]
	}
	var patch *monkey.PatchGuard
	patch = monkey.Patch(os.Create, func(name string) (*os.File, error) {
		patch.Unpatch()
		return nil, err
	})
}

// OneShotForceOSReadError patches *os.File.Read to return an error. Once it has been triggered it
// removes the patch and operates as per normal.
// This patch requires the -gcflags=-l to operate correctly e.g. go test -gcflags=-l ./pkg/sys
func OneShotForceOSReadError(errs ...error) {
	err := os.ErrInvalid
	if len(errs) > 0 {
		err = errs[0]
	}
	var patch *monkey.PatchGuard
	patch = monkey.PatchInstanceMethod(reflect.TypeOf((*os.File)(nil)), "Read", func(*os.File, []byte) (int, error) {
		patch.Unpatch()
		return 0, err
	})
}

// OneShotForceOSReadAtError patches *os.File.ReadAt to return an error. Once it has been triggered it
// removes the patch and operates as per normal.
// This patch requires the -gcflags=-l to operate correctly e.g. go test -gcflags=-l ./pkg/sys
func OneShotForceOSReadAtError(errs ...error) {
	err := os.ErrInvalid
	if len(errs) > 0 {
		err = errs[0]
	}
	var patch *monkey.PatchGuard
	patch = monkey.PatchInstanceMethod(reflect.TypeOf((*os.File)(nil)), "ReadAt", func(*os.File, []byte, int64) (int, error) {
		patch.Unpatch()
		return 0, err
	})
}

// OneShotForceOSReaddirnamesError patches *os.File.Readdirnames to return an error. Once it has been triggered it
// removes the patch and operates as per normal.
// This patch requires the -gcflags=-l to operate correctly e.g. go test -gcflags=-l ./pkg/sys
func OneShotForceOSReaddirnamesError(errs ...error) {
	err := os.ErrInvalid
	if len(errs) > 0 {
		err = errs[0]
	}
	var patch *monkey.PatchGuard
	patch = monkey.PatchInstanceMethod(reflect.TypeOf((*os.File)(nil)), "Readdirnames", func(*os.File, int) ([]string, error) {
		patch.Unpatch()
		return []string{}, err
	})
}

// OneShotForceOSReadlinkError patches os.Readlink to return an error. Once it has been triggered it
// removes the patch and operates as per normal.
// This patch requires the -gcflags=-l to operate correctly e.g. go test -gcflags=-l ./pkg/sys
func OneShotForceOSReadlinkError(errs ...error) {
	err := os.ErrInvalid
	if len(errs) > 0 {
		err = errs[0]
	}
	var patch *monkey.PatchGuard
	patch = monkey.Patch(os.Readlink, func(name string) (string, error) {
		patch.Unpatch()
		return "", err
	})
}

// OneShotForceOSSyncError patches *os.File.Sync to Sync the file then return an error. Once it has been triggered it
// removes the patch and operates as per normal.
// This patch requires the -gcflags=-l to operate correctly e.g. go test -gcflags=-l ./pkg/sys
func OneShotForceOSSyncError(errs ...error) {
	err := os.ErrInvalid
	if len(errs) > 0 {
		err = errs[0]
	}
	var patch *monkey.PatchGuard
	patch = monkey.PatchInstanceMethod(reflect.TypeOf((*os.File)(nil)), "Sync", func(f *os.File) error {
		patch.Unpatch()
		f.Sync()
		return err
	})
}

// OneShotForceOSTruncateError patches *os.File.Truncate to return an error. Once it has been triggered it
// removes the patch and operates as per normal.
// This patch requires the -gcflags=-l to operate correctly e.g. go test -gcflags=-l ./pkg/sys
func OneShotForceOSTruncateError(errs ...error) {
	err := os.ErrInvalid
	if len(errs) > 0 {
		err = errs[0]
	}
	var patch *monkey.PatchGuard
	patch = monkey.PatchInstanceMethod(reflect.TypeOf((*os.File)(nil)), "Truncate", func(*os.File, int64) error {
		patch.Unpatch()
		return err
	})
}

// OneShotForceOSWriteAtError patches *os.File.WriteAt to return an error. Once it has been triggered it
// removes the patch and operates as per normal.
// This patch requires the -gcflags=-l to operate correctly e.g. go test -gcflags=-l ./pkg/sys
func OneShotForceOSWriteAtError(errs ...error) {
	err := os.ErrInvalid
	if len(errs) > 0 {
		err = errs[0]
	}
	var patch *monkey.PatchGuard
	patch = monkey.PatchInstanceMethod(reflect.TypeOf((*os.File)(nil)), "WriteAt", func(*os.File, []byte, int64) (int, error) {
		patch.Unpatch()
		return 0, err
	})
}
