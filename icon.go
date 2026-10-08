package widgets

import (
	"embed"
	"fmt"
	"image"
	"io/fs"

	"github.com/guigui-gui/guigui"
	"github.com/guigui-gui/guigui/basicwidget"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/srwiley/oksvg"
	"github.com/srwiley/rasterx"
)

//go:embed icons/*.svg
var icon_directory embed.FS

type Icon struct {
	guigui.DefaultWidget

	size image.Point

	prevoious_color_mod ebiten.ColorMode
	ebiten_img          ebiten.Image
	image_widget basicwidget.Image
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

	svg, err := oksvg.ReadIconStream(file, oksvg.StrictErrorMode)
	if err != nil {
		return err
	}

	img := svg_to_raster(svg)
	icon_ref.ebiten_img = *ebiten.NewImageFromImage(img)
	return nil
}

func (icon_ref *Icon) Build(context *guigui.Context, adder *guigui.ChildAdder) error {
	return nil
}

func (icon_ref *Icon) Layout(context *guigui.Context, widgetBounds *guigui.WidgetBounds, layouter *guigui.ChildLayouter) {
	layouter.LayoutWidget(&icon_ref.image_widget, widgetBounds.Bounds())
}

func (icon_ref *Icon) Measure(context *guigui.Context, constraints guigui.Constraints) image.Point {
	return icon_ref.size
}
