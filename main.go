package main

import "github.com/ayhonz/termix/editor"

func main() {
	var editor editor.Editor // do I have to do this?

	defer editor.Quit()
	editor.Run()
}
