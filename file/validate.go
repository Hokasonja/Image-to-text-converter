package file

import (
	"errors"
	"strings"
)

var ext = []string{"png", "jpg"}

func Validate(filename string) (string, error) {
	r := strings.Split(filename, ".")
	l := len(r)

	flag := false
	for _, e := range ext {
		if r[l-1] == e {
			flag = true
			break
		}
	}

	if flag == true {
		return strings.Join(r, "."), nil
	} else {
		return "", errors.New("invalid extension")
	}
}
