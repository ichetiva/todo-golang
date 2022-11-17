package misc

import (
	"strings"
)

type NonValidToken struct{}

func (e *NonValidToken) Error() string {
	return "Non-valid token"
}

func GetTokenFromHeader(header string) (string, error) {
	if !strings.Contains(header, "Bearer") {
		return "", &NonValidToken{}
	}
	token := strings.Split(header, " ")[1]
	return token, nil
}
