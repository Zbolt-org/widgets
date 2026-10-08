package main

import (
	"log"
	"widgets"

	"github.com/guigui-gui/guigui"
	"github.com/guigui-gui/guigui/basicwidget"
	"github.com/hajimehoshi/ebiten/v2"
)

type Root struct {
	guigui.DefaultWidget
	background basicwidget.Background

	close    widgets.Icon
	info_tip widgets.Icon
}

func (r *Root) Build(context *guigui.Context, adder *guigui.ChildAdder) error {
	adder.AddWidget(&r.background)

	context.SetPreferredColorMode(ebiten.ColorModeDark)
	err := r.close.Load("icons/close.svg", nil)
	if err != nil {
		log.Println(err.Error())
	}
	adder.AddWidget(&r.close)
	return nil
}

func (r *Root) Layout(context *guigui.Context, widgetBounds *guigui.WidgetBounds, layouter *guigui.ChildLayouter) {
	layouter.LayoutWidget(&r.background, widgetBounds.Bounds())
	layouter.LayoutWidget(&r.close, widgetBounds.Bounds())
}

func main() {
	if err := guigui.Run(&Root{}, &guigui.RunOptions{
		Title: "My App",
	}); err != nil {
		log.Fatal(err)
	}
}
