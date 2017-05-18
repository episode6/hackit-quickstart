// copied and modified from https://github.com/etgryphon/stringUp/blob/31534ccd8cac1d3d205c906ea5722928b045ed8c/stringUp.go
package main

import (
	"bytes"
	"regexp"
)

var camelingRegex = regexp.MustCompile("[0-9A-Za-z]+")

func camelCase(src string) string {
	byteSrc := []byte(src)
	chunks := camelingRegex.FindAll(byteSrc, -1)
	for idx, val := range chunks {
		chunks[idx] = bytes.Title(val)
	}
	return string(bytes.Join(chunks, nil))
}
