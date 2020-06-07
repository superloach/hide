package main

import (
	"github.com/superloach/hide"
	"github.com/superloach/hide/easy"
)

func main() {
	h, err := hide.MkHide()
	if err != nil {
		panic(err)
	}

	err = h.Run()
	h.State.Save()
	if err == easy.Exit || err == nil {
		println("bye-bye")
	} else if err != nil {
		panic(err)
	}
}
