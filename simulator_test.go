package main

import (
	"testing"
)

func TestStringInSlicePositive(t *testing.T) {

	slice := []string{"foo", "bar", "bli"}
	str := "bli"

	got := stringInSlice(str, slice)
	want := true

	if got != want {
		t.Errorf("\n stringInSlice(\"bli\", {\"foo\", \"bar\", \"bli\"}) \n got: %v; want: %v;", got, want)
	}
}

func TestStringInSliceNegative(t *testing.T) {

	slice := []string{"foo", "bar", "bli"}
	str := "nono"

	got := stringInSlice(str, slice)
	want := false

	if got != want {
		t.Errorf("\n stringInSlice(\"nono\", {\"foo\", \"bar\", \"bli\"}) \n got: %v; want: %v;", got, want)
	}
}
