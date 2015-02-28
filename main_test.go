package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNothing(t *testing.T) {
	assert.NotNil(t, "lol")
}

// I do not yet know how testing works in Go. But I need to test:
// 1. That an unset server fails
// 2. That a set server passes
// 3. That a non-proper method fails
// 4. that a lower case method works
