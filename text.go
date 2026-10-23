package widgets

import (
	"image"

	"github.com/guigui-gui/guigui"
	"github.com/guigui-gui/guigui/basicwidget"
)

type Text struct {
	guigui.DefaultWidget

	text        basicwidget.Text
	contextmenu basicwidget.ContextMenuArea[struct{}]
}

func (d *Text) Text() *basicwidget.Text {
	return &d.text
}

func (d *Text) Build(context *guigui.Context, adder *guigui.ChildAdder) error {
	text := d.Text()
	adder.AddWidget(text)
	return nil
}

func (d *Text) Layout(context *guigui.Context, widgetBounds *guigui.WidgetBounds, layouter *guigui.ChildLayouter) {
	b := widgetBounds.Bounds()
	layouter.LayoutWidget(&d.text, b)
}

func (d *Text) Measure(context *guigui.Context, constraints guigui.Constraints) image.Point {
	return d.text.Measure(context, constraints)
}
