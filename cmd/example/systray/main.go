package main

import (
	"fmt"
	"github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon"
)

func main() {
	fmt.Println("Start")
	systray.Run(onReady, onExit)
	fmt.Println("After run")
}

func onReady() {
	fmt.Println("onReady")
	systray.SetIcon(icon.Data)
	systray.SetTitle("Example Title")
	systray.SetTooltip("Example Tooltip")
	mQuit := systray.AddMenuItem("Quit", "Quit application")
	mQuit.SetIcon(icon.Data)

	go func() {
		<-mQuit.ClickedCh
		fmt.Println("Requesting quit")
		systray.Quit()
		fmt.Println("Finished quitting")
	}()
}

func onExit() {
	fmt.Println("onExit")
}
