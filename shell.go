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

	// Loop
	for {
		// Input
		var input string
		fmt.Printf("Q:>")
		if scanner.Scan() {
			input = scanner.Text()
		}

		// Process
		iter = model.GenerateContentStream(ctx, genai.Text(input))

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
		}
		fmt.Println()
	}
}
