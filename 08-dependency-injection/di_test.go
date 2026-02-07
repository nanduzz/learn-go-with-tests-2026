package main

import (
	"bytes"
	"testing"
)

func TestGreat(t *testing.T) {
	buffer := &bytes.Buffer{}
	Great(buffer, "Go Developer")

	got := buffer.String()
	want := "Hello, Go Developer"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

}
