package main

import (
	"fmt"
	"os"
)

func main() {
	// Options
	if len(os.Args) < 2 {
		shell()
	} else if os.Args[len(os.Args)-1] == "-h" || os.Args[len(os.Args)-1] == "--help" {
		fmt.Println("Gemini CLI by Hanson Zhang")
		fmt.Println("	Run \"gemini\" to spawn a conversation shell.")
		fmt.Println("	Run \"gemini '<YOUR QUESTION>'\" to ask gemini.")
		fmt.Println("	Run \"gemini '<YOUR QUESTION>' '<IMG PATH>' ... -i\" to ask gemini with images. You can put multiple images here.")
		fmt.Println()
		fmt.Println("	Run \"gemini -m\" or \"gemini --model\" to get all available models.")
		fmt.Println("	Run \"gemini -m '<MODEL NAME>'\" or \"gemini --model '<MODEL NAME>'\" to set model.")
		return
	} else if os.Args[1] == "-m" || os.Args[1] == "--model" {
		if len(os.Args) > 2 {
			setModel(os.Args[2])
		}
		printAvailableModels()
	} else {
		oneTime(os.Args)
	}
}
