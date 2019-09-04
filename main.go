package main

import "time"

import "hide/game"
import "hide/scene"
import "hide/images"

import "github.com/jakecoffman/cp"

func main() {
	hide := game.MakeGame("HIDE", 600, 480)

	_, _ = scene.MakeTimedTextScene(hide, "welcome to the HIDE demo!\nthis should go away after 3 seconds", time.Second*3)

	next, _ := scene.MakeTextScene(hide, "great, this is the next scene!\npress -> to view an image of mario :P")
	next.Back = -1

	menu, _ := scene.MakeMenuScene(hide, "mario menu")

	mario := images.ByteImage(images.MarioPng)

	mario1, m1id := scene.MakeImageScene(hide, images.FileImage("./images/mario.png"))
	mario1.Caption = "mario.png\npress <- to return to the menu"
	mario1.Next = -1
	menu.AddOption("mario file version", m1id)

	mario2, m2id := scene.MakeImageScene(hide, mario)
	mario2.Caption += "images.MarioPng\npress <- to return to the menu"
	mario2.Next = -1
	menu.AddOption("mario f2bs version", m2id)

	mario3, m3id := scene.MakeChipmunkScene(hide)
	mario3.MakeRect(mario, float64(hide.Width()) / 2, float64(hide.Height()) / 2, 100)
	mario3.MakeRect(mario, float64(hide.Width()) / 2, float64(hide.Height()), cp.INFINITY)
	menu.AddOption("mario chipmunk demo", m3id)

	hide.Run()
}
