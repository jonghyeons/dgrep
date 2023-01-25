package main

import (
	"fmt"
	"testing"
)

func TestScan(t *testing.T) {
	scan(`test`, `sample`)
	fmt.Println()
	scan(`a`, `sample`)
}
