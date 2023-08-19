package utils

func IsValidNamingConvention(naming_convention string) bool {
	namingConventions := []string{"camel", "snake", "pascal"}
	isValidConventions := false
	for i := 0; i < len(namingConventions); i++ {
		if namingConventions[i] == naming_convention {
			isValidConventions = true
			break
		}
	}
	return isValidConventions
}
