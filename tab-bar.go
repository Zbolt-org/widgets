package widgets

import (
	"image"
	"widgets/utils"

	"github.com/guigui-gui/guigui"
	"github.com/guigui-gui/guigui/basicwidget/basicwidgetdraw"
	"github.com/hajimehoshi/ebiten/v2"
)

type tab_bar_widget[T comparable] struct {
	guigui.DefaultWidget

	selected_tab_index int
	tab_data_items     []TabItem[T]
	tab_items          guigui.WidgetSlice[*tab_item_widget[T]]
	layout             guigui.LinearLayout

	last_measurement image.Point
}

// on_close
// on_switch
// on_widget_select
// select_item
// selected_item
func (d *tab_bar_widget[T]) Env(context *guigui.Context, key guigui.EnvKey, source *guigui.EnvSource) (any, bool) {
	widget, ok := source.Origin.(*tab_item_widget[T])
	if !ok {
		return nil, false
	}
	switch key {
	case tab_item_event_selected:
		d.selected_tab_index = widget.get_index()
	case tab_item_event_close:
		println("Item closed", widget.data.Text)
	}
	return nil, true
}

func (d *tab_bar_widget[T]) SetTabItems(items []TabItem[T]) {
	d.tab_data_items = items
}

func (d *tab_bar_widget[T]) Build(context *guigui.Context, adder *guigui.ChildAdder) error {
	d.tab_items.SetLen(len(d.tab_data_items))

	for i, tab_item_widget := range d.tab_items.All() {
		tab_item_widget.set_index(i)
		tab_item_widget.set_data(d.tab_data_items[i])
		tab_item_widget.set_selected(d.selected_tab_index == i)
		adder.AddWidget(tab_item_widget)
	}

	return nil
}

func (d *tab_bar_widget[T]) Layout(context *guigui.Context, widgetBounds *guigui.WidgetBounds, layouter *guigui.ChildLayouter) {
	padding := tab_item_radius(context)
	d.layout = guigui.LinearLayout{
		Direction: guigui.LayoutDirectionHorizontal,
		Padding:   guigui.Padding{padding, padding, padding, padding},
		Gap:       utils.XSmallGap(context),
		Items:     make([]guigui.LinearLayoutItem, d.tab_items.Len(), d.tab_items.Len()+1),
	}

	for i, tab_item := range d.tab_items.All() {
		d.layout.Items[i].Widget = tab_item
	}
	d.layout.Items = append(d.layout.Items, guigui.LinearLayoutItem{
		Size: guigui.FlexibleSize(1),
	})

	b := widgetBounds.Bounds()
	d.layout.LayoutWidgets(context, b, layouter)
}

func (d *tab_bar_widget[T]) Draw(context *guigui.Context, widgetBounds *guigui.WidgetBounds, dst *ebiten.Image) {
	radius := tab_item_radius(context)
	background_clr := basicwidgetdraw.BackgroundSecondaryColor(context.ColorMode())
	basicwidgetdraw.DrawRoundedRect(context, dst, widgetBounds.Bounds(), background_clr, radius)

	/*
		border_type := basicwidgetdraw.RoundedRectBorderTypeRegular
		clr1, clr2 := basicwidgetdraw.BorderColors(context.ColorMode(), border_type)
		width := 1 * float32(context.Scale())
		basicwidgetdraw.DrawRoundedRectBorder(context, dst, widgetBounds.Bounds(), clr1, clr2, radius, width, border_type)
	*/
}

func (d *tab_bar_widget[T]) Measure(context *guigui.Context, constraints guigui.Constraints) image.Point {
	d.last_measurement = d.layout.Measure(context, constraints)
	return d.last_measurement
}
