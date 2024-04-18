package main

import "testing"

func TestHello(t *testing.T) {
	want := "Hello GO"

	got := hello()
	if got != want {
		t.Fatalf("want %s, got %s\n", want, got)
	}
}
