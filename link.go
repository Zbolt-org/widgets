package widgets

import (
	"net/url"

	opener "codeberg.org/udan-jayanith/Opener"
	"github.com/guigui-gui/guigui"
	"github.com/guigui-gui/guigui/basicwidget"
	"github.com/hajimehoshi/dialog"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var (
	DefualtURL_Style basicwidget.TextStyle = basicwidget.TextStyle{}
)

func init() {
	DefualtURL_Style.SetColor(basicwidget.AccentTintColor())
	DefualtURL_Style.SetUnderline(true)
}

// Note: This can alos open file paths so be makesure the url is what you want to open.
// If text value != url a tooltip will show the url otherwise the tooltip will not show.
type Link struct {
	basicwidget.Text
	u       string
	style   *basicwidget.TextStyle
	tooltip basicwidget.TooltipArea
}

func (d *Link) SetURL_Str(u string) {
	d.u = u
}

func (d *Link) SetURL(u *url.URL) error {
	uu := u.String()
	d.u = uu
	return nil
}

func (d *Link) SetBaseStyle(style *basicwidget.TextStyle) {
	d.style = style
}

func (d *Link) Build(context *guigui.Context, adder *guigui.ChildAdder) error {
	if d.style == nil {
		d.style = &DefualtURL_Style
	}
	d.Text.SetBaseStyle(d.style)

	if d.Text.Value() != d.u {
		d.tooltip.SetText(d.u)
		adder.AddWidget(&d.tooltip)
	}
	return d.Text.Build(context, adder)
}

func (d *Link) Layout(context *guigui.Context, widgetBounds *guigui.WidgetBounds, layouter *guigui.ChildLayouter) {
	b := widgetBounds.Bounds()
	if d.Text.Value() != d.u {
		layouter.LayoutWidget(&d.tooltip, b)
	}
	d.Text.Layout(context, widgetBounds, layouter)
}

func (d *Link) HandlePointingInput(context *guigui.Context, widgetBounds *guigui.WidgetBounds) guigui.HandleInputResult {
	if widgetBounds.IsHitAtCursor() && inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {
		err := opener.Open(d.u)
		if err != nil {
			msgb := dialog.Message(err.Error())
			go msgb.Error()
		}
	}
	return d.Text.HandlePointingInput(context, widgetBounds)
}
