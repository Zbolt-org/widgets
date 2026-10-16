package widgets

import (
	"image"
	"widgets/utils"

	"github.com/guigui-gui/guigui"
	"github.com/guigui-gui/guigui/basicwidget"
	"github.com/guigui-gui/guigui/basicwidget/basicwidgetdraw"
	"github.com/hajimehoshi/ebiten/v2"
)

type tab_bar_content_widget[T comparable] struct {
	guigui.DefaultWidget

	selected_tab_index int
	tab_data_items     []TabItem[T]
	tab_items          guigui.WidgetSlice[*tab_item_widget[T]]
	layout             guigui.LinearLayout

	holding_widget         *tab_item_widget[T]
	switch_canidate_widget *tab_item_widget[T]

	on_tab_item_close_fn  func(context *guigui.Context, index int, data TabItem[T])
	on_tab_item_select_fn func(context *guigui.Context, index int, data TabItem[T])
	on_tab_item_switch_fn func(context *guigui.Context, index1 int, data1 TabItem[T], index2 int, data2 TabItem[T])
}

func (d *tab_bar_content_widget[T]) TabItemsLenght() int {
	return len(d.tab_data_items)
}

func (d *tab_bar_content_widget[T]) SelectTabItem(index int) {
	d.selected_tab_index = index
}

func (d *tab_bar_content_widget[T]) SelectedTabItem() (index int, ok bool) {
	return d.selected_tab_index, len(d.tab_data_items) > 0
}

func (d *tab_bar_content_widget[T]) OnTabClose(fn func(context *guigui.Context, index int, data TabItem[T])) {
	d.on_tab_item_close_fn = fn
}

func (d *tab_bar_content_widget[T]) OnTabSelect(fn func(context *guigui.Context, index int, data TabItem[T])) {
	d.on_tab_item_select_fn = fn
}

func (d *tab_bar_content_widget[T]) OnTabSwitch(fn func(context *guigui.Context, index1 int, data1 TabItem[T], index2 int, data2 TabItem[T])) {
	d.on_tab_item_switch_fn = fn
}

func (d *tab_bar_content_widget[T]) Env(context *guigui.Context, key guigui.EnvKey, source *guigui.EnvSource) (any, bool) {
	widget, ok := source.Origin.(*tab_item_widget[T])
	if !ok {
		return nil, false
	}
	switch key {
	case tab_item_event_selected:
		d.selected_tab_index = widget.get_index()
		if d.on_tab_item_select_fn != nil {
			d.on_tab_item_select_fn(context, widget.get_index(), widget.get_data())
		}
	case tab_item_event_close:
		if d.on_tab_item_close_fn != nil {
			d.on_tab_item_close_fn(context, widget.get_index(), widget.get_data())
		}
	case tab_item_event_holding:
		if d.holding_widget != nil {
			return d.holding_widget == widget, true
		} else if len(d.tab_data_items) == 1 {
			return false, true
		}
		d.holding_widget = widget
		return true, true
	case tab_item_event_cancel_holding:
		if d.switch_canidate_widget != d.holding_widget && d.switch_canidate_widget != nil && d.switch_canidate_widget.data.Movable && d.on_tab_item_switch_fn != nil {
			d.on_tab_item_switch_fn(context, widget.get_index(), widget.get_data(), d.switch_canidate_widget.get_index(), d.switch_canidate_widget.get_data())
		}

		d.cancel_tab_item_moving()
	case tab_item_event_switch_canidate:
		d.switch_canidate_widget = widget
	}
	return nil, false
}

func (d *tab_bar_content_widget[T]) cancel_tab_item_moving() {
	d.holding_widget = nil
	d.switch_canidate_widget = nil
}

func (d *tab_bar_content_widget[T]) SetTabItems(items []TabItem[T]) {
	if len(d.tab_data_items) != len(items) {
		d.cancel_tab_item_moving()
	}
	if len(items)-1 <= d.selected_tab_index {
		d.selected_tab_index = max(len(items)-1, 0)
	}
	d.tab_data_items = items
}

