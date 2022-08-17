package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func main() {

	var OS string = runtime.GOOS
	apk := os.Args[1]
	var ni bool = false
	out, err := exec.Command("apktool").Output()
	fmt.Println("your input is", apk)
	if err != nil {
		ni = false
	} else {
		ni = true
	}

	if ni == true {
		switch OS {
		case "windows":
			fmt.Println("You are using windows\nDownload the apktool and run")
		case "linux":
			fmt.Println("You are using linux")
			cmd := exec.Command("apktool", "d", apk).Run()
			fmt.Print("cat ", apk, "/res/values/strings.xml\n")
			_ = cmd
			_ = out
			_ = ni
		}
	} else {
		fmt.Println("Apktool is not installed in your machine")
	}

}
