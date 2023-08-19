package utils

func UserExists(users []string, newUser string) bool {
	for _, element := range users {
		if element == newUser {
			return true
		}
	}
	return false
}