func (d *tab_bar_content_widget[T]) Build(context *guigui.Context, adder *guigui.ChildAdder) error {
	if len(d.tab_data_items) == 1 {
		d.cancel_tab_item_moving()
		if d.holding_widget != nil {
			d.holding_widget.cancel_moving()
		}
	}

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

func (d *tab_bar_content_widget[T]) Layout(context *guigui.Context, widgetBounds *guigui.WidgetBounds, layouter *guigui.ChildLayouter) {
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

func (d *tab_bar_content_widget[T]) HandlePointingInput(context *guigui.Context, widgetBounds *guigui.WidgetBounds) guigui.HandleInputResult {
	if d.holding_widget != nil && !widgetBounds.IsHitAtCursor() {
		d.holding_widget.cancel_moving()
		d.cancel_tab_item_moving()
		return guigui.HandleInputByWidget(d)
	}
	return guigui.HandleInputResult{}
}

func (d *tab_bar_content_widget[T]) Draw(context *guigui.Context, widgetBounds *guigui.WidgetBounds, dst *ebiten.Image) {
	radius := tab_item_radius(context)
	background_clr := basicwidgetdraw.BackgroundSecondaryColor(context.ColorMode())
	basicwidgetdraw.DrawRoundedRect(context, dst, widgetBounds.Bounds(), background_clr, radius)
}

func (d *tab_bar_content_widget[T]) Measure(context *guigui.Context, constraints guigui.Constraints) image.Point {
	return d.layout.Measure(context, constraints)
}

type tab_bar_widget[T comparable] struct {
	guigui.DefaultWidget

	panel   basicwidget.Panel
	content tab_bar_content_widget[T]

	layout guigui.LinearLayout
}

func (d *tab_bar_widget[T]) SetTabItems(items []TabItem[T]) {
	d.content.SetTabItems(items)
}

func (d *tab_bar_widget[T]) TabItemsLenght() int {
	return d.content.TabItemsLenght()
}

func (d *tab_bar_widget[T]) SelectTabItem(index int) {
	d.content.SelectTabItem(index)
}

func (d *tab_bar_widget[T]) SelectedTabItem() (index int, ok bool) {
	return d.content.SelectedTabItem()
}

func (d *tab_bar_widget[T]) OnTabClose(fn func(context *guigui.Context, index int, data TabItem[T])) {
	d.content.OnTabClose(fn)
}

func (d *tab_bar_widget[T]) OnTabSelect(fn func(context *guigui.Context, index int, data TabItem[T])) {
	d.content.OnTabSelect(fn)
}

func (d *tab_bar_widget[T]) OnTabSwitch(fn func(context *guigui.Context, index1 int, data1 TabItem[T], index2 int, data2 TabItem[T])) {
	d.content.OnTabSwitch(fn)
}

func (d *tab_bar_widget[T]) Build(context *guigui.Context, adder *guigui.ChildAdder) error {
	d.panel.SetContent(&d.content)
	d.panel.SetContentConstraints(basicwidget.PanelContentConstraintsFixedWidth)
	adder.AddWidget(&d.panel)
	return nil
}

func (d *tab_bar_widget[T]) Layout(context *guigui.Context, widgetBounds *guigui.WidgetBounds, layouter *guigui.ChildLayouter) {
	d.layout = guigui.LinearLayout{
		Direction: guigui.LayoutDirectionHorizontal,
		Items: []guigui.LinearLayoutItem{
			{
				Widget: &d.panel,
				Size:   guigui.FlexibleSize(1),
			},
		},
	}
	d.layout.LayoutWidgets(context, widgetBounds.Bounds(), layouter)
}

func (d *tab_bar_widget[T]) Measure(context *guigui.Context, constraints guigui.Constraints) image.Point {
	return d.content.Measure(context, constraints)
}
