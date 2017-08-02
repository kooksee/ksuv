package app

import "os"

func IsDir(dir string) bool {
	fi, err := os.Stat(dir)
	return err == nil && fi.IsDir()
}
