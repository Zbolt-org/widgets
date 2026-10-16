package widgets

import (
	"embed"
	"image"
	"io/fs"

	"github.com/guigui-gui/guigui"
	"github.com/guigui-gui/guigui/basicwidget"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

//go:embed icons/*
var icon_directory embed.FS

// TODO: Add an transparency option for icon
// TODO: Add a cache for icons.
type Icon struct {
	guigui.DefaultWidget

	size image.Point

	img_details  *ImageDetails
	image_widget basicwidget.Image

	load_inputs struct {
		icon string
		fs   fs.FS
		size image.Point
	}

	clickable   bool
	on_click_fn func(context *guigui.Context, icon_ref *Icon)
}

func (icon_ref *Icon) OnClick(fn func(context *guigui.Context, icon_ref *Icon)) {
	icon_ref.on_click_fn = fn
}

func (icon_ref *Icon) Clickable() bool {
	return icon_ref.clickable
}

func (icon_ref *Icon) SetClickable(clickable bool) {
	icon_ref.clickable = clickable
}

func (icon_ref *Icon) HasIcon() bool {
	return icon_ref.img_details != nil
}

func (icon_ref *Icon) fs() fs.FS {
	if icon_ref.load_inputs.fs == nil {
		return icon_directory
	}
	return icon_ref.load_inputs.fs
}

func (icon_ref *Icon) Load(icon string, fs fs.FS, size image.Point) error {
	if icon_ref.load_inputs.icon == icon && icon_ref.load_inputs.fs == fs && icon_ref.load_inputs.size == size {
		return nil
	}
	icon_ref.load_inputs.icon = icon
	icon_ref.load_inputs.fs = fs
	icon_ref.load_inputs.size = size

	if fs == nil {
		fs = icon_directory
	}
	if size.X != 0 && size.Y != 0 {
		icon_ref.size = size
	}

	img_details, err := OpenImage(icon, fs)
	if err != nil {
		return err
	}
	icon_ref.img_details = img_details
	//change_color(icon_ref.ebiten_img, basicwidgetdraw.TextColor(ebiten.ColorModeLight, true))
	return nil
}

func (icon_ref *Icon) is_black() bool {
	if icon_ref.img_details == nil {
		return false
	}
	black, ok := icon_ref.img_details.Details.(bool)
	if !ok {
		return false
	}
	return black
}

func (icon_ref *Icon) set_is_black(black bool) (ok bool) {
	return ChangeImageDetails(icon_ref.load_inputs.icon, icon_ref.fs(), black)
}

/*
	func color_to_rgba_8bit(clr color.Color) (uint8, uint8, uint8, uint8) {
		r, g, b, a := clr.RGBA()
		return uint8(r), uint8(g), uint8(b), uint8(a)
	}

	func change_color(img *ebiten.Image, clr color.Color) {
		bounds := img.Bounds()
		pixels := make([]byte, 4*bounds.Dx()*bounds.Dy())
		img.ReadPixels(pixels)
		r, g, b, a := color_to_rgba_8bit(clr)
		transparency_diff := 255 - a
		for i := 4; i <= len(pixels); i += 4 {
			if pixels[i-1] == 0 {
				continue
			}
			rgba := pixels[i-4 : i]
			rgba[0] = r
			rgba[1] = g
			rgba[2] = b
			if rgba[3] >= transparency_diff {
				rgba[3] -= transparency_diff
			}
		}
		img.WritePixels(pixels)
	}
*/
func inverse_color(img *ebiten.Image) {
	bounds := img.Bounds()
	pixels := make([]byte, 4*bounds.Dx()*bounds.Dy())
	img.ReadPixels(pixels)
	for i := 4; i <= len(pixels); i += 4 {
		if pixels[i-1] == 0 {
			continue
		}
		rgba := pixels[i-4 : i]
		rgba[0] = ^rgba[0]
		rgba[1] = ^rgba[1]
		rgba[2] = ^rgba[2]
	}
	img.WritePixels(pixels)
}

func (icon_ref *Icon) Build(context *guigui.Context, adder *guigui.ChildAdder) error {
	if context.ColorMode() == ebiten.ColorModeLight && !icon_ref.is_black() {
		inverse_color(icon_ref.img_details.Img)
		icon_ref.set_is_black(true)
		guigui.RequestRedraw(&icon_ref.image_widget)
	} else if context.ColorMode() == ebiten.ColorModeDark && icon_ref.is_black() {
		inverse_color(icon_ref.img_details.Img)
		icon_ref.set_is_black(false)
		guigui.RequestRedraw(&icon_ref.image_widget)
	}

	icon_ref.image_widget.SetImage(icon_ref.img_details.Img)
	adder.AddWidget(&icon_ref.image_widget)
	return nil
}

func (icon_ref *Icon) Layout(context *guigui.Context, widgetBounds *guigui.WidgetBounds, layouter *guigui.ChildLayouter) {
	layouter.LayoutWidget(&icon_ref.image_widget, widgetBounds.Bounds())
}

func (icon_ref *Icon) Measure(context *guigui.Context, constraints guigui.Constraints) image.Point {
	if icon_ref.size.X == 0 && icon_ref.size.Y == 0 {
		line_height := basicwidget.LineHeight(context)
		return image.Pt(line_height, line_height)
	}
	return icon_ref.size
}

func (icon_ref *Icon) HandlePointingInput(context *guigui.Context, widgetBounds *guigui.WidgetBounds) guigui.HandleInputResult {
	if widgetBounds.IsHitAtCursor() && inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) && icon_ref.on_click_fn != nil && icon_ref.clickable {
		icon_ref.on_click_fn(context, icon_ref)
		return guigui.HandleInputByWidget(icon_ref)
	}
	return guigui.HandleInputResult{}
}

func (icon_ref *Icon) CursorShape(context *guigui.Context, widgetBounds *guigui.WidgetBounds) (ebiten.CursorShapeType, bool) {
	if icon_ref.clickable && widgetBounds.IsHitAtCursor() {
		return ebiten.CursorShapePointer, true
	}
	return 0, false
}
