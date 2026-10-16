package widgets

import (
	"github.com/guigui-gui/guigui"
	"github.com/guigui-gui/guigui/basicwidget"
)

type Button struct {
	basicwidget.Button
	tooltip      basicwidget.TooltipArea
	tooltip_text string
}

func (d *Button) SetTooltip(text string) {
	d.tooltip_text = text
}

func (d *Button) Build(context *guigui.Context, adder *guigui.ChildAdder) error {
	if d.tooltip_text != "" {
		d.tooltip.SetText(d.tooltip_text)
		adder.AddWidget(&d.tooltip)
	}
	return d.Button.Build(context, adder)
}

func (d *Button) Layout(context *guigui.Context, widgetBounds *guigui.WidgetBounds, layouter *guigui.ChildLayouter) {
	d.Button.Layout(context, widgetBounds, layouter)
	if d.tooltip_text != "" {
		layouter.LayoutWidget(&d.tooltip, widgetBounds.Bounds())
	}
}
