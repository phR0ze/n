package sys

import (
	"github.com/phR0ze/n/pkg/opt"
)

// sets the follow opt by default i.e. if not already set
func defaultFollowOpt(opts *[]*opt.Opt, val bool) {
	if follow := opt.Find(*opts, "follow"); follow == nil {
		*opts = append(*opts, &opt.Opt{Key: "follow", Val: val})
	}
}

// get the info option value
func infoOpt(opts []*opt.Opt) *FileInfo {
	if info := opt.Find(opts, "info"); info != nil {
		if val, ok := info.Val.(*FileInfo); ok {
			return val
		}
	}
	return nil
}

// get the follow option value
func followOpt(opts []*opt.Opt) (result bool) {
	if follow := opt.Find(opts, "follow"); follow != nil {
		if val, ok := follow.Val.(bool); ok {
			result = val
		}
	}
	return
}

// create the follow opt with the given value
func newFollowOpt(val bool) *opt.Opt {
	return &opt.Opt{Key: "follow", Val: val}
}

// create a new option containing a file info
func newInfoOpt(info *FileInfo) *opt.Opt {
	return &opt.Opt{Key: "info", Val: info}
}
