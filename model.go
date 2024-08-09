package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
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

	model = client.GenerativeModel(getModel())
	return client, model
}

func getModel() (modelName string) {
	// Get config file
	file, err := os.Open("currentModel.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Read config file
	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	currentModel := string(data)

	// Check available
	if !modelAvailable(currentModel) {
		currentModel = "gemini-1.5-flash"
		setModel(currentModel)
	}

	return currentModel
}

func setModel(modelName string) {
	// Check available
	if !modelAvailable(modelName) {
		fmt.Println("Model not available. Please set a new model.")
		return
	}

	// Get config file
	file, err := os.Create("currentModel.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Write config file
	_, err = file.WriteString(modelName)
	if err != nil {
		log.Fatal(err)
	}
}

func modelAvailable(modelName string) (available bool) {
	availableModels := getAvailableModels()
	for _, model := range availableModels {
		if modelName == model {
			return true
		}
	}
	return false
}

func getAvailableModels() (modelNames []string) {
	// Open the file for reading
	file, err := os.Open("availableModels.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a new Scanner for the file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Append each line to the modelNames slice
		modelNames = append(modelNames, scanner.Text())
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return modelNames
}

func printAvailableModels() {
	availableModels := getAvailableModels()
	for _, model := range availableModels {
		if getModel() == model {
			fmt.Println("*" + model)
		} else {
			fmt.Println(" " + model)
		}
	}
}
