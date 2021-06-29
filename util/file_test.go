package util

import (
	"fmt"
)

import (
	"testing"
)

func doAbTest() {
	fmt.Println("doAbTest")
}

func AbTest(flag bool) {
	if flag {
		go doAbTest()
	}
	fmt.Println("AbTest")
}

func AbTestWrapper() {
	AbTest(true)
}
func TestGoroutine(t *testing.T) {
	AbTestWrapper()
	//time.Sleep(time.Second)
}
