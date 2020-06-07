package state

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
)

func Load() *State {
	cfg, err := os.UserConfigDir()
	if err != nil {
		println(err.Error())
		return BlankState
	}
	dir := path.Join(cfg, "hide")
	file := path.Join(dir, Filename)

	err = os.MkdirAll(dir, 0777)
	if err != nil {
		println(err.Error())
		return BlankState
	}

	f, err := os.OpenFile(file, os.O_RDONLY, 0777)
	if err != nil {
		println(err.Error())
		return BlankState
	}

	data, err := ioutil.ReadAll(f)
	if err != nil {
		println(err.Error())
		return BlankState
	}

	f.Close()

	c := &State{}

	err = json.Unmarshal(data, c)
	if err != nil {
		println(err.Error())
		return BlankState
	}

	return c
}
