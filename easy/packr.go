package easy

import (
	"bytes"
	"encoding/json"
	"image/png"

	"github.com/gobuffalo/packr/v2"
	"github.com/hajimehoshi/ebiten"

	"github.com/superloach/hide/level"
)

var ImageBox = packr.New("ImageBox", "./../assets/images")
var LevelBox = packr.New("LevelBox", "./../assets/levels")

func LoadImage(name string) (*ebiten.Image, error) {
	data, err := ImageBox.Find(name + ".png")
	if err != nil {
		return nil, err
	}
	read := bytes.NewReader(data)
	img, err := png.Decode(read)
	if err != nil {
		return nil, err
	}
	return ebiten.NewImageFromImage(img, ebiten.FilterDefault)
}

func LoadLevel(name string) (*level.Level, error) {
	data, err := LevelBox.Find(name + ".json")
	if err != nil {
		return nil, err
	}

	l := &level.Level{}
	err = json.Unmarshal(data, l)
	if err != nil {
		return l, err
	}

	for id, name := range l.Tiles {
		img, err := LoadImage(name)
		if err != nil {
			return l, err
		}
		l.TileImages[id] = img
	}

	return l, nil
}
