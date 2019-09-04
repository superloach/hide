package scene

import "hide/hiderr"

import "io/ioutil"

func MakeFileImageScene(game GameI, filename string) (*ImageScene, int) {
	dat, err := ioutil.ReadFile(filename)
	hiderr.Do(err)

	s, id := MakeImageScene(game, []byte(dat))
	s.Caption = filename

	return s, id
}
