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
// 3. That -v works
// 4. That one can post info
// 5. That one can use HEAD, etc
