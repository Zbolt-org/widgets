package widgets

import (
	"image"

	"github.com/guigui-gui/guigui"
	"github.com/guigui-gui/guigui/basicwidget"
)

type InfoHint struct {
	guigui.DefaultWidget
	tooltip_area basicwidget.TooltipArea
	tooltip_text string

	infohint_icon Icon

	layout guigui.LinearLayout
}

func (infohint *InfoHint) SetText(text string) {
	infohint.tooltip_text = text
}

func (infohint *InfoHint) Build(context *guigui.Context, adder *guigui.ChildAdder) error {
	infohint.tooltip_area.SetText(infohint.tooltip_text)
	adder.AddWidget(&infohint.tooltip_area)

	if !infohint.infohint_icon.HasIcon() {
		infohint.infohint_icon.Load("icons/info-tip.png", nil)
	}
	adder.AddWidget(&infohint.infohint_icon)
	return nil
}

func (infohint *InfoHint) Layout(context *guigui.Context, widgetBounds *guigui.WidgetBounds, layouter *guigui.ChildLayouter) {
	b := widgetBounds.Bounds()

	infohint.layout = guigui.LinearLayout{
		Direction: guigui.LayoutDirectionVertical,
		Items: []guigui.LinearLayoutItem{
			{
				Size: guigui.FlexibleSize(1),
			},
			{
				Layout: guigui.LinearLayout{
					Direction: guigui.LayoutDirectionHorizontal,
					Items: []guigui.LinearLayoutItem{
						{
							Size: guigui.FlexibleSize(1),
						},
						{
							Widget: &infohint.infohint_icon,
						},
						{
							Size: guigui.FlexibleSize(1),
						},
					},
				},
			},
			{
				Size: guigui.FlexibleSize(1),
			},
		},
	}

	tooltip_area_container_bounds := infohint.layout.ItemBoundsAt(1, context, b)
	layout := infohint.layout.Items[1].Layout.(guigui.LinearLayout)
	tooltip_area_bounds := layout.ItemBoundsAt(1, context, tooltip_area_container_bounds)
	layouter.LayoutWidget(&infohint.tooltip_area, tooltip_area_bounds)
	infohint.layout.LayoutWidgets(context, b, layouter)
}

func (infohint *InfoHint) Measure(context *guigui.Context, constraints guigui.Constraints) image.Point {
	return infohint.infohint_icon.Measure(context, constraints)
}
