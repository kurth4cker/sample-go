// SPDX-License-Identifier: GPL-3.0-or-later
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

package test

import (
	"bytes"
	"testing"

	"codeberg.org/kurth4cker/go-sample"
	"codeberg.org/kurth4cker/go-sample/assert"
)

func TestFhelloln(t *testing.T) {
	buffer := new(bytes.Buffer)
	sample.Fhelloln(buffer, "world")

	want := "hello world\n"
	got := buffer.String()

	assert.Equal(t, got, want)
}

func TestShello(t *testing.T) {
	want := "hello world"
	got := sample.Shello("world")

	assert.Equal(t, got, want)
}
