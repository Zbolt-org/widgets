package widgets

import (
	"image"
	"widgets/utils"

	"github.com/guigui-gui/guigui"
)

type TabContainer[T comparable] struct {
	guigui.DefaultWidget

	tab_items []TabItem[T]
	tab_bar   tab_bar_widget[T]

	layout guigui.LinearLayout
}

func (d *TabContainer[T]) SetTabItems(items []TabItem[T]) {
	d.tab_items = items
}

func (d *TabContainer[T]) Build(context *guigui.Context, adder *guigui.ChildAdder) error {
	d.tab_bar.SetTabItems(d.tab_items)
	adder.AddWidget(&d.tab_bar)
	return nil
}

func (d *TabContainer[T]) Layout(context *guigui.Context, widgetBounds *guigui.WidgetBounds, layouter *guigui.ChildLayouter) {
	d.layout = guigui.LinearLayout{
		Direction: guigui.LayoutDirectionVertical,
		Gap:       utils.SmallGap(context),
		Items: []guigui.LinearLayoutItem{
			{
				Widget: &d.tab_bar,
			},
			{
				Size: guigui.FlexibleSize(1),
			},
		},
	}

	d.layout.LayoutWidgets(context, widgetBounds.Bounds(), layouter)
}

func (d *TabContainer[T]) Measure(context *guigui.Context, constraints guigui.Constraints) image.Point {
	return d.layout.Measure(context, constraints)
}
