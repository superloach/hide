package images

import "image"
import "bytes"

import "github.com/superloach/hide/hiderr"

import "github.com/gobuffalo/packr"
import "github.com/hajimehoshi/ebiten"

func ByteImage(imageBytes []byte) *ebiten.Image {
	img, _, err := image.Decode(bytes.NewReader(imageBytes))
	hiderr.Do(err)
	image, _ := ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	return image
}

func Image(name string) *ebiten.Image {
	dat, err := packr.NewBox(".").Find(name)
	hiderr.Do(err)
	return ByteImage([]byte(dat))
}
