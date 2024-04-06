package utils

import (
	"fmt"
	"testing"
)

func TestSRand(t *testing.T) {
	fmt.Println(SRand(10))
	fmt.Println(rune('A'), rune('Z'))
}
