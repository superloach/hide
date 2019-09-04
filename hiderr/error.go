package hiderr;

import "os";
import "errors";

func Do(err error) {
	if err != nil {
		panic(err);
	}
}

func Msg(msg string) {
	Do(errors.New(msg));
}

func Quit() {
	println("bye-bye :)");
	os.Exit(0);
}
