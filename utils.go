package gokindlebt

import (
	"os/user"
	"strconv"
	"syscall"
)

func UseBluetoothPrivileges() error {
	group, err := user.LookupGroup("bluetooth")
	if err != nil {
		return err
	}
	user, err := user.Lookup("bluetooth")
	if err != nil {
		return err
	}

	groupId, err := strconv.Atoi(group.Gid)
	if err != nil {
		return err
	}
	userId, err := strconv.Atoi(user.Uid)
	if err != nil {
		return err
	}

	err = syscall.Setgid(groupId)
	if err != nil {
		return err
	}
	err = syscall.Setuid(userId)
	if err != nil {
		return err
	}

	return nil
}

func IsASCIIPrintable(data []byte) bool {
	for _, b := range data {
		if b < 0x20 || b > 0x7E {
			return false
		}
	}
	return true
}
