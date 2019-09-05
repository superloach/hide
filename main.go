package main

import "time"

import "github.com/superloach/hide/game"
import "github.com/superloach/hide/scene"
import "github.com/superloach/hide/images"

import "github.com/jakecoffman/cp"

func main() {
	hide := game.MakeGame("HIDE", 600, 480)

	_, _ = scene.MakeTimedTextScene(hide, "welcome to the HIDE demo!\nthis should go away after 3 seconds", time.Second*3)

	next, _ := scene.MakeTextScene(hide, "great, this is the next scene!\npress -> to view an image of mario :P")
	next.Back = -1

	menu, _ := scene.MakeMenuScene(hide, "mario menu")

	mario := images.ByteImage(images.MarioPng)
	ball := images.ByteImage(images.BallPng)

	mario1, m1id := scene.MakeImageScene(hide, images.FileImage("./images/mario.png"))
	mario1.Caption = "mario.png\npress <- to return to the menu"
	mario1.Next = -1

	menu.AddOption("mario file version", m1id)

	mario2, m2id := scene.MakeImageScene(hide, mario)
	mario2.Caption += "images.MarioPng\npress <- to return to the menu"
	mario2.Next = -1

	menu.AddOption("mario f2bs version", m2id)

	mario3, m3id := scene.MakeChipScene(hide)

	_ = mario3.MakeRect(mario, 100, 480-128, cp.INFINITY)

	var count float64 = 32
	var i float64
	for i = 0; i < count; i++ {
		_ = mario3.MakeCirc(ball, 100+8*float64(int(i)%4), 0+8*i, 10)
	}

	menu.AddOption("mario chipmunk demo", m3id)

	hide.Run()
}
