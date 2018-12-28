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

func Test_randomElement(t *testing.T) {
	array := []string{"127.0.0.1:9001", "192.168.1.9527:9001", "192.168.1.9528:9001", "192.168.1.9529:9001",
	"192.168.1.9530:9001"}
	element, _ := router.RandomElement(array)
	fmt.Println("randomElement:" + element)
}