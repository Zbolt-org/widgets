package widgets

import (
	"image"
	"io/fs"
	"widgets/utils"

	"github.com/guigui-gui/guigui"
	"github.com/guigui-gui/guigui/basicwidget"
	"github.com/guigui-gui/guigui/basicwidget/basicwidgetdraw"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type TabIcon struct {
	Icon string
	Fs   fs.FS
}

type TabItem[T comparable] struct {
	Text              string
	Icon              TabIcon
	Closable, Movable bool
	Widget            guigui.Widget
	Value             T
}

func (t *TabItem[T]) is_icon() bool {
	return t.Icon.Icon != ""
}

type tab_item_widget[T comparable] struct {
	guigui.DefaultWidget

	data TabItem[T]

	layout guigui.LinearLayout

	icon       Icon
	text       basicwidget.Text
	close_icon Icon

	was_hit_at_cursor bool
	selected          bool
}

func (d *tab_item_widget[T]) set_data(data TabItem[T]) {
	d.data = data
}

func (d *tab_item_widget[T]) get_data() TabItem[T] {
	return d.data
}

func (d *tab_item_widget[T]) icon_size(context *guigui.Context) image.Point {
	lh := basicwidget.LineHeight(context)
	lh = (lh / 2) + (lh / 4)
	return image.Pt(lh, lh)
}

func (d *tab_item_widget[T]) Build(context *guigui.Context, adder *guigui.ChildAdder) error {
	icon_size := d.icon_size(context)

	if d.data.Closable {
		d.close_icon.Load("icons/close.png", nil, icon_size)
		d.close_icon.SetClickable(true)
		adder.AddWidget(&d.close_icon)
	}

	if d.data.is_icon() {
		d.icon.Load(d.data.Icon.Icon, d.data.Icon.Fs, icon_size)
		adder.AddWidget(&d.icon)
	}

	// TODO: bold the text if the tab is selected.
	t := basicwidget.TextStyle{}
	t.SetBold(d.selected)
	d.text.SetBaseStyle(&t)
	d.text.SetValue(d.data.Text)
	d.text.SetEllipsisString("...")
	adder.AddWidget(&d.text)
	return nil
}

func tab_item_padding(context *guigui.Context) guigui.Padding {
	u := basicwidget.UnitSize(context)
	return guigui.Padding{
		Top:    u / 12,
		Bottom: u / 12,
		Start:  u / 3,
		End:    u / 3,
	}
}

func tab_item_radius(context *guigui.Context) int {
	padding := tab_item_padding(context)
	return (padding.Start + padding.Top) / 2
}

func (d *tab_item_widget[T]) Layout(context *guigui.Context, widgetBounds *guigui.WidgetBounds, layouter *guigui.ChildLayouter) {
	text_width := d.text.Measure(context, guigui.FixedHeightConstraints(basicwidget.LineHeight(context))).X
	max_text_width := basicwidget.UnitSize(context) * 12
	text_width = min(text_width, max_text_width)

	items := []guigui.LinearLayoutItem{
		{
			Widget: &d.text,
			Size:   guigui.FixedSize(text_width),
		},
	}

	if d.data.is_icon() {
		item := []guigui.LinearLayoutItem{
			{
				Size:   guigui.FixedSize(d.icon.Measure(context, guigui.Constraints{}).X),
				Widget: &d.icon,
			},
		}
		items = append(item, items...)

	}

	if d.data.Closable {
		items = append(items, guigui.LinearLayoutItem{
			Size:   guigui.FixedSize(d.close_icon.Measure(context, guigui.Constraints{}).X),
			Widget: &d.close_icon,
		})
	}

	padding := tab_item_padding(context)
	d.layout = guigui.LinearLayout{
		Direction: guigui.LayoutDirectionHorizontal,
		Gap:       utils.XSmallGap(context),
		Padding:   padding,
		Items:     items,
	}

	d.layout.LayoutWidgets(context, widgetBounds.Bounds(), layouter)
}

func (d *tab_item_widget[T]) HandlePointingInput(context *guigui.Context, widgetBounds *guigui.WidgetBounds) guigui.HandleInputResult {
	if widgetBounds.IsHitAtCursor() && inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {
		d.selected = !d.selected
		guigui.RequestRedraw(d)
		return guigui.HandleInputByWidget(d)
	} else if widgetBounds.IsHitAtCursor() || d.was_hit_at_cursor {
		d.was_hit_at_cursor = widgetBounds.IsHitAtCursor()
		guigui.RequestRedraw(d)
		return guigui.HandleInputByWidget(d)
	}
	return guigui.HandleInputResult{}
}

func (d *tab_item_widget[T]) Draw(context *guigui.Context, widgetBounds *guigui.WidgetBounds, dst *ebiten.Image) {
	r := tab_item_radius(context)
	b := widgetBounds.Bounds()

	background_clr := basicwidgetdraw.BackgroundSecondaryColor(context.ColorMode())
	border_type := basicwidgetdraw.RoundedRectBorderTypeOutset
	if d.selected {
		border_type = basicwidgetdraw.RoundedRectBorderTypeInset
	}
	basicwidgetdraw.DrawRoundedRect(context, dst, b, background_clr, r)

	clr1, clr2 := basicwidgetdraw.BorderColors(context.ColorMode(), border_type)
	width := 1 * float32(context.Scale())

	if widgetBounds.IsHitAtCursor() || d.selected {
		basicwidgetdraw.DrawRoundedRectBorder(context, dst, b, clr1, clr2, r, width, border_type)
		//basicwidgetdraw.DrawRoundedRectBorderWithSharpCorners(context, dst, b, clr1, clr2, r, width, border_type, corners)
	}
	// TODO: change the background color based on cursor hover & selected & drag.
}

func (d *tab_item_widget[T]) Measure(context *guigui.Context, _ guigui.Constraints) image.Point {
	return d.layout.Measure(context, guigui.Constraints{})
}

func (d *tab_item_widget[T]) CursorShape(context *guigui.Context, widgetBounds *guigui.WidgetBounds) (ebiten.CursorShapeType, bool) {
	if widgetBounds.IsHitAtCursor() {
		return ebiten.CursorShapePointer, true
	}
	return 0, false
}
