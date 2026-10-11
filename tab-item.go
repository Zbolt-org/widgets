package widgets

import (
	"image"
	"io/fs"
	"widgets/utils"

	"github.com/guigui-gui/guigui"
	"github.com/guigui-gui/guigui/basicwidget"
	"github.com/guigui-gui/guigui/basicwidget/basicwidgetdraw"
	"github.com/hajimehoshi/ebiten/v2"
)

type TabItem[T comparable] struct {
	Text string
	Icon struct {
		icon string
		fs   fs.FS
	}
	Closable, Movable bool
	Value             T
}

func (t *TabItem[T]) is_icon() bool {
	return t.Icon.icon != ""
}

type tab_item_widget[T comparable] struct {
	guigui.DefaultWidget

	data TabItem[T]

	layout guigui.LinearLayout

	icon       Icon
	text       basicwidget.Text
	close_icon Icon
}

func (d *tab_item_widget[T]) set_data(data TabItem[T]) {
	d.data = data
}

func (d *tab_item_widget[T]) get_data() TabItem[T] {
	return d.data
}

func (d *tab_item_widget[T]) Build(context *guigui.Context, adder *guigui.ChildAdder) error {
	if d.data.Closable {
		d.close_icon.Load("icons/close.png", nil, image.Point{})
		adder.AddWidget(&d.close_icon)
	}

	if d.data.is_icon() {
		d.icon.Load(d.data.Icon.icon, d.data.Icon.fs, image.Point{})
		adder.AddWidget(&d.icon)
	}

	// TODO: bold the text if the tab is selected.
	d.text.SetValue(d.data.Text)
	adder.AddWidget(&d.text)
	return nil
}

func (d *tab_item_widget[T]) Layout(context *guigui.Context, widgetBounds *guigui.WidgetBounds, layouter *guigui.ChildLayouter) {
	items := []guigui.LinearLayoutItem{
		{
			Widget: &d.text,
			Size:   guigui.FlexibleSize(1),
		},
	}

	if d.data.Closable {
		item := []guigui.LinearLayoutItem{
			{
				Widget: &d.close_icon,
			},
		}
		items = append(item, items...)
	}

	if d.data.is_icon() {
		items = append(items, guigui.LinearLayoutItem{
			Widget: &d.icon,
		})
	}

	d.layout = guigui.LinearLayout{
		Direction: guigui.LayoutDirectionHorizontal,
		Gap:       utils.SmallGap(context),
		Padding:   utils.SmallPadding(context),
		Items:     items,
	}

	d.layout.LayoutWidgets(context, widgetBounds.Bounds(), layouter)
}

func (d *tab_item_widget[T]) Draw(context *guigui.Context, widgetBounds *guigui.WidgetBounds, dst *ebiten.Image) {
	//clr1, clr2 := basicwidgetdraw.BorderColors(context.ColorMode(), border_type)
	//border_type := basicwidgetdraw.RoundedRectBorderTypeRegular

	background_clr := basicwidgetdraw.BackgroundColor(context.ColorMode())
	r := utils.SmallBorderRadius(context)
	basicwidgetdraw.DrawRoundedRect(context, dst, widgetBounds.Bounds(), background_clr, r)

	// TODO: change the background color based on cursor hover & selected & drag.
}

func (d *tab_item_widget[T]) Measure(context *guigui.Context, _ guigui.Constraints) image.Point {
	return d.layout.Measure(context, guigui.Constraints{})
}
