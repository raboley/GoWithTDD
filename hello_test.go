package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Russell")
	want := "Hello, Russell"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
