package output

import (
	"bytes"
)

func Output(multiline []string) string {
	var buffer bytes.Buffer

	for _, s := range multiline {
		buffer.WriteString(s + "\n")
	}

	return buffer.String()
}
