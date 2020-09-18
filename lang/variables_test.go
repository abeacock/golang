package main

import "testing"

func TestVariables(t *testing.T) {
	var a = 1
	if a != 1 {
		t.Error()
	}

	var b = 'b'
	if b != 'b' {
		t.Error()
	}

	var c string = "c"
	if c != "c" {
		t.Error()
	}

	d := 1
	if d != 1 {
		t.Error()
	}

	e := "e"
	if e != "e" {
		t.Error()
	}

	var f int = 2
	if f != 2 {
		t.Error()
	}
}
