package widgets

import (
	"image"
	"time"

	"github.com/guigui-gui/guigui"
	"github.com/guigui-gui/guigui/basicwidget"
	"github.com/guigui-gui/guigui/basicwidget/basicwidgetdraw"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type InfiniteLoadingBar struct {
	guigui.DefaultWidget

	p1, p2 int
	t      time.Time
}

func (d *InfiniteLoadingBar) Build(context *guigui.Context, adder *guigui.ChildAdder) error {
	return nil
}

func (d *InfiniteLoadingBar) Layout(context *guigui.Context, widgetBounds *guigui.WidgetBounds, layouter *guigui.ChildLayouter) {
}

func (d *InfiniteLoadingBar) Tick(context *guigui.Context, widgetBounds *guigui.WidgetBounds) error {
	if time.Since(d.t).Milliseconds() < 6 {
		return nil
	}

	if d.p1 < 100 {
		d.p1++
	}
	if d.p1 > 20 {
		d.p2++
	}

	if d.p2 == 100 {
		d.p1 = 0
		d.p2 = 0
	}

	d.t = time.Now()
	guigui.RequestRedraw(d)
	return nil
}

func (d *InfiniteLoadingBar) Draw(context *guigui.Context, widgetBounds *guigui.WidgetBounds, dst *ebiten.Image) {
	background_clr := basicwidgetdraw.BackgroundColor(context.ColorMode())
	bar_clr, _ := basicwidgetdraw.BorderAccentColors(context.ColorMode(), basicwidgetdraw.RoundedRectBorderTypeRegular)

	b := widgetBounds.Bounds()
	vector.FillRect(dst, float32(b.Min.X), float32(b.Min.Y), float32(b.Dx()), float32(b.Dy()), background_clr, false)

	x1 := b.Min.X + (b.Dx()*d.p1)/100
	x2 := b.Min.X + (b.Dx()*d.p2)/100
	if d.p2 == 0 {
		x2 = b.Min.X
	}
	vector.FillRect(dst, float32(x2), float32(b.Min.Y), float32(x1-x2), float32(b.Dy()), bar_clr, false)
}

func (d *InfiniteLoadingBar) Measure(context *guigui.Context, constraints guigui.Constraints) image.Point {
	var s image.Point

	if h, ok := constraints.FixedHeight(); ok {
		s.Y = h
	} else {
		s.Y = basicwidget.UnitSize(context) / 6
	}

	if w, ok := constraints.FixedWidth(); ok {
		s.X = w
	} else {
		s.X = context.AppBounds().Dx()
	}
	return s
}
