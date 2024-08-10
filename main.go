package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Starting Personal AI Model Project")

	// Collect data
	data, err := collectTextData("./personal_data")
	if err != nil {
		log.Fatalf("Error collecting data: %v", err)
	}

	// Preprocess data
	for i, text := range data {
		data[i] = preprocessText(text)
	}

	// Set up neural network
	setupNeuralNetwork()

	fmt.Println("Project steps completed. Ready for further development.")
}
