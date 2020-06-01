package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func process(input io.Reader, output io.Writer) {
	findUnique(input, output)
}

func TestProcess1(t *testing.T) {
	input, _ := os.Open("test_1.txt")
	output := &bytes.Buffer{}
	process(input, output)

	got := output.String()
	want := "597"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestProcess2(t *testing.T) {
	input, _ := os.Open("test_2.txt")
	output := &bytes.Buffer{}
	process(input, output)

	got := output.String()
	want := "167890"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestProcess3(t *testing.T) {
	input, _ := os.Open("test_3.txt")
	output := &bytes.Buffer{}
	process(input, output)

	got := output.String()
	want := "59789671238967123"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestProcess4(t *testing.T) {
	input, _ := os.Open("test_4.txt")
	output := &bytes.Buffer{}
	process(input, output)

	got := output.String()
	want := "0"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
