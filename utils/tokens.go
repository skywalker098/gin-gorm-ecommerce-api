package utils

import "encoding/base64"

func GenerateBasicAuthToken(email string) string {
	token := base64.StdEncoding.EncodeToString([]byte(email))
	return token
}

//decode token
func DecodeBasicAuthToken(token string) string {
	//basic auth
	decoded, _ := base64.StdEncoding.DecodeString(token)
	return string(decoded)
}
