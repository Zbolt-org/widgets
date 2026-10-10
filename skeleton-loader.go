package widgets

import (
	"image"
	"image/color"
	"time"
	"widgets/utils"

	"github.com/guigui-gui/guigui"
	"github.com/guigui-gui/guigui/basicwidget/basicwidgetdraw"
	"github.com/hajimehoshi/ebiten/v2"
)

type SkeletonLoader[T guigui.Widget] struct {
	guigui.DefaultWidget

	widget     LazyWidget[T]
	is_loading bool

	backward    bool
	color_level int

	t time.Time
}

func (d *SkeletonLoader[T]) Loading() bool {
	return d.is_loading
}

func (d *SkeletonLoader[T]) SetLoading(loading bool) {
	d.is_loading = loading
}

func (d *SkeletonLoader[T]) Widget() T {
	return d.widget.Widget()
}

func (d *SkeletonLoader[T]) Tick(context *guigui.Context, widgetBounds *guigui.WidgetBounds) error {
	if time.Since(d.t).Milliseconds() < 60 {
		return nil
	}

	guigui.RequestRedraw(d)

	if d.color_level >= 20 {
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
	if !d.is_loading {
		adder.AddWidget(d.widget.Widget())
	}
	return nil
}

func (d *SkeletonLoader[T]) Layout(context *guigui.Context, widgetBounds *guigui.WidgetBounds, layouter *guigui.ChildLayouter) {
	if !d.is_loading {
		layouter.LayoutWidget(d.widget.Widget(), widgetBounds.Bounds())
	}
}

func (d *SkeletonLoader[T]) Draw(context *guigui.Context, widgetBounds *guigui.WidgetBounds, dst *ebiten.Image) {
	if !d.is_loading {
		return
	}
	b := widgetBounds.Bounds()

	var clr_diffrence uint8
	if context.ColorMode() == ebiten.ColorModeLight {
		clr_diffrence = uint8(240 - d.color_level)
	} else {
		clr_diffrence = uint8(40 + d.color_level)
	}

	clr := color.RGBA{clr_diffrence, clr_diffrence, clr_diffrence, 255}
	basicwidgetdraw.DrawRoundedRect(context, dst, b, clr, utils.SmallBorderRadius(context))
}

func (d *SkeletonLoader[T]) Measure(context *guigui.Context, constraints guigui.Constraints) image.Point {
	return d.widget.Widget().Measure(context, constraints)
}
