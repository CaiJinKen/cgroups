package cgroups

import "os"

func checkPath(p string) error {
	if !exists(p) {
		return os.MkdirAll(p, 755)
	}
	return nil
}

func exists(f string) bool {
	_, err := os.Stat(f)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true

}
