package widgets

import (
	"image"
	"io/fs"

	"github.com/guigui-gui/guigui"
	"github.com/hajimehoshi/ebiten/v2"
)

/*
func (item *tab_item[T]) padding(ctx *gui.Context) gui.Padding {
	return basic.NewPadding(0, widget.LineHeight(ctx)/2)
}

func (item *tab_item[T]) gap(ctx *gui.Context) int {
	return widget.UnitSize(ctx) / 4
}
*/

type TabItem[T comparable] struct {
	Text string
	Icon struct {
		icon string
		fs   fs.FS
	}
	Value T
}

type tab_item_widget[T comparable] struct {
	guigui.DefaultWidget

	data TabItem[T]
}

func (d *tab_item_widget[T]) set_data(data TabItem[T]) {
	d.data = data
}

func (d *tab_item_widget[T]) get_data() TabItem[T] {
	return d.data
}

func (d *tab_item_widget[T]) Build(context *guigui.Context, adder *guigui.ChildAdder) error {
	return nil
}

func (d *tab_item_widget[T]) Layout(context *guigui.Context, widgetBounds *guigui.WidgetBounds, layouter *guigui.ChildLayouter) {
}

func (d *tab_item_widget[T]) Draw(context *guigui.Context, widgetBounds *guigui.WidgetBounds, dst *ebiten.Image) {
}

func (d *tab_item_widget[T]) Measure(context *guigui.Context, constraints guigui.Constraints) image.Point {
	var s image.Point
	return s
}
