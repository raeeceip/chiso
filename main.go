package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	fmt.Println("Starting Personal AI Model Project")

	// Initialize Enhanced Belief System Processor
	beliefProcessor := NewBeliefSystemProcessor("./beliefs", "./personal_data/journals")

	// Update beliefs from journals
	err := beliefProcessor.UpdateBeliefsFromJournals()
	if err != nil {
		log.Printf("Error updating beliefs from journals: %v", err)
	}

	// Save updated beliefs
	err = beliefProcessor.SaveBeliefs()
	if err != nil {
		log.Printf("Error saving updated beliefs: %v", err)
	}

	// Collect data
	data, err := collectTextData("./personal_data")
	if err != nil {
		log.Fatalf("Error collecting data: %v", err)
	}
	fmt.Printf("Collected %d content items\n", len(data))

	// Create ContentManager
	cm := NewContentManager(data)

	// Example usage of ContentManager
	journals := cm.FilterByType(Journal)
	fmt.Printf("Found %d journal entries\n", len(journals))

	recentContent := cm.FilterByDateRange(time.Now().AddDate(0, -1, 0), time.Now())
	fmt.Printf("Found %d content items from the last month\n", len(recentContent))

	aiRelatedContent := cm.SearchByKeyword("artificial intelligence")
	fmt.Printf("Found %d content items related to AI\n", len(aiRelatedContent))

	personalGrowthContent := cm.FilterByTags([]string{"personal_growth", "self_improvement"})
	fmt.Printf("Found %d content items tagged with personal growth or self-improvement\n", len(personalGrowthContent))

	cm.SortByDate() // Sort all content by date

	// Preprocess data
	processedData := preprocessAllData(cm.Contents, beliefProcessor)
	fmt.Println("Data preprocessing and belief system processing completed")

	// Set up and train neural network
	err = setupAndTrainNetwork(processedData)
	if err != nil {
		log.Fatalf("Error in neural network setup and training: %v", err)
	}

	fmt.Println("Neural network training completed")

	// Example of using the trained model with belief system
	exampleInput := "How can we promote ethical conduct in AI development?"
	response := generateResponseFromNetwork(exampleInput)
	processedResponse := beliefProcessor.ProcessResponse(exampleInput, response)
	fmt.Printf("Input: %s\nResponse: %s\n", exampleInput, processedResponse)

	// Start the interactive loop
	interactiveLoop(beliefProcessor)
}

func preprocessAllData(contents []Content, bsp *BeliefSystemProcessor) []string {
	processedData := make([]string, len(contents))
	for i, content := range contents {
		preprocessed := preprocessText(content.Text)
		processedData[i] = bsp.ProcessResponse(content.Text, preprocessed)
	}
	return processedData
}

func interactiveLoop(bsp *BeliefSystemProcessor) {
	fmt.Println("Enter your questions (type 'exit' to quit):")
	for {
		var input string
		fmt.Print("> ")
		fmt.Scanln(&input)

		if input == "exit" {
			break
		}

		response := generateResponseFromNetwork(input)
		processedResponse := bsp.ProcessResponse(input, response)
		fmt.Printf("AI: %s\n", processedResponse)
	}
	fmt.Println("Goodbye!")
}
