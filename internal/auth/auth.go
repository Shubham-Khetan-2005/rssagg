package auth

import (
	"errors"
	"net/http"
	"strings"
)

//GetAPIKey will extract API key from the headers of the HTTP request
//Example:
//Authorization: ApiKey {insert api key here}

func GetAPIKey(headers http.Header) (string,error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no information found in Authorisation")
	}
	
	vals := strings.Split(val, " ")
	if len(vals) != 2 || vals[0] != "ApiKey" {
		return "", errors.New("invalid authorisation header")
	}
	
	return vals[1], nil
}