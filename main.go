package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: ezcap screen | ios | android")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "screen":

		cmd := exec.Command("screencapture", "-x", "screenshot.png")
		err := cmd.Run()
		if err != nil {
			fmt.Println("error", err)
			os.Exit(1)
		}
	case "ios":

		cmd := exec.Command("xcrun", "simctl", "screenshot", "screenshot_simulator.png")
		err := cmd.Run()
		if err != nil {
			fmt.Println("error", err)
			os.Exit(1)

		}
	case "android":
		cmd := exec.Command("adb", "exec-out", "screencap", "-p")
		output, err := cmd.Output()
		if err != nil {
			fmt.Println("error", err)
			os.Exit(1)
		}
		err = os.WriteFile("screenshot_android.png", output, 0644)
		if err != nil {
			fmt.Println("error writing file:", err)
			os.Exit(1)
		}
	}
}
