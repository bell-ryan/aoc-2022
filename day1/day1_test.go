package day1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	want, got := 68923, Part1()
	assert.Equal(t, want, got)
}

func BenchmarkPart1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = Part1()
	}
}
func TestPart2(t *testing.T) {
	want, got := 200044, Part2()
	assert.Equal(t, want, got)
}

func BenchmarkPart2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = Part2()
	}
}
