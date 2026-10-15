package main

import (
	"testing"
	"time"
)

// TestTimingSensitive is intentionally flaky: it fails about 75% of the time.
func TestTimingSensitive(t *testing.T) {
	if time.Now().UnixNano()%4 != 0 {
		t.Fatal("timing-sensitive check failed")
	}
}
