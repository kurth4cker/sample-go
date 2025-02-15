// SPDX-License-Identifier: Unlicense OR MIT
// SPDX-FileCopyrightText: 2019 The Gio authors

// Copied from https://gioui.org/doc/learn/get-started

package main

import (
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/widget/material"
)

func run(window *app.Window) error {
	theme := material.NewTheme()
	var ops op.Ops
	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)

			title := material.H1(theme, "hello Gio")
			maroon := color.NRGBA{R: 127, G: 0, B: 0, A: 255}
			title.Color = maroon
			title.Alignment = text.Middle
			title.Layout(gtx)

			e.Frame(gtx.Ops)
		}
	}
}

func main() {
	go func() {
		window := new(app.Window)
		if err := run(window); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}
