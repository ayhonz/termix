package main

import "github.com/ayhonz/termix/editor"

func main() {
	e := editor.Init()	
	defer e.Quit()
	e.Run()
}
