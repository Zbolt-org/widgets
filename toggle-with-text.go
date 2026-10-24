package widgets

import (
	"image"
	"github.com/Zbolt-org/widgets/utils"

	"github.com/guigui-gui/guigui"
	"github.com/guigui-gui/guigui/basicwidget"
)

type LabeledToggle struct {
	guigui.DefaultWidget

	label        string
	label_widget basicwidget.Text
	toggel       basicwidget.Toggle

	layout guigui.LinearLayout
}

func (d *LabeledToggle) Toggle() *basicwidget.Toggle {
	return &d.toggel
}

func (d *LabeledToggle) SetLabel(label string) {
	d.label = label
}

func (d *LabeledToggle) Build(context *guigui.Context, adder *guigui.ChildAdder) error {
	d.label_widget.SetValue(d.label)
	adder.AddWidget(&d.label_widget)

	adder.AddWidget(&d.toggel)
	return nil
}

func (d *LabeledToggle) Layout(context *guigui.Context, widgetBounds *guigui.WidgetBounds, layouter *guigui.ChildLayouter) {
	d.layout = guigui.LinearLayout{
		Gap:       utils.SmallGap(context),
		Direction: guigui.LayoutDirectionHorizontal,
		Items: []guigui.LinearLayoutItem{
			{
				Layout: utils.VerticalCenter(&d.label_widget),
			},
			{
				Layout: utils.VerticalCenter(&d.toggel),
			},
		},
	}

	d.layout.LayoutWidgets(context, widgetBounds.Bounds(), layouter)
}

func (d *LabeledToggle) Measure(context *guigui.Context, constraints guigui.Constraints) image.Point {
	return d.layout.Measure(context, constraints)
}
