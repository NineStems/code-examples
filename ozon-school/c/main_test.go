package main

import (
	"testing"
)

func BenchmarkCodegen(b *testing.B) {
	for i := 0; i < 5; i++ {
		DoIt()
	}
}
