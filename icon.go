package widgets

import (
	"embed"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io/fs"

	"github.com/guigui-gui/guigui"
	"github.com/guigui-gui/guigui/basicwidget"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/srwiley/oksvg"
	"github.com/srwiley/rasterx"
)

//go:embed icons/*
var icon_directory embed.FS

type Icon struct {
	guigui.DefaultWidget

	size image.Point

	prevoious_color_mod ebiten.ColorMode
	ebiten_img          *ebiten.Image
	image_widget        basicwidget.Image
}

func svg_to_raster(svg *oksvg.SvgIcon) image.Image {
	w, h := int(svg.ViewBox.W), int(svg.ViewBox.H)
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	scanner := rasterx.NewScannerGV(w, h, img, img.Bounds())
	raster := rasterx.NewDasher(w, h, scanner)

	svg.Draw(raster, 1.0)
	return img
}

func (icon_ref *Icon) Load(icon string, fs fs.FS, size ...image.Point) error {
	if fs == nil {
		fs = icon_directory
	}

	if len(size) > 1 {
		return fmt.Errorf("Extra arguments")
	} else if len(size) == 0 {
		icon_ref.size.X = 24
		icon_ref.size.Y = 24
	} else {
		icon_ref.size = size[0]
	}

	file, err := fs.Open(icon)
	if err != nil {
		return err
	}
	defer file.Close()

	img, err := png.Decode(file)
	icon_ref.ebiten_img = ebiten.NewImageFromImage(img)
	icon_ref.prevoious_color_mod = ebiten.ColorModeUnknown
	return nil
}

func replace_color(img *ebiten.Image, search, replace_with color.RGBA) {
	size := img.Bounds().Size()
	for y := 0; y <= size.Y; y++ {
		// TODO: Improve this. This currently check 1 pixel at a time. Check multiple pixels at a time by using ReadAt
		for x := 0; x <= size.X; x++ {
			if img.At(x, y) == search {
				img.Set(x, y, replace_with)
			}
		}
	}
}

func (icon_ref *Icon) Build(context *guigui.Context, adder *guigui.ChildAdder) error {
	if icon_ref.prevoious_color_mod != context.ColorMode() {
		switch context.ColorMode() {
		case ebiten.ColorModeDark:
			replace_color(icon_ref.ebiten_img, color.RGBA{0, 0, 0, 255}, color.RGBA{255, 255, 255, 255})
		case ebiten.ColorModeLight, ebiten.ColorModeUnknown:
			replace_color(icon_ref.ebiten_img, color.RGBA{255, 255, 255, 255}, color.RGBA{0, 0, 0, 255})
		}
		icon_ref.prevoious_color_mod = context.ColorMode()
	}
	icon_ref.image_widget.SetImage(icon_ref.ebiten_img)
	adder.AddWidget(&icon_ref.image_widget)
	return nil
}

func (icon_ref *Icon) Layout(context *guigui.Context, widgetBounds *guigui.WidgetBounds, layouter *guigui.ChildLayouter) {
	layouter.LayoutWidget(&icon_ref.image_widget, widgetBounds.Bounds())
}

func (icon_ref *Icon) Measure(context *guigui.Context, constraints guigui.Constraints) image.Point {
	return icon_ref.size
}
