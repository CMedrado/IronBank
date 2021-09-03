package transfer

import "strings"

func CheckAuthorizationHeaderType(header string) (string, error) {
	headerSplit := strings.Split(header, " ")
	if "Basic" != headerSplit[0] {
		return "", ErrInvalidCredential
	}
	return headerSplit[1], nil
}
