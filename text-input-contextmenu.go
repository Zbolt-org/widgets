package widgets

import (
	"image"

	"github.com/guigui-gui/guigui"
	"github.com/guigui-gui/guigui/basicwidget"
)

type InputContextMenuInterface interface {
	Copy() bool
	Paste() bool
	Cut() bool
	SelectAll()
	Undo() bool
	Redo() bool
}

type InputContextMenu[T guigui.Widget] struct {
	guigui.DefaultWidget

	widget LazyWidget[T]
	input  InputContextMenuInterface

	context_menu_area           basicwidget.ContextMenuArea[int]
	context_menu_items_disables [6]bool

	additional_contextmenu_items []basicwidget.PopupMenuItem[int]
	on_additional_item_selected  func(context *guigui.Context, index int)
}

// OnAdditionalItemSelected sets the callback for when an additional context menu item is selected.
func (d *InputContextMenu[T]) OnAdditionalItemSelected(fn func(context *guigui.Context, index int)) {
	d.on_additional_item_selected = fn
}

// SetAditionalItems sets the additional context menu items.
func (d *InputContextMenu[T]) SetAditionalItems(addition_items []basicwidget.PopupMenuItem[int]) {
	d.additional_contextmenu_items = addition_items
}

// Only for defualt items. Sets whether the item is disabled.
func (d *InputContextMenu[T]) SetDisable(item context_menu_item, disable bool) {
	if item >= DefualtContextMenuItemEnd {
		panic("SetDisable only valid for defualt context menu items")
	}
	d.context_menu_items_disables[item] = disable
}

// Only for defualt items. Returns whether the item is disabled.
func (d *InputContextMenu[T]) Disable(item context_menu_item) bool {
	if item >= DefualtContextMenuItemEnd {
		panic("SetDisable only valid for defualt context menu items")
	}
	return d.context_menu_items_disables[item]
}

func (d *InputContextMenu[T]) Widget() T {
	return d.widget.Widget()
}

func (d *InputContextMenu[T]) SetInput(input InputContextMenuInterface) {
	d.input = input
}

func (d *InputContextMenu[T]) Build(context *guigui.Context, adder *guigui.ChildAdder) error {
	context_menu := d.context_menu_area.PopupMenu()
	contextmenu_items := []basicwidget.PopupMenuItem[int]{
		{
			Text:     "Copy",
			KeyText:  "Ctrl+C",
			Disabled: d.context_menu_items_disables[ContextMenuItemCopy],
			Value:    ContextMenuItemCopy,
		},
		{
			Text:     "Paste",
			KeyText:  "Ctrl+V",
			Disabled: d.context_menu_items_disables[ContextMenuItemPaste],
			Value:    ContextMenuItemPaste,
		},
		{
			Text:     "Cut",
			KeyText:  "Ctrl+X",
			Disabled: d.context_menu_items_disables[ContextMenuItemCut],
			Value:    ContextMenuItemCut,
		},
		{
			Text:     "Select All",
			KeyText:  "Ctrl+A",
			Disabled: d.context_menu_items_disables[ContextMenuItemSelectAll],
			Value:    ContextMenuItemSelectAll,
		},
		{
			Text:     "Undo",
			KeyText:  "Ctrl+Z",
			Disabled: d.context_menu_items_disables[ContextMenuItemUndo],
			Value:    ContextMenuItemUndo,
		},
		{
			Text:     "Redo",
			KeyText:  "Ctrl+Y",
			Disabled: d.context_menu_items_disables[ContextMenuItemRedo],
			Value:    ContextMenuItemRedo,
		},
	}

	if len(d.additional_contextmenu_items) > 0 {
		contextmenu_items = append(contextmenu_items, basicwidget.PopupMenuItem[int]{
			Border: true,
		})
		contextmenu_items = append(contextmenu_items, d.additional_contextmenu_items...)
	}
	context_menu.SetItems(contextmenu_items)
	context_menu.OnItemSelected(func(context *guigui.Context, index int) {
		if index >= DefualtContextMenuItemEnd {
			if d.on_additional_item_selected == nil {
				return
			}
			d.on_additional_item_selected(context, index-DefualtContextMenuItemEnd)
			return
		}

		switch index {
		case ContextMenuItemCopy:
			d.input.Copy()
		case ContextMenuItemPaste:
			d.input.Paste()
		case ContextMenuItemCut:
			d.input.Cut()
		case ContextMenuItemSelectAll:
			d.input.SelectAll()
		case ContextMenuItemUndo:
			d.input.Undo()
		case ContextMenuItemRedo:
			d.input.Redo()
		}
	})
	adder.AddWidget(&d.context_menu_area)
	return nil
}

func (d *InputContextMenu[T]) Layout(context *guigui.Context, widgetBounds *guigui.WidgetBounds, layouter *guigui.ChildLayouter) {
	b := widgetBounds.Bounds()
	layouter.LayoutWidget(&d.context_menu_area, b)
	layouter.LayoutWidget(d.Widget(), b)
}

func (d *InputContextMenu[T]) Measure(context *guigui.Context, constraints guigui.Constraints) image.Point {
	return d.Widget().Measure(context, constraints)
}
