package sys

import (
	"github.com/phR0ze/n/pkg/opt"
)

// InfoOpt creates a new option info option with the given *FileInfo
// -------------------------------------------------------------------------------------------------
func InfoOpt(val *FileInfo) *opt.Opt {
	return &opt.Opt{Key: "info", Val: val}
}

// get the info option from the options slice defaulting to nil
func getInfoOpt(opts []*opt.Opt) *FileInfo {
	if o := opt.Get(opts, "info"); o != nil {
		if val, ok := o.Val.(*FileInfo); ok {
			return val
		}
	}
	return nil
}

// FollowOpt creates a new follow option with the given value
// -------------------------------------------------------------------------------------------------
func FollowOpt(val bool) *opt.Opt {
	return &opt.Opt{Key: "follow", Val: val}
}

// get the follow option from the options slice defaulting to false
func getFollowOpt(opts []*opt.Opt) (result bool) {
	if o := opt.Get(opts, "follow"); o != nil {
		if val, ok := o.Val.(bool); ok {
			result = val
		}
	}
	return
}

// adds a follow option to the options slice if not found with the given value
func defaultFollowOpt(opts *[]*opt.Opt, val bool) {
	if o := opt.Get(*opts, "follow"); o == nil {
		*opts = append(*opts, &opt.Opt{Key: "follow", Val: val})
	}
}

// OnlyDirsOpt creates a new only dirs option with the given value
// -------------------------------------------------------------------------------------------------
func OnlyDirsOpt(val bool) *opt.Opt {
	return &opt.Opt{Key: "onlyDirs", Val: val}
}

// get the only dirs option from the options slice defaulting to false
func getOnlyDirsOpt(opts []*opt.Opt) (result bool) {
	if o := opt.Get(opts, "onlyDirs"); o != nil {
		if val, ok := o.Val.(bool); ok {
			result = val
		}
	}
	return
}

// OnlyFilesOpt creates a new only dirs option with the given value
// -------------------------------------------------------------------------------------------------
func OnlyFilesOpt(val bool) *opt.Opt {
	return &opt.Opt{Key: "onlyFiles", Val: val}
}

// get the only dirs option from the options slice defaulting to false
func getOnlyFilesOpt(opts []*opt.Opt) (result bool) {
	if o := opt.Get(opts, "onlyFiles"); o != nil {
		if val, ok := o.Val.(bool); ok {
			result = val
		}
	}
	return
}

// RecurseOpt creates a new recurse option with the given value
// -------------------------------------------------------------------------------------------------
func RecurseOpt(val bool) *opt.Opt {
	return &opt.Opt{Key: "recurse", Val: val}
}

// get the recurse option from the options slice defaulting to false
func getRecurseOpt(opts []*opt.Opt) (result bool) {
	if o := opt.Get(opts, "recurse"); o != nil {
		if val, ok := o.Val.(bool); ok {
			result = val
		}
	}
	return
}

// RootOpt creates a new root option with the given value
// -------------------------------------------------------------------------------------------------
func RootOpt(val bool) *opt.Opt {
	return &opt.Opt{Key: "root", Val: val}
}

// get the root option from the options slice defaulting to false
func getRootOpt(opts []*opt.Opt) (result bool) {
	if o := opt.Get(opts, "root"); o != nil {
		if val, ok := o.Val.(bool); ok {
			result = val
		}
	}
	return
}
