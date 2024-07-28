package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/iterator"
)

func shell() {
	// Init
	ctx := context.Background()
	client, model := initModel(ctx)
	defer client.Close()
	var iter *genai.GenerateContentResponseIterator
	scanner := bufio.NewScanner(os.Stdin)
	cs := model.StartChat()
	cs.History = []*genai.Content{}

	// Loop
	for {
		// Input
		var input string
		fmt.Printf("Q:>")
		if scanner.Scan() {
			input = scanner.Text()
		}

		// Process
		iter = cs.SendMessageStream(ctx, genai.Text(input))

		parts := []genai.Part{}

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
			fmt.Printf("%s", resp.Candidates[0].Content.Parts[0])
			parts = append(parts, resp.Candidates[0].Content.Parts[0])
		}
		fmt.Println()

		// Add history
		cs.History = append(cs.History, &genai.Content{
			Parts: []genai.Part{genai.Text(input)},
			Role:  "user",
		})
		cs.History = append(cs.History, &genai.Content{
			Parts: parts,
			Role:  "model",
		})
	}
}
