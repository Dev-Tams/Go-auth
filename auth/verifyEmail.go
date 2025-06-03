
package auth


import (
	"regexp"
)

func IsValidEmail(email string) bool {
    pattern := `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
    check := regexp.MustCompile(pattern)
    return check.MatchString(email)
}

