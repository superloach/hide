package main

import (
	"github.com/superloach/hide"
	"github.com/superloach/hide/util"
)

func main() {
	h := hide.MkHide()
	err := h.Run()
	if err != nil && err != util.Exit {
		panic(err)
	}
}
