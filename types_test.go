package main

import (
	"fmt"
	"testing"
)

func TestNewAccount(t *testing.T) {
	acc, err := NewAccount("a", "b", "hunger")

	fmt.Println(acc)

	if err != nil {
		t.Errorf("there was an error")
	}
}
