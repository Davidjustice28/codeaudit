package utils

import "strings"

func IsConfigOption(value string) bool {
	configTuple := strings.Split(value, "=")
	if len(configTuple) == 2 {
		return true
	} else {
		return false
	}
}
