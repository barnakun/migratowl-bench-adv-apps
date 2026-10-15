package main

import "testing"

func TestHello(t *testing.T) {
	if got := Hello("world"); got != "hello, world" {
		t.Fatalf("got %q", got)
	}
}
