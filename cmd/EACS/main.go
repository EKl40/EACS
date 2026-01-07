package main

import (
	"github.com/getlantern/systray"
	"github.com/EKl40/EACS/internal/tray"
)

func main() {
	systray.Run(tray.OnReady, nil)
}
