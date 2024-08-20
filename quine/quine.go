package quine

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/anthropic-ai/anthropic-sdk-go"
)

func readDirectory() (map[string]string, error) {
	files := make(map[string]string)

	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() && info.Name() == "vendor" {
			return filepath.SkipDir
		}

		if info.IsDir() {
			return nil
		}

		content, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		files[path] = string(content)
		return nil
	})

	return files, err
}

func prepareForClaude(files map[string]string) string {
	var sb strings.Builder

	sb.WriteString("Please analyze and reproduce the following Go project structure with the given file contents. " +
		"After reproducing the structure, provide a brief analysis of the project's purpose and structure, " +
		"and suggest any improvements or best practices that could be applied:\n\n")

	for path, content := range files {
		sb.WriteString(fmt.Sprintf("File: %s\n", path))
		sb.WriteString("Content:\n")
		sb.WriteString(content)
		sb.WriteString("\n\n")
	}

	return sb.String()
}

func sendToClaudeAPI(message string) (string, error) {
	apiKey := os.Getenv("ANTHROPIC_API_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("ANTHROPIC_API_KEY environment variable not set")
	}

	client, err := anthropic.NewClient(apiKey)
	if err != nil {
		return "", fmt.Errorf("error creating Anthropic client: %v", err)
	}

	resp, err := client.CreateCompletion(
		context.Background(),
		&anthropic.CreateCompletionRequest{
			Model:     anthropic.GPT4,
			MaxTokens: 1000,
			Prompt:    fmt.Sprintf("Human: %s\n\nAssistant:", message),
		},
	)
	if err != nil {
		return "", fmt.Errorf("error creating completion: %v", err)
	}

	return resp.Completion, nil
}

func main() {
	files, err := readDirectory()
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		return
	}

	message := prepareForClaude(files)

	fmt.Println("Sending message to Claude API...")
	response, err := sendToClaudeAPI(message)
	if err != nil {
		fmt.Printf("Error sending to Claude API: %v\n", err)
		return
	}

	fmt.Println("Response from Claude:")
	fmt.Println(response)

	err = ioutil.WriteFile("claude_response.txt", []byte(response), 0644)
	if err != nil {
		fmt.Printf("Error writing response to file: %v\n", err)
	} else {
		fmt.Println("Response written to claude_response.txt")
	}
}
