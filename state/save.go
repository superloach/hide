package state

import (
	"encoding/json"
	"os"
	"path"
)

func (s *State) Save() {
	data, err := json.MarshalIndent(s, "", "\t")
	if err != nil {
		panic(err)
	}

	cfg, err := os.UserConfigDir()
	if err != nil {
		panic(err)
	}
	dir := path.Join(cfg, "hide")
	file := path.Join(dir, Filename)

	err = os.MkdirAll(dir, 0777)
	if err != nil {
		panic(err)
	}

	f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}

	_, err = f.Write(data)
	if err != nil {
		panic(err)
	}

	f.Close()
}
