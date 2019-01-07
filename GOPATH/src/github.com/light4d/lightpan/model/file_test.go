package model

import (
	"fmt"
	"testing"
)

func TestParseFile(t *testing.T) {
	r := "/media@timeloveboy/开发/CODE/github.com/light4d/lightpan/GOPATH/src/github.com/light4d/lightpan/model/file_test.go"
	fmt.Println(ParseFile(r))
	r = "/timeloveboy/lightpan/postman.json"
	fmt.Println(ParseFile(r))
}
