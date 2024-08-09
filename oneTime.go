package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/iterator"
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

func oneTime(args []string) {
	// Init
	ctx := context.Background()
	client, model := initModel(ctx)
	defer client.Close()
	var iter *genai.GenerateContentResponseIterator

	// Process
	hasImage := args[len(args)-1] == "-i" || args[len(args)-1] == "--image"
	if !hasImage {
		iter = model.GenerateContentStream(ctx, genai.Text(args[1]))
	} else {
		prompt := []genai.Part{}
		for idx, addr := range args {
			if idx > 1 && idx < len(args)-1 {
				imgData, err := os.ReadFile(addr)
				if err != nil {
					log.Fatal(err)
				}
				prompt = append(prompt, genai.ImageData(suffixToType(addrToSuffix(addr)), imgData))
			}
		}
		prompt = append(prompt, genai.Text(args[1]))
		iter = model.GenerateContentStream(ctx, prompt...)
	}

	// Print
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
		_, err = fmt.Printf("%s", resp.Candidates[0].Content.Parts[0])
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println()
}
