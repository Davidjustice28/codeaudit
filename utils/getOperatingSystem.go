package utils

import "runtime"

func GetOs() string {
	systemFound := runtime.GOOS
	return systemFound
}
