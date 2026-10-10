package widgets

/*
import (
	"image"
	"image/color"
	"widgets/utils"

	"github.com/guigui-gui/guigui"
	"github.com/guigui-gui/guigui/basicwidget/basicwidgetdraw"
	"github.com/hajimehoshi/ebiten/v2"
)

type WidgetWithBorder[T guigui.Widget] struct {
	guigui.DefaultWidget

	widget LazyWidget[T]

	custome_borders bool
	clr1, clr2      color.Color
	radius          int
	borderWidth     float32
	borderType      basicwidgetdraw.RoundedRectBorderType
}

func (d *WidgetWithBorder[T]) SetBorder(clr1, clr2 color.Color, radius int, borderWidth float32, borderType basicwidgetdraw.RoundedRectBorderType) {
	d.clr1 = clr1
	d.clr2 = clr2
	d.radius = radius
	d.borderWidth = borderWidth
	d.borderType = borderType
	d.custome_borders = true
}

func (d *WidgetWithBorder[T]) set_defualt_border(context *guigui.Context) {
	clr1, clr2 := basicwidgetdraw.BorderColors(context.ColorMode(), basicwidgetdraw.RoundedRectBorderTypeRegular)
	d.SetBorder(clr1, clr2, utils.SmallBorderRadius(context), 1, basicwidgetdraw.RoundedRectBorderTypeRegular)
}

func (d *WidgetWithBorder[T]) border(context *guigui.Context) (clr1, clr2 color.Color, radius int, borderWidth float32, borderType basicwidgetdraw.RoundedRectBorderType) {
	if !d.custome_borders {
		d.set_defualt_border(context)
	}
	return d.clr1, d.clr2, d.radius, d.borderWidth, d.borderType
}

func (d *WidgetWithBorder[T]) Build(context *guigui.Context, adder *guigui.ChildAdder) error {
	adder.AddWidget(d.widget.Widget())
	return nil
}

func (d *WidgetWithBorder[T]) Layout(context *guigui.Context, widgetBounds *guigui.WidgetBounds, layouter *guigui.ChildLayouter) {
	layouter.LayoutWidget(d.widget.Widget(), widgetBounds.Bounds())
}

func (d *WidgetWithBorder[T]) Draw(context *guigui.Context, widgetBounds *guigui.WidgetBounds, dst *ebiten.Image) {
	clr1, clr2, r, w, t := d.border(context)
	basicwidgetdraw.DrawRoundedRectBorder(context, dst, widgetBounds.Bounds(), clr1, clr2, r, w, t)
}

func (d *WidgetWithBorder[T]) Measure(context *guigui.Context, constraints guigui.Constraints) image.Point {
	return d.widget.Widget().Measure(context, constraints)
}
*/
