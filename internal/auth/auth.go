package auth

import (
	"errors"
	"net/http"
	"strings"
)

// Get API key from the headers of the request
// EXAMPLE:
// Authorization: ApiKey {insert api key here}
func GetApiKey(headers http.Header) (string, error) {
	value := headers.Get("Authorization")
	if value == "" {
		return "", errors.New("no authentication header found")
	}

	values := strings.Split(value, " ")
	if len(values) != 2 {
		return "", errors.New("invalid authentication header")
	}
	if values[0] != "ApiKey" {
		return "", errors.New("invalid first part of authentication header")
	}
	return values[1], nil
}