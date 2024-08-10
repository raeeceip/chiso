package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	fmt.Println("Starting Personal AI Model Project")

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
	processedData := make([]string, len(cm.Contents))
	for i, content := range cm.Contents {
		processedData[i] = preprocessText(content.Text)
	}
	fmt.Println("Data preprocessing completed")

	// Set up and train neural network
	err = setupAndTrainNetwork(processedData)
	if err != nil {
		log.Fatalf("Error in neural network setup and training: %v", err)
	}

	fmt.Println("Project steps completed. Neural network is ready for use.")
}
