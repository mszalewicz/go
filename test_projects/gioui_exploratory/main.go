// My main
package main

import (
	"image/color"
	"log"
	"os"

	"this/gui"

	"gioui.org/app"
	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget/material"
)

func main() {
	go func() {
		window := new(app.Window)
		err := run(window)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

var (
	BEIGE        = color.NRGBA{R: 247, G: 239, B: 229, A: 255}
	BLACK        = color.NRGBA{R: 0, G: 0, B: 0, A: 255}
	BLUE         = color.NRGBA{R: 150, G: 201, B: 244, A: 255}
	CHARCOAL     = color.NRGBA{R: 0, G: 0, B: 0, A: 150}
	GREEN        = color.NRGBA{R: 100, G: 196, B: 166, A: 255}
	GREY         = color.NRGBA{R: 100, G: 100, B: 100, A: 255}
	GREY_LIGHT   = color.NRGBA{R: 235, G: 235, B: 235, A: 255}
	ORANGE       = color.NRGBA{R: 235, G: 178, B: 10, A: 100}
	PURPLE       = color.NRGBA{R: 161, G: 100, B: 196, A: 255}
	PURPLE_LIGHT = color.NRGBA{R: 161, G: 100, B: 196, A: 120}
	RED          = color.NRGBA{R: 238, G: 78, B: 78, A: 220}
	WHITE        = color.NRGBA{R: 255, G: 255, B: 255, A: 255}
	TEXTSIZE     = unit.Sp(30)
)

func run(window *app.Window) error {
	theme := material.NewTheme()
	var ops op.Ops
	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			// This graphics context is used for managing the rendering state.
			gtx := app.NewContext(&ops, e)

			const buttonSize = 12

			theme.Bg = GREY_LIGHT
			theme.Fg = BLACK

			hugeBtnInfo, hugeBtnWidget := gui.CreateDefaultButton(theme, "HUGE BUTTON!!! ", 0, 10, 10, 10, 5)
			hugeBtnInfo.Font.Weight = font.Bold

			openBtnInfo, openBtnWidget := gui.CreateDefaultButton(theme, "OPEN", 0, 10, 10, 5, 5)
			openBtnInfo.Font.Weight = font.Bold

			deleteBtnInfo, deleteBtnWidget := gui.CreateDefaultButton(theme, "DELETE", 0, 10, 10, 5, 10)
			deleteBtnInfo.Font.Weight = font.Bold

			splitButtons := new(gui.SplitView)
			splitButtonsWigdet := func(gtx layout.Context) layout.Dimensions { return splitButtons.Layout(gtx, openBtnWidget, deleteBtnWidget, 0.5) }

			splitEntry := new(gui.SplitView)
			splitEntry.Layout(gtx, hugeBtnWidget, splitButtonsWigdet, 0.7)

			e.Frame(gtx.Ops)
		}
	}
}
