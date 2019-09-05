package sys

import (
	"os"
	"reflect"

	"github.com/bouk/monkey"
)

// Patches *io.File.Close to Close the file then return an error. One it has been triggered it
// removes the patch and operatores as per normal requires the -gcflags=-l to operate correctly
// e.g. go test -gcflags=-l ./pkg/sys
func OneShotForceIOCloseError() {
	var guard *monkey.PatchGuard
	guard = monkey.PatchInstanceMethod(reflect.TypeOf((*os.File)(nil)), "Close", func(f *os.File) error {
		guard.Unpatch()
		if err := f.Close(); err != nil {
			return err
		}
		return os.ErrInvalid
	})
}

// Patches *io.File.Sync to Sync the file then return an error. One it has been triggered it
// removes the patch and operatores as per normal requires the -gcflags=-l to operate correctly
// e.g. go test -gcflags=-l ./pkg/sys
func OneShotForceIOSyncError() {
	var guard *monkey.PatchGuard
	guard = monkey.PatchInstanceMethod(reflect.TypeOf((*os.File)(nil)), "Sync", func(f *os.File) error {
		guard.Unpatch()
		if err := f.Sync(); err != nil {
			return err
		}
		return os.ErrInvalid
	})
}
