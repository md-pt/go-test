package main

import "testing"

func TestHello(t *testing.T) {
	ans := Hello()
	if ans != "Hello World" {
		t.Errorf("Hello() = %s; want Hello World", ans)
	}
}

// Test
