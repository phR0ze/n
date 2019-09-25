// +build linux

package sys

import (
	"github.com/pkg/errors"
	"golang.org/x/sys/unix"
)

// SwitchUser to the given user/group
// Note the bug around switching uid/gid in linux https://github.com/golang/go/issues/1435
// http://timetobleed.com/5-things-you-dont-know-about-user-ids-that-will-destroy-you/
// requires you drop the group before the user and use a safe solution
func SwitchUser(uid, euid, suid, gid, egid, sgid int) (err error) {
	if err = unix.Setresgid(gid, egid, sgid); err != nil {
		err = errors.Wrap(err, "failed to set gid while switching user")
		return
	}
	if err = unix.Setresuid(uid, euid, suid); err != nil {
		err = errors.Wrap(err, "failed to set uid while switching user")
		return
	}
	return
}
