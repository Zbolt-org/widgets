package widgets

import (
	"image"

	"github.com/guigui-gui/guigui"
	"github.com/guigui-gui/guigui/basicwidget"
)

type ToggleWithText struct {
	guigui.DefaultWidget

	text   basicwidget.Text
	basicwidget.Toggle
}

func (*ToggleWithText) Build(context *guigui.Context, adder *guigui.ChildAdder) error {
	return nil
}

func (*ToggleWithText) Layout(context *guigui.Context, widgetBounds *guigui.WidgetBounds, layouter *guigui.ChildLayouter) {
}

func (d *ToggleWithText) Measure(context *guigui.Context, constraints guigui.Constraints) image.Point {
	var s image.Point
	return s
}
