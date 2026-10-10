package widgets

import (
	"image"
	"widgets/utils"

	"github.com/guigui-gui/guigui"
	"github.com/guigui-gui/guigui/basicwidget"
)

type action_notification_content struct {
	guigui.DefaultWidget

	layout guigui.LinearLayout

	action_value string
	action       basicwidget.Text
}

func (d *action_notification_content) SetAction(action string) {
	d.action_value = action
}

func (d *action_notification_content) Build(context *guigui.Context, adder *guigui.ChildAdder) error {
	d.action.SetValue(d.action_value)
	adder.AddWidget(&d.action)
	return nil
}

func (d *action_notification_content) Layout(context *guigui.Context, widgetBounds *guigui.WidgetBounds, layouter *guigui.ChildLayouter) {
	d.layout = guigui.LinearLayout{
		Padding:   utils.MediumPadding(context),
		Direction: guigui.LayoutDirectionHorizontal,
		Items: []guigui.LinearLayoutItem{
			{
				Size: guigui.FlexibleSize(1),
			},
			{
				Layout: utils.VerticalCenter(&d.action),
			},
			{
				Size: guigui.FlexibleSize(1),
			},
		},
	}
}

func (d *action_notification_content) Measure(context *guigui.Context, constraints guigui.Constraints) image.Point {
	return d.layout.Measure(context, constraints)
}

type ActionNotification struct {
	guigui.DefaultWidget

	action_content action_notification_content
	popup          basicwidget.Popup

	layout guigui.LinearLayout
}

func (d *ActionNotification) Build(context *guigui.Context, adder *guigui.ChildAdder) error {
	d.action_content.SetAction("Hello world")

	d.popup.SetModal(true)
	d.popup.SetAnimated(true)
	if !d.popup.IsOpen() {
		d.popup.SetOpen(true)
	}
	d.popup.SetContent(&d.action_content)
	adder.AddWidget(&d.popup)
	return nil
}

func (d *ActionNotification) Layout(context *guigui.Context, widgetBounds *guigui.WidgetBounds, layouter *guigui.ChildLayouter) {
	b := context.AppBounds()
	d.layout = guigui.LinearLayout{
		Direction: guigui.LayoutDirectionHorizontal,
		Items: []guigui.LinearLayoutItem{
			{
				Size: guigui.FlexibleSize(1),
			},
			{
				Layout: guigui.LinearLayout{
					Direction: guigui.LayoutDirectionVertical,
					Items: []guigui.LinearLayoutItem{
						{
							Size: guigui.FlexibleSize(1),
						},
						{
							Widget: &d.popup,
						},
						{
							Size: guigui.FixedSize(basicwidget.UnitSize(context) * 2),
						},
					},
				},
			},
			{
				Size: guigui.FlexibleSize(1),
			},
		},
	}
	d.layout.LayoutWidgets(context, b, layouter)
}

func (d *ActionNotification) Measure(context *guigui.Context, _ guigui.Constraints) image.Point {
	return d.action_content.Measure(context, guigui.Constraints{})
	/*
		if len(d.layout.Items) == 0 {
			return image.Point{}
		}

		b := context.AppBounds()
		outer_layout_bounds := d.layout.ItemBoundsAt(1, context, b)
		inner_layout := d.layout.Items[1].Layout.(guigui.LinearLayout)

		return inner_layout.ItemBoundsAt(1, context, outer_layout_bounds).Size()
	*/
}
