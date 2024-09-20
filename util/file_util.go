package util

import (
	"os"
)

// return (isExisted, isDir, err)
func IsDirOrFile(path string) (bool, bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, false, nil
		}
		return false, false, err
	} else {
		return true, info.IsDir(), nil
	}
}
