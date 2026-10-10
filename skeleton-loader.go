package widgets

import (
	"image"
	"image/color"
	"time"

	"github.com/guigui-gui/guigui"
	"github.com/guigui-gui/guigui/basicwidget"
	"github.com/guigui-gui/guigui/basicwidget/basicwidgetdraw"
	"github.com/hajimehoshi/ebiten/v2"
)

type SkeletonLoader[T guigui.Widget] struct {
	guigui.DefaultWidget

	widget LazyWidget[T]

	backward    bool
	color_level int

	t time.Time
}

func (d *SkeletonLoader[T]) Widget() T {
	return d.widget.Widget()
}

func (d *SkeletonLoader[T]) Tick(context *guigui.Context, widgetBounds *guigui.WidgetBounds) error {
	if time.Since(d.t).Milliseconds() < 60 {
		return nil
	}

	guigui.RequestRedraw(d)

	if d.color_level >= 8 {
		d.backward = true
	} else if d.color_level == 0 {
		d.backward = false
	}

	if d.backward {
		d.color_level--
	} else {
		d.color_level++
	}

	d.t = time.Now()
	return nil
}

func (d *SkeletonLoader[T]) Build(context *guigui.Context, adder *guigui.ChildAdder) error {
	//adder.AddWidget(d.widget.Widget())
	return nil
}

func (d *SkeletonLoader[T]) Layout(context *guigui.Context, widgetBounds *guigui.WidgetBounds, layouter *guigui.ChildLayouter) {
	//layouter.LayoutWidget(d.widget.Widget(), widgetBounds.Bounds())
}

func (d *SkeletonLoader[T]) Draw(context *guigui.Context, widgetBounds *guigui.WidgetBounds, dst *ebiten.Image) {
	b := widgetBounds.Bounds()

	var clr_diffrence uint8
	if context.ColorMode() == ebiten.ColorModeLight {
		clr_diffrence = uint8(240 - d.color_level)
	} else {
		clr_diffrence = uint8(40 + d.color_level)
	}
	basicwidgetdraw.DrawRoundedRect(context, dst, b, color.RGBA{clr_diffrence, clr_diffrence, clr_diffrence, 255}, basicwidget.UnitSize(context)/2)
}

func (d *SkeletonLoader[T]) Measure(context *guigui.Context, constraints guigui.Constraints) image.Point {
	return d.widget.Widget().Measure(context, constraints)
}
