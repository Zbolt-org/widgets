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

	last_measurement       image.Point
	holding_widget         *tab_item_widget[T]
	switch_canidate_widget *tab_item_widget[T]
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
	case tab_item_event_holding:
		if d.holding_widget != nil {
			return d.holding_widget == widget, true
		}
		d.holding_widget = widget
		return true, true
	case tab_item_event_cancel_holding:
		if d.switch_canidate_widget != d.holding_widget && d.switch_canidate_widget != nil && d.switch_canidate_widget.data.Movable {
			println("Switch canidate", d.switch_canidate_widget.get_data().Text)
		}

		d.cancel_tab_item_moving()
	case tab_item_event_switch_canidate:
		d.switch_canidate_widget = widget
	}
	return nil, true
}

func (d *tab_bar_widget[T]) cancel_tab_item_moving() {
	d.holding_widget = nil
	d.switch_canidate_widget = nil
}

func (d *tab_bar_widget[T]) SetTabItems(items []TabItem[T]) {
	if len(d.tab_data_items) != len(items) {
		d.cancel_tab_item_moving()
	} else if len(items)-1 <= d.selected_tab_index {
		d.selected_tab_index = max(len(items)-1, 0)
	}
	d.tab_data_items = items
}

func (d *tab_bar_widget[T]) Build(context *guigui.Context, adder *guigui.ChildAdder) error {
	d.tab_items.SetLen(len(d.tab_data_items))

	for i, tab_item_widget := range d.tab_items.All() {
		tab_item_widget.set_index(i)
		tab_item_widget.set_data(d.tab_data_items[i])
		tab_item_widget.set_selected(d.selected_tab_index == i)
		if d.holding_widget == tab_item_widget {
			continue
		}
		adder.AddWidget(tab_item_widget)
	}

	if d.holding_widget != nil {
		adder.AddWidget(d.holding_widget)
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
		if tab_item == d.holding_widget {
			size := tab_item.Measure(context, guigui.Constraints{})
			d.layout.Items[i].Size = guigui.FixedSize(size.X)
		} else {
			d.layout.Items[i].Widget = tab_item
		}
	}
	d.layout.Items = append(d.layout.Items, guigui.LinearLayoutItem{
		Size: guigui.FlexibleSize(1),
	})

	b := widgetBounds.Bounds()
	d.layout.LayoutWidgets(context, b, layouter)

	if d.holding_widget != nil {
		size := d.holding_widget.Measure(context, guigui.Constraints{})
		cursor_position := image.Pt(ebiten.CursorPosition())
		bounds := widgetBounds.Bounds()
		bounds.Min.X = cursor_position.X - d.holding_widget.RelativeCursorPosition().X
		bounds.Max.X = bounds.Min.X + size.X
		bounds.Min.Y += d.layout.Padding.Top
		bounds.Max.Y -= d.layout.Padding.Bottom
		layouter.LayoutWidget(d.holding_widget, bounds)
	}
}

func (d *tab_bar_widget[T]) HandlePointingInput(context *guigui.Context, widgetBounds *guigui.WidgetBounds) guigui.HandleInputResult {
	if d.holding_widget != nil && !widgetBounds.IsHitAtCursor() {
		d.holding_widget.cancel_moving()
		d.cancel_tab_item_moving()
		return guigui.HandleInputByWidget(d)
	}
	return guigui.HandleInputResult{}
}

func (d *tab_bar_widget[T]) Draw(context *guigui.Context, widgetBounds *guigui.WidgetBounds, dst *ebiten.Image) {
	radius := tab_item_radius(context)
	background_clr := basicwidgetdraw.BackgroundSecondaryColor(context.ColorMode())
	basicwidgetdraw.DrawRoundedRect(context, dst, widgetBounds.Bounds(), background_clr, radius)
}

func (d *tab_bar_widget[T]) Measure(context *guigui.Context, constraints guigui.Constraints) image.Point {
	d.last_measurement = d.layout.Measure(context, constraints)
	return d.last_measurement
}
