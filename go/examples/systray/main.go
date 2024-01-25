package main

import (
	_ "embed"
	"os"

	"github.com/getlantern/systray"
)

//go:embed favicon.ico
var icon []byte

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetIcon(icon)
	systray.SetTitle("Awesome App")
	systray.SetTooltip("Pretty awesome超级棒")
	mQuit := systray.AddMenuItem("Quit", "Quit the whole app")

	go func() {
		for {
			select {
			case <-mQuit.ClickedCh:
				os.Exit(0)
			}
		}
	}()

	// Sets the icon of a menu item. Only available on Mac and Windows.
	mQuit.SetIcon(icon)
}

func onExit() {
	// clean up here
}
