package widgets

import (
	"image"
	"io/fs"

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
	//guigui.DefaultWidget

	basicwidget.TextInput
	context_menu_area           basicwidget.ContextMenuArea[int]
	context_menu_items_disables [6]bool

	additional_contextmenu_items []basicwidget.PopupMenuItem[int]
	on_additional_item_selected  func(context *guigui.Context, index int)

	text_content TextContent

	img_details *ImageDetails
	icon_name   string
	icon_fs     fs.FS
}

func (d *TextInput) SetIcon(icon_name string, fs fs.FS) error {
	if fs == nil {
		fs = icon_directory
	}
	if d.icon_name == icon_name && d.icon_fs == fs {
		return nil
	}

	details, err := OpenImage(icon_name, fs)
	if err != nil {
		return err
	}
	d.img_details = details
	d.icon_name = icon_name
	d.icon_fs = fs
	return nil
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

		text_input := &d.TextInput
		switch index {
		case ContextMenuItemCopy:
			text_input.Copy()
		case ContextMenuItemPaste:
			text_input.Paste()
		case ContextMenuItemCut:
			text_input.Cut()
		case ContextMenuItemSelectAll:
			text_input.SelectAll()
		case ContextMenuItemUndo:
			text_input.Undo()
		case ContextMenuItemRedo:
			text_input.Redo()
		}
	})
	adder.AddWidget(&d.context_menu_area)

	if d.img_details != nil {
		handle_icon_color_mod(context, d.img_details, d.icon_name, d.icon_fs, &d.TextInput)
		d.TextInput.SetIcon(d.img_details.Img)
	}
	return d.TextInput.Build(context, adder)
}

func (d *TextInput) Layout(context *guigui.Context, widgetBounds *guigui.WidgetBounds, layouter *guigui.ChildLayouter) {
	b := widgetBounds.Bounds()
	d.TextInput.Layout(context, widgetBounds, layouter)
	layouter.LayoutWidget(&d.context_menu_area, b)
}

func (d *TextInput) Measure(context *guigui.Context, constraints guigui.Constraints) image.Point {
	return d.TextInput.Measure(context, constraints)
}
