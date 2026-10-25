package widgets

import (
	"io/fs"

	"github.com/guigui-gui/guigui"
	"github.com/guigui-gui/guigui/basicwidget"
)

type Button struct {
	basicwidget.Button
	tooltip      basicwidget.TooltipArea
	tooltip_text string

	img_details *ImageDetails
	icon_name   string
	icon_fs     fs.FS
}

func (d *Button) SetIcon(icon_name string, fs fs.FS) error {
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

func (d *Button) SetTooltip(text string) {
	d.tooltip_text = text
}

func (d *Button) Build(context *guigui.Context, adder *guigui.ChildAdder) error {
	if d.tooltip_text != "" {
		d.tooltip.SetText(d.tooltip_text)
		adder.AddWidget(&d.tooltip)
	}

	if d.img_details != nil {
		handle_icon_color_mod(context, d.img_details, d.icon_name, d.icon_fs, &d.Button)
		d.Button.SetIcon(d.img_details.Img)
	}
	return d.Button.Build(context, adder)
}

func (d *Button) Layout(context *guigui.Context, widgetBounds *guigui.WidgetBounds, layouter *guigui.ChildLayouter) {
	d.Button.Layout(context, widgetBounds, layouter)
	if d.tooltip_text != "" {
		layouter.LayoutWidget(&d.tooltip, widgetBounds.Bounds())
	}
}
