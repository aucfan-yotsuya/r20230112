package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestForRun1a(t *testing.T) {
	const (
		B = ""
		N = "*"
	)
	assert.False(t, forRun1a(0, B, B), "case 0, both are blanks")
	assert.False(t, forRun1a(0, B, N), "case 0, blank, not blank")
	assert.False(t, forRun1a(0, N, B), "case 0, not blank, blank")
	assert.False(t, forRun1a(0, N, N), "case 0, both are not blanks")
	assert.False(t, forRun1a(1, B, B), "case 1, both are blanks")
	assert.True(t, forRun1a(1, B, N), "case 1, blank, not blank")
	assert.True(t, forRun1a(1, N, B), "case 1, not blank, blank")
	assert.True(t, forRun1a(1, N, N), "case 1, both are not blanks")
}
func TestForRun2a(t *testing.T) {
	const (
		B = ""
		N = "*"
	)
	assert.False(t, forRun2a(0, B, B), "case 0, both are blanks")
	assert.True(t, forRun2a(0, B, N), "case 0, blank, not blank")
	assert.True(t, forRun2a(0, N, B), "case 0, not blank, blank")
	assert.True(t, forRun2a(0, N, N), "case 0, both are not blanks")
	assert.False(t, forRun2a(1, B, B), "case 1, both are blanks")
	assert.False(t, forRun2a(1, B, N), "case 1, blank, not blank")
	assert.False(t, forRun2a(1, N, B), "case 1, not blank, blank")
	assert.False(t, forRun2a(1, N, N), "case 1, both are not blanks")
}
func TestForRun1b(t *testing.T) {
	const (
		B = ""
		N = "*"
	)
	assert.False(t, forRun1b(0, B, B), "case 0, both are blanks")
	assert.False(t, forRun1b(0, B, N), "case 0, blank, not blank")
	assert.False(t, forRun1b(0, N, B), "case 0, not blank, blank")
	assert.False(t, forRun1b(0, N, N), "case 0, both are not blanks")
	assert.False(t, forRun1b(1, B, B), "case 1, both are blanks")
	assert.True(t, forRun1b(1, B, N), "case 1, blank, not blank")
	assert.True(t, forRun1b(1, N, B), "case 1, not blank, blank")
	assert.True(t, forRun1b(1, N, N), "case 1, both are not blanks")
}
func TestForRun2b(t *testing.T) {
	const (
		B = ""
		N = "*"
	)
	assert.False(t, forRun2b(0, B, B), "case 0, both are blanks")
	assert.True(t, forRun2b(0, B, N), "case 0, blank, not blank")
	assert.True(t, forRun2b(0, N, B), "case 0, not blank, blank")
	assert.True(t, forRun2b(0, N, N), "case 0, both are not blanks")
	assert.False(t, forRun2b(1, B, B), "case 1, both are blanks")
	assert.False(t, forRun2b(1, B, N), "case 1, blank, not blank")
	assert.False(t, forRun2b(1, N, B), "case 1, not blank, blank")
	assert.False(t, forRun2b(1, N, N), "case 1, both are not blanks")
}
func TestIsPositive(t *testing.T) {
	assert.False(t, isPositive(-1), "case -1")
	assert.False(t, isPositive(0), "case 0")
	assert.True(t, isPositive(1), "case 1")
}
func TestIsZero(t *testing.T) {
	assert.False(t, isZero(-1), "case -1")
	assert.True(t, isZero(0), "case 0")
	assert.False(t, isZero(1), "case 1")
}
func TestEitherNotBlank(t *testing.T) {
	const (
		B = ""
		N = "*"
	)
	assert.False(t, eitherNotBlank(B, B), "case both are blanks")
	assert.True(t, eitherNotBlank(B, N), "case blank, not blank")
	assert.True(t, eitherNotBlank(N, B), "case not blank, blank")
	assert.True(t, eitherNotBlank(N, N), "case both are not blanks")
}
func TestIsBlank(t *testing.T) {
	const (
		B = ""
		N = "*"
	)
	assert.True(t, isBlank(B), "case blank")
	assert.False(t, isBlank(N), "case not blank")
}
