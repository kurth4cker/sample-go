// SPDX-License-Identifier: GPL-3.0-or-later
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

package main

import (
	"fmt"
	"testing"

	"codeberg.org/kurth4cker/go-sample"
)

func TestGreet(t *testing.T) {
	names := []string{
		"world",
		"kurth4cker",
		"kthzk",
	}

	for _, name := range names {
		want := fmt.Sprintf("hello %s", name)
		got := sample.Greet(name)

		if got != want {
			t.Errorf("got %q, want %q, given %q",
				got, want, name)
		}
	}
}
