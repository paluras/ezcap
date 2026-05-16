package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: ezcap screen | ios | android")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "screen":
		captureScreen("screenshot.png")
		printPath("screenshot.png")
	case "ios":
		captureIOS("screenshoot_ios.png")
		printPath("screenshoot_ios.png")

	case "android":
		captureAndroid("screenshot_android.png")
		printPath("screenshot_android.png")
	}
}
