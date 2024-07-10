package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetApiKey(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	arr := strings.Split(authHeader, "ApiKey ")
	if len(arr) < 2 {
		return "", errors.New("Unauthorized, please provide an api key")
	}
	return arr[1], nil
}
