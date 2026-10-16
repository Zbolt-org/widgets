package widgets

import (
	"image/png"
	"io/fs"
	"weak"

	"github.com/hajimehoshi/ebiten/v2"
)

type image_cache_key struct {
	file_path string
	fs        fs.FS
}

type ImageDetails struct {
	Img     *ebiten.Image
	Details any
	Cache   bool
}

var (
	image_cache map[image_cache_key]weak.Pointer[ImageDetails]
)

func init() {
	image_cache = make(map[image_cache_key]weak.Pointer[ImageDetails], 8)
}

// TODO: use program path as defualt fs.FS if the fs is nil.
func OpenImage(file_path string, fs fs.FS) (*ImageDetails, error) {
	k := image_cache_key{
		file_path: file_path,
		fs:        fs,
	}

	v, ok := image_cache[k]
	if ok && v.Value() != nil {
		vv := v.Value()
		vv.Cache = true
		return vv, nil
	}

	file, err := fs.Open(file_path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		return nil, err
	}
	ebiten_img := ebiten.NewImageFromImage(img)
	img_details := ImageDetails{
		Img: ebiten_img,
	}
	image_cache[k] = weak.Make(&img_details)
	return &img_details, nil
}

func ChangeImageDetails(file_path string, fs fs.FS, details any) (ok bool) {
	k := image_cache_key{
		file_path: file_path,
		fs:        fs,
	}

	v, ok := image_cache[k]
	if ok && v.Value() != nil {
		vv := v.Value()
		vv.Details = details
	}
	return ok
}
