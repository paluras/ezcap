package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func captureScreen(path string) error {
	cmd := exec.Command("screencapture", "-x", path)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error capturing screen:", err)
		os.Exit(1)
	}
	return nil
}

func captureIOS(path string) error {
	cmd := exec.Command("xcrun", "simctl", "io", "booted", "screenshot", path)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error capturing iOS simulator:", err)
		os.Exit(1)
	}
	return nil
}

func captureAndroid(path string) error {
	cmd := exec.Command("adb", "exec-out", "screencap", "-p")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error capturing Android emulator:", err)
		os.Exit(1)
	}
	err = os.WriteFile(path, output, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		os.Exit(1)
	}
	return nil
}

func printPath(filename string) {
	dir, _ := os.Getwd()
	fullPath := filepath.Join(dir, filename)
	fmt.Println("Saved to:", fullPath)
}
