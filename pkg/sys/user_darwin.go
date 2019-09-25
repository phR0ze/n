// +build darwin

package sys

import (
	"github.com/pkg/errors"
)

// SwitchUser to the given user/group
// Note the bug around switching uid/gid in linux https://github.com/golang/go/issues/1435
// http://timetobleed.com/5-things-you-dont-know-about-user-ids-that-will-destroy-you/
// requires you drop the group before the user and use a safe solution
func SwitchUser(uid, euid, suid, gid, egid, sgid int) (err error) {
	err = errors.Wrap(err, "Not implemented for darwin")
	return
}
