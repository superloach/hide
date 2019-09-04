package images

import "bytes"
import "image"
import "io/ioutil"

import "github.com/superloach/hide/hiderr"

import "github.com/hajimehoshi/ebiten"

func ByteImage(imageBytes []byte) *ebiten.Image {
	img, _, err := image.Decode(bytes.NewReader(imageBytes))
	hiderr.Do(err)
	image, _ := ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	return image
}

func FileImage(filename string) *ebiten.Image {
	dat, err := ioutil.ReadFile(filename)
	hiderr.Do(err)
	return ByteImage([]byte(dat))
}
