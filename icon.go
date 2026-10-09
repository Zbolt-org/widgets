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

type Icon struct {
	guigui.DefaultWidget

	size image.Point

	is_black     bool
	ebiten_img   *ebiten.Image
	image_widget basicwidget.Image
}

func (icon_ref *Icon) HasIcon() bool {
	return icon_ref.ebiten_img != nil
}

func (icon_ref *Icon) Load(icon string, fs fs.FS, size ...image.Point) error {
	if fs == nil {
		fs = icon_directory
	}

	if len(size) > 1 {
		panic("Extra arguments")
	} else if len(size) == 1 {
		icon_ref.size = size[0]
	}

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
	return nil
}

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
		icon_ref.image_widget.SetImage(icon_ref.ebiten_img)
	} else if context.ColorMode() == ebiten.ColorModeDark && icon_ref.is_black {
		inverse_color(icon_ref.ebiten_img)
		icon_ref.is_black = false
	}

	adder.AddWidget(&icon_ref.image_widget)
	return nil
}

func (icon_ref *Icon) Layout(context *guigui.Context, widgetBounds *guigui.WidgetBounds, layouter *guigui.ChildLayouter) {
	layouter.LayoutWidget(&icon_ref.image_widget, widgetBounds.Bounds())
}

func (icon_ref *Icon) Measure(context *guigui.Context, constraints guigui.Constraints) image.Point {
	if icon_ref.size.X == 0 || icon_ref.size.Y == 0 {
		line_height := basicwidget.LineHeight(context)
		return image.Pt(line_height, line_height)
	}
	return icon_ref.size
}
