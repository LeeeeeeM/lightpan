package tests

import (
	"testing"
	"github.com/light4d/lightpan/router"
	"github.com/stretchr/testify/assert"
	"fmt"
)

func Test_contentType(t *testing.T) {
	fmt.Print("contenttype:" + router.ContentType("xx.mp3") + "\n")
	assert.Equal(t, router.ContentType("xx.mp3"), "audio/mpeg")
}