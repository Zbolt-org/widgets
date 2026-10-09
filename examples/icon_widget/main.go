package main

import (
	"log"
	"widgets"

	"github.com/guigui-gui/guigui"
	"github.com/guigui-gui/guigui/basicwidget"
)

type Root struct {
	guigui.DefaultWidget
	background basicwidget.Background

	close    widgets.Icon
	info_tip widgets.InfoHint

	layout guigui.LinearLayout
}

func (r *Root) Build(context *guigui.Context, adder *guigui.ChildAdder) error {
	adder.AddWidget(&r.background)

	err := r.close.Load("icons/close.png", nil)
	if err != nil {
		log.Println(err.Error())
	}
	adder.AddWidget(&r.close)

	r.info_tip.SetText("This is a infomation hint.")
	adder.AddWidget(&r.info_tip)
	return nil
}

func (r *Root) Layout(context *guigui.Context, widgetBounds *guigui.WidgetBounds, layouter *guigui.ChildLayouter) {
	layouter.LayoutWidget(&r.background, widgetBounds.Bounds())

	u := basicwidget.UnitSize(context)
	if len(r.layout.Items) == 0 {
		r.layout = guigui.LinearLayout{
			Gap:       u / 4,
			Padding:   guigui.Padding{u, u, u, u},
			Items: []guigui.LinearLayoutItem{
				{
					Widget: &r.close,
				},
				{
					Widget: &r.info_tip,
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
