package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func suffixToType(suffix string) (imgType string) {
	if suffix == "jpg" || suffix == "jpeg" || suffix == "JPG" {
		return "jpeg"
	} else if suffix == "png" || suffix == "PNG" {
		return "png"
	} else if suffix == "gif" {
		return "gif"
	}
	return ""
}

func addrToSuffix(addr string) (suffix string) {
	idx := strings.LastIndex(addr, ".")
	if idx == -1 {
		return ""
	}
	return addr[idx+1:]
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Type gemini 'Question' to ask questions")
		return
	}

	hasImage := os.Args[len(os.Args)-1] == "-i" || os.Args[len(os.Args)-1] == "--image"

	ctx := context.Background()

	// Access your API key as an environment variable (see "Set up your API key" above)
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash")

	var iter *genai.GenerateContentResponseIterator

	if !hasImage {
		iter = model.GenerateContentStream(ctx, genai.Text(os.Args[1]))
	} else {
		prompt := []genai.Part{}
		for idx, addr := range os.Args {
			if idx > 1 && idx < len(os.Args)-1 {
				imgData, err := os.ReadFile(addr)
				if err != nil {
					log.Fatal(err)
				}
				prompt = append(prompt, genai.ImageData(suffixToType(addrToSuffix(addr)), imgData))
			}
		}
		prompt = append(prompt, genai.Text(os.Args[1]))
		iter = model.GenerateContentStream(ctx, prompt...)
	}

	fmt.Println()

	for {
		resp, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		// print resp
		fmt.Printf("%s", resp.Candidates[0].Content.Parts[0])
	}

	fmt.Println()
}
