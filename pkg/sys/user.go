package sys

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"syscall"

	"github.com/pkg/errors"
)

// User wraps the os.User interface and provide additional helper functions
type User struct {
	obj      *user.User // handle on the actual OS object to use where needed
	UID      int        // Uid is the user ID, will be 0 when sudo is used.
	GID      int        // Gid is the primary group ID, will be 0 when sudo is used.
	Name     string     // Username is the login name, will be root when sudo is used.
	Home     string     // HomeDir is the path to the user's home directory (if they have one).
	RealUID  int        // Real user uid behind the sudo mask
	RealGID  int        // Real user uid behind the sudo mask
	RealName string     // Real user name behind the sudo mask
	RealHome string     // Real user home behind the sudo mask
}

// convert an *user.User to an internal *User
func newUserStruct(obj *user.User) (u *User, err error) {
	if obj == nil {
		err = errors.Wrap(err, "user object is nil")
		return
	}

	// Convert user's uid and gid to integers
	var uid, gid int64
	if uid, err = strconv.ParseInt(obj.Uid, 0, 0); err != nil {
		err = errors.Wrap(err, "failed to convert user's uid into an int. Not a POSIX system?")
		return
	}
	if gid, err = strconv.ParseInt(obj.Gid, 0, 0); err != nil {
		err = errors.Wrap(err, "failed to convert user's gid into an int. Not a POSIX system?")
		return
	}

	// Unmask the sudo user if using sudo
	suid, sgid := -1, -1
	var realName, realHome string
	if uid == 0 {

		// We're actually root so set accordingly
		if id := os.Getenv("SUDO_UID"); id == "" {
			suid, sgid = int(uid), int(gid)
			realName = obj.Username
			realHome = obj.HomeDir
		} else {
			var _suser *user.User
			if _suser, err = user.LookupId(id); err != nil {
				err = errors.Wrap(err, "failed to get real user behind sudo mask")
				return
			}
			var suser *User
			if suser, err = newUserStruct(_suser); err != nil {
				return
			}
			suid, sgid = suser.UID, suser.GID
			realName = suser.Name
			realHome = suser.Home
		}
	}

	// Create the new User struct
	u = &User{
		obj:      obj,
		UID:      int(uid),
		GID:      int(gid),
		Name:     obj.Username,
		Home:     obj.HomeDir,
		RealUID:  suid,
		RealGID:  sgid,
		RealName: realName,
		RealHome: realHome,
	}
	return
}

// CurrentUser gets the current system user. The internal user.Current() caches this result.
func CurrentUser(cache ...bool) (u *User, err error) {
	_cache := true
	if len(cache) > 0 {
		_cache = cache[0]
	}

	// Get the current system user using the cached path
	var obj *user.User
	if _cache {
		if obj, err = user.Current(); err != nil {
			err = errors.Wrap(err, "failed to get current user")
			return
		}

		// Convert to internal object
		if u, err = newUserStruct(obj); err != nil {
			return
		}
	} else {
		// Get user using a non cached path
		if u, err = LookupUserById(os.Getuid()); err != nil {
			err = errors.Wrapf(err, "failed to get current user")
			return
		}
	}

	return
}

// LookupUserById gets the user for the given id
func LookupUserById(uid int) (u *User, err error) {

	// Convert the uid into a string
	id := fmt.Sprintf("%d", uid)

	// Lookup the given user by id
	var obj *user.User
	if obj, err = user.LookupId(id); err != nil {
		err = errors.Wrapf(err, "failed to get user for id %d", uid)
		return
	}

	// Convert to internal object
	if u, err = newUserStruct(obj); err != nil {
		return
	}

	return
}

// IsRoot detects if the current user has root permissions based on the user uid
func (u *User) IsRoot() bool {
	if u == nil {
		return false
	}
	return u.UID == 0
}

// DropSudo switches back to the original user under the sudo mask.
// Preserves the ability to raise Sudo again.
func DropSudo() (err error) {
	if u, e := CurrentUser(); e == nil && u.IsRoot() {
		uid, gid := u.RealUID, u.RealGID
		if err = SwitchUser(uid, uid, 0, gid, gid, 0); err != nil {
			return
		}
	}
	return
}

// DropSudoP switches back to the original user under the sudo mask.
// Does not preserve the ability to raise Sudo again.
func DropSudoP() (err error) {
	if u, e := CurrentUser(); e == nil && u.IsRoot() {
		uid, gid := u.RealUID, u.RealGID
		if err = SwitchUser(uid, uid, uid, gid, gid, gid); err != nil {
			return
		}
	}
	return
}

// Sudo switches back to sudo root. Returns an error if not allowed
func Sudo() (err error) {
	if err = SwitchUser(0, 0, 0, 0, 0, 0); err != nil {
		return
	}
	return
}

// SwitchUser to the given user/group
// Note the bug around switching uid/gid in linux https://github.com/golang/go/issues/1435
// http://timetobleed.com/5-things-you-dont-know-about-user-ids-that-will-destroy-you/
// requires you drop the group before the user and use a safe solution
func SwitchUser(uid, euid, suid, gid, egid, sgid int) (err error) {
	if err = syscall.Setresgid(gid, egid, sgid); err != nil {
		err = errors.Wrap(err, "failed to set gid while switching user")
		return
	}
	if err = syscall.Setresuid(uid, euid, suid); err != nil {
		err = errors.Wrap(err, "failed to set uid while switching user")
		return
	}
	return
}

// UserHome returns the absolute home directory for the current user.
// os.UserHomeDir is used internall because as of Go 1.9 cgo is no longer required
func UserHome() (result string, err error) {
	if result, err = os.UserHomeDir(); err != nil {
		err = errors.Wrap(err, "failed to compute the user's home directory")
		return
	}

	// Now ensure we have an absolute value here, trimming out all redirects
	if result, err = filepath.Abs(result); err != nil {
		err = errors.Wrapf(err, "failed to compute the absolute path for %s", result)
		return
	}
	return
}

// UserIsRoot detects if the current user has root permissions based on the user uid
func UserIsRoot() (root bool) {
	return os.Getuid() == 0
}

// UserIsRealRoot detects if the current user is the root user not by SUDO
func UserIsRealRoot() (root bool) {
	if u, err := CurrentUser(); err == nil {
		return u.RealGID == 0
	}
	return false
}

// UserUID is the currrent user's uid
func UserUID() int {
	return os.Getuid()
}

// UserGID is the currrent user's gid
func UserGID() int {
	return os.Getgid()
}

// UserName is the currrent user's name
func UserName() string {
	if u, err := CurrentUser(); err == nil {
		return u.Name
	}
	return ""
}

// UserRealUID as opposed to UID is the current user's uid behind the sudo mask
func UserRealUID() int {
	if u, err := CurrentUser(); err == nil {
		return u.RealUID
	}
	return -1
}

// UserRealGID as opposed to GID is the current user's gid behind the sudo mask
func UserRealGID() int {
	if u, err := CurrentUser(); err == nil {
		return u.RealGID
	}
	return -1
}

// UserRealName as opposed to Name is the current user's name behind the sudo mask
func UserRealName() string {
	if u, err := CurrentUser(); err == nil {
		return u.RealName
	}
	return ""
}

// UserRealHome as opposed to UserHome is the current user's home behind the sudo mask
func UserRealHome() string {
	if u, err := CurrentUser(); err == nil {
		return u.RealHome
	}
	return ""
}
