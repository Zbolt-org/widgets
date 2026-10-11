package widgets

import (
	"image"

	"github.com/guigui-gui/guigui"
	"github.com/guigui-gui/guigui/basicwidget"
)

type context_menu_item = int

const (
	ContextMenuItemCopy context_menu_item = iota + 0
	ContextMenuItemPaste
	ContextMenuItemCut
	ContextMenuItemSelectAll
	ContextMenuItemUndo
	ContextMenuItemRedo
	DefualtContextMenuItemEnd
)

type TextContent int

const (
	TextContentPlainText TextContent = iota + 0
	TextContentHTML
)

// TextInput is a widget for text input with context menu support.
type TextInput struct {
	guigui.DefaultWidget

	text_input                  basicwidget.TextInput
	context_menu_area           basicwidget.ContextMenuArea[int]
	context_menu_items_disables [6]bool

	additional_contextmenu_items []basicwidget.PopupMenuItem[int]
	on_additional_item_selected  func(context *guigui.Context, index int)

	text_content TextContent
}

func (d *TextInput) TextInput() *basicwidget.TextInput {
	return &d.text_input
}

func (d *TextInput) SetTextContent(textcontent TextContent) {
	d.text_content = textcontent
}

func (d *TextInput) TextContent() TextContent {
	return d.text_content
}

// OnAdditionalItemSelected sets the callback for when an additional context menu item is selected.
func (d *TextInput) OnAdditionalItemSelected(fn func(context *guigui.Context, index int)) {
	d.on_additional_item_selected = fn
}

// SetAditionalItems sets the additional context menu items.
func (d *TextInput) SetAditionalItems(addition_items []basicwidget.PopupMenuItem[int]) {
	d.additional_contextmenu_items = addition_items
}

// Only for defualt items. Sets whether the item is disabled.
func (d *TextInput) SetDisable(item context_menu_item, disable bool) {
	if item >= DefualtContextMenuItemEnd {
		panic("SetDisable only valid for defualt context menu items")
	}
	d.context_menu_items_disables[item] = disable
}

// Only for defualt items. Returns whether the item is disabled.
func (d *TextInput) Disable(item context_menu_item) bool {
	if item >= DefualtContextMenuItemEnd {
		panic("SetDisable only valid for defualt context menu items")
	}
	return d.context_menu_items_disables[item]
}

func (d *TextInput) Build(context *guigui.Context, adder *guigui.ChildAdder) error {
	adder.AddWidget(&d.text_input)

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
			text_input := d.TextInput()
			text_input.Copy()
		case ContextMenuItemPaste:
			text_input := d.TextInput()
			text_input.Paste()
		case ContextMenuItemCut:
			text_input := d.TextInput()
			text_input.Cut()
		case ContextMenuItemSelectAll:
			text_input := d.TextInput()
			text_input.SelectAll()
		case ContextMenuItemUndo:
			text_input := d.TextInput()
			text_input.Undo()
		case ContextMenuItemRedo:
			text_input := d.TextInput()
			text_input.Redo()
		}
	})
	adder.AddWidget(&d.context_menu_area)
	return nil
}

func (d *TextInput) Layout(context *guigui.Context, widgetBounds *guigui.WidgetBounds, layouter *guigui.ChildLayouter) {
	b := widgetBounds.Bounds()
	layouter.LayoutWidget(&d.text_input, b)
	layouter.LayoutWidget(&d.context_menu_area, b)
}

func (d *TextInput) Measure(context *guigui.Context, constraints guigui.Constraints) image.Point {
	return d.text_input.Measure(context, constraints)
}
