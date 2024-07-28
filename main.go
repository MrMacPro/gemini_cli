package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func initModel(ctx context.Context) (client *genai.Client, model *genai.GenerativeModel) {
	// Access your API key as an environment variable (see "Set up your API key" above)
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	//defer client.Close()

	model = client.GenerativeModel("gemini-1.5-flash")
	return client, model
}

func main() {
	// Help
	if os.Args[len(os.Args)-1] == "-h" || os.Args[len(os.Args)-1] == "--help" {
		fmt.Println("Gemini CLI by Hanson Zhang")
		fmt.Println("	Run \"gemini\" to spawn a conversation shell.")
		fmt.Println("	Run \"gemini '<YOUR QUESTION>'\" to ask gemini.")
		fmt.Println("	Run \"gemini '<YOUR QUESTION>' '<IMG PATH>' ... -i\" to ask gemini with images. You can put multiple images here.")
		return
	}

	if len(os.Args) < 2 {
		shell()
	} else {
		oneTime(os.Args)
	}
}
