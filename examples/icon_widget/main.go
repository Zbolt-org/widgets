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
	toggle     widgets.LabeledToggle

	close    widgets.Icon
	info_tip widgets.InfoHint

	btn widgets.SkeletonLoader[*basicwidget.Button]

	layout guigui.LinearLayout
}

func (r *Root) Build(context *guigui.Context, adder *guigui.ChildAdder) error {
	adder.AddWidget(&r.background)
	r.toggle.SetLabel("Color mode")
	r.toggle.Toggle().OnValueChanged(func(context *guigui.Context, value bool) {
		if value {
			context.SetPreferredColorMode(ebiten.ColorModeDark)
		} else {
			context.SetPreferredColorMode(ebiten.ColorModeLight)
		}
	})
	adder.AddWidget(&r.toggle)

	if !r.close.HasIcon() {
		err := r.close.Load("icons/close.png", nil)
		if err != nil {
			return err
		}
	}
	adder.AddWidget(&r.close)
	r.info_tip.SetText("This is a infomation hint.")
	adder.AddWidget(&r.info_tip)

	r.btn.Widget().SetText("Hello")
	adder.AddWidget(&r.btn)
	return nil
}

func (r *Root) Layout(context *guigui.Context, widgetBounds *guigui.WidgetBounds, layouter *guigui.ChildLayouter) {
	layouter.LayoutWidget(&r.background, widgetBounds.Bounds())

	u := basicwidget.UnitSize(context)
	if len(r.layout.Items) == 0 {
		r.layout = guigui.LinearLayout{
			Gap:       u / 4,
			Padding:   guigui.Padding{u, u, u, u},
			Direction: guigui.LayoutDirectionHorizontal,
			Items: []guigui.LinearLayoutItem{
				{
					Widget: &r.close,
				},
				{
					Widget: &r.info_tip,
				},
				{
					Widget: &r.toggle,
				},
				{
					Widget: &r.btn,
				},
			},
		}
	}

	r.layout.LayoutWidgets(context, widgetBounds.Bounds(), layouter)
}

func main() {
	if err := guigui.Run(&Root{}, &guigui.RunOptions{
		Title: "My App",
	}); err != nil {
		log.Fatal(err)
	}
}
