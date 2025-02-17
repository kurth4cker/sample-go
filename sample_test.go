// SPDX-License-Identifier: GPL-3.0-or-later
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

package main

import (
	"testing"

	"codeberg.org/kurth4cker/go-sample"
)

func TestShello(t *testing.T) {
	want := "hello world"
	got := sample.Shello("world")

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
