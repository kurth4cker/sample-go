// SPDX-License-Identifier: GPL-3.0-or-later
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

package main

import (
	"bytes"
	"testing"

	"codeberg.org/kurth4cker/go-sample"
)

func TestFhelloln(t *testing.T) {
	buffer := new(bytes.Buffer)
	sample.Fhelloln(buffer, "world")

	want := "hello world\n"
	got := buffer.String()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestShello(t *testing.T) {
	want := "hello world"
	got := sample.Shello("world")

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
