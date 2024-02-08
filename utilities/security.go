package utilities

import (
	"encoding/base64"
	"errors"
	"strings"
)

func DecodeBasicAuth(header string) ([]string, error) {
	StatsdClient.Inc("utilities.security", 1, 1.0)
	if len(strings.TrimSpace(header)) == 0 {
		return nil, errors.New("No header content present")
	}

	header = header[6:]
	bearer, err := base64.StdEncoding.DecodeString(header)
	if err != nil {
		return nil, errors.New("No basic auth header content present")
	}

	return strings.Split(string(bearer), ":"), nil
}
