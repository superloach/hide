package scene;

import "github.com/hajimehoshi/ebiten";

// generic scene
type Scene interface {
	Update(*ebiten.Image);
}
