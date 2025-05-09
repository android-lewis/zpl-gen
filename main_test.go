package main

import "testing"

func TestMain(t *testing.T) {
	want := "Hello World"
	got := hello()

	if want != got {
		t.Fatalf("wanted %s got %s", want, got)
	}
}
