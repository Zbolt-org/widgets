package widgets

import (
	"image"
	"github.com/Zbolt-org/widgets/utils"

	"github.com/guigui-gui/guigui"
	"github.com/guigui-gui/guigui/basicwidget"
)

type TabContainer[T comparable] struct {
	guigui.DefaultWidget

	tab_bar         tab_bar_widget[T]
	items           []TabItem[T]
	selected_widget guigui.Widget

	unselected_widget_defualt basicwidget.Text
	unselected_widget         guigui.Widget
	layout                    guigui.LinearLayout
}

func (d *TabContainer[T]) SetUnselectedWidget(widget guigui.Widget) {
	d.unselected_widget = widget
}

func (d *TabContainer[T]) UnselectedWidget() guigui.Widget {
	if d.unselected_widget == nil {
		return &d.unselected_widget_defualt
	}
	return d.unselected_widget
}

func (d *TabContainer[T]) SetTabItems(items []TabItem[T]) {
	d.items = items
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
	if len(d.items) == 0 {
		d.unselected_widget_defualt.SetValue("Nothing to show")
		d.unselected_widget_defualt.SetEllipsisString("...")
		d.unselected_widget_defualt.SetHorizontalAlign(basicwidget.HorizontalAlignCenter)
		d.unselected_widget_defualt.SetVerticalAlign(basicwidget.VerticalAlignMiddle)
		adder.AddWidget(d.UnselectedWidget())
	} else {
		d.tab_bar.SetTabItems(d.items)
		selected_widget_index, _ := d.tab_bar.SelectedTabItem()
		d.selected_widget = d.items[selected_widget_index].Widget
		adder.AddWidget(d.selected_widget)
		adder.AddWidget(&d.tab_bar)
	}
	return nil
}

func (d *TabContainer[T]) Layout(context *guigui.Context, widgetBounds *guigui.WidgetBounds, layouter *guigui.ChildLayouter) {
	d.layout = guigui.LinearLayout{
		Direction: guigui.LayoutDirectionVertical,
		Gap:       utils.SmallGap(context),
	}

	if len(d.items) == 0 {
		d.layout.Items = []guigui.LinearLayoutItem{
			{
				Widget: d.UnselectedWidget(),
				Size:   guigui.FlexibleSize(1),
			},
		}
	} else {
		d.layout.Items = []guigui.LinearLayoutItem{
			{
				Widget: &d.tab_bar,
			},
			{
				Widget: d.selected_widget,
				Size:   guigui.FlexibleSize(1),
			},
		}
	}

	d.layout.LayoutWidgets(context, widgetBounds.Bounds(), layouter)
}

func (d *TabContainer[T]) Measure(context *guigui.Context, constraints guigui.Constraints) image.Point {
	return d.layout.Measure(context, constraints)
}
