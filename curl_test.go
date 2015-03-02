package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNothing(t *testing.T) {
	assert.NotNil(t, "lol")
}

// I do not yet know how testing works in Go. But I need to test:
// 1. That JSON and non-JSON responses are parsed
// 2. That -i works
