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

// checks to see if the follow is flag is set
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
