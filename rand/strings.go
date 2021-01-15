package rand

import (
	"crypto/rand"
	"encoding/base64"
)

const RememberTokenBytes = 32

// will help us generate n random bytes, or will
// return an arror if there was one. This uses the
// crypto/rand pkg so it is safe to use with things
// like remember tokens
func Bytes(n int) ([]byte, error) {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return nil, err
	}
	return b, nil
}

// will generate a byte slice of size nBytes and then return a string
// that is the base64 URL encoded version of that byte slice
func String(nBytes int) (string, error) {
	b, err := Bytes(nBytes)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

// is a helper func designed to generate remember tokens of a predetermined byte size
func RememberToken() (string, error) {
	return String(RememberTokenBytes)
}

/*
	fmt.Println(String(10))
	fmt.Println(RememberToken())
*/
