package widgets

import (
	"embed"
	"image"
	"image/png"
	"io/fs"

	"github.com/guigui-gui/guigui"
	"github.com/guigui-gui/guigui/basicwidget"
	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed icons/*
var icon_directory embed.FS

// TODO: Add an transparency option for icon
// TODO: Add a cache for icons.
type Icon struct {
	guigui.DefaultWidget

	size image.Point

	is_black     bool
	ebiten_img   *ebiten.Image
	image_widget basicwidget.Image

	load_inputs struct {
		icon string
		fs   fs.FS
		size image.Point
	}

	clickable bool
}

func (icon_ref *Icon) Clickable() bool {
	return icon_ref.clickable
}

func (icon_ref *Icon) SetClickable(clickable bool) {
	icon_ref.clickable = clickable
}

func (icon_ref *Icon) HasIcon() bool {
	return icon_ref.ebiten_img != nil
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
	// TODO: Consider a cache
	file, err := fs.Open(icon)
	if err != nil {
		return err
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		return err
	}
	icon_ref.ebiten_img = ebiten.NewImageFromImage(img)
	//change_color(icon_ref.ebiten_img, basicwidgetdraw.TextColor(ebiten.ColorModeLight, true))
	return nil

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
	if context.ColorMode() == ebiten.ColorModeLight && !icon_ref.is_black {
		inverse_color(icon_ref.ebiten_img)
		icon_ref.is_black = true
		guigui.RequestRedraw(&icon_ref.image_widget)
	} else if context.ColorMode() == ebiten.ColorModeDark && icon_ref.is_black {
		inverse_color(icon_ref.ebiten_img)
		icon_ref.is_black = false
		guigui.RequestRedraw(&icon_ref.image_widget)
	}

	icon_ref.image_widget.SetImage(icon_ref.ebiten_img)
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

func (icon_ref *Icon) CursorShape(context *guigui.Context, widgetBounds *guigui.WidgetBounds) (ebiten.CursorShapeType, bool) {
	if icon_ref.clickable && widgetBounds.IsHitAtCursor() {
		return ebiten.CursorShapePointer, true
	}
	return 0, false
}
