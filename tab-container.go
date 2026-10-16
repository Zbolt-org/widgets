package widgets

import (
	"image"
	"widgets/utils"

	"github.com/guigui-gui/guigui"
)

type TabContainer[T comparable] struct {
	guigui.DefaultWidget

	tab_bar   tab_bar_widget[T]
	layout guigui.LinearLayout
}

func (d *TabContainer[T]) SetTabItems(items []TabItem[T]) {
	d.tab_bar.SetTabItems(items)
}

func (d *TabContainer[T]) TabItemsLenght() int {
	return d.tab_bar.TabItemsLenght()
}

func (d *TabContainer[T]) SelectTabItem(index int) {
	d.tab_bar.SelectTabItem(index)
}

func (d *TabContainer[T]) SelectedTabItem() (index int, ok bool) {
	return d.tab_bar.SelectedTabItem()
}

func (d *TabContainer[T]) OnTabClose(fn func(context *guigui.Context, index int, data TabItem[T])) {
	d.tab_bar.OnTabClose(fn)
}

func (d *TabContainer[T]) OnTabSelect(fn func(context *guigui.Context, index int, data TabItem[T])) {
	d.tab_bar.OnTabSelect(fn)
}

func (d *TabContainer[T]) OnTabSwitch(fn func(context *guigui.Context, index1 int, data1 TabItem[T], index2 int, data2 TabItem[T])) {
	d.tab_bar.OnTabSwitch(fn)
}

func (d *TabContainer[T]) Build(context *guigui.Context, adder *guigui.ChildAdder) error {
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
