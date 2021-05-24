package util

import "os"

// path exists or not
func IsFileExist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	if _, err := os.Stat(path); err == nil {
		return true
	}
	return false
}
