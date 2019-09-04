package main;

import "time";

import "hide/game";
import "hide/scene";
import "hide/images";

func main() {
	hide := game.MakeGame("HIDE", 600, 480);

	_, _ = scene.MakeTimedTextScene(hide, "welcome to the HIDE demo!\nthis should go away after 3 seconds", time.Second * 3);

	next, _ := scene.MakeTextScene(hide, "great, this is the next scene!\npress -> to view an image of mario :P");
	next.Back = -1;

	menu, _ := scene.MakeMenuScene(hide, "mario menu");

	mario1, m1id := scene.MakeFileImageScene(hide, "./images/mario.png");
	mario1.Caption += "\npress <- to return to the menu";
	mario1.Next = -1;
	menu.AddOption("mario file version", m1id);

	mario2, m2id := scene.MakeImageScene(hide, images.MarioPng);
	mario2.Caption += "./images/mario_png.go\npress <- to return to the menu";
	mario2.Next = -1;
	menu.AddOption("mario f2bs version", m2id);

	hide.Run();
}
