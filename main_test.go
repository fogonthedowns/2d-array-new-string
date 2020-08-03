package main

import (
	"testing"
)

func TestFun(t *testing.T) {
	in := []string{"Daisy", "Rose", "Hyacinth", "Poppy"}
	out := newString(in)
	expected := "DRHPaoyoisapsecpyiynth"

	if out != expected {
		t.Errorf("got %s, want %s", out, expected)
	}

}
