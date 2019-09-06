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
	if info := opt.Find(opts, "info"); info != nil {
		if val, ok := info.Val.(*FileInfo); ok {
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
	if follow := opt.Find(opts, "follow"); follow != nil {
		if val, ok := follow.Val.(bool); ok {
			result = val
		}
	}
	return
}

// adds a follow option to the options slice if not found with the given value
func defaultFollowOpt(opts *[]*opt.Opt, val bool) {
	if follow := opt.Find(*opts, "follow"); follow == nil {
		*opts = append(*opts, &opt.Opt{Key: "follow", Val: val})
	}
}

// RecurseOpt creates a new recurse option with the given value
// -------------------------------------------------------------------------------------------------
func RecurseOpt(val bool) *opt.Opt {
	return &opt.Opt{Key: "recurse", Val: val}
}

// get the recurse option from the options slice defaulting to false
func getRecurseOpt(opts []*opt.Opt) (result bool) {
	if recurse := opt.Find(opts, "recurse"); recurse != nil {
		if val, ok := recurse.Val.(bool); ok {
			result = val
		}
	}
	return
}
