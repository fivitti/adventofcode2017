package fileutils

import "os"

func FileExists(path string) error {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return err
	}
	return nil
}

