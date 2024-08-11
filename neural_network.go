package main

import (
	"fmt"
	"log"
	"strings"

	"gorgonia.org/gorgonia"
	"gorgonia.org/tensor"
)

var (
	model *gorgonia.VM
	g     *gorgonia.ExprGraph
	x     *gorgonia.Node
	y     *gorgonia.Node
)

func setupAndTrainNetwork(data []string) error {
	// Check if we have any data
	if len(data) == 0 {
		return fmt.Errorf("no data provided for training")
	}

	// For simplicity, let's assume each piece of data is a fixed size of 100 words
	// In a real scenario, you'd need to implement proper tokenization and padding
	inputSize := 100
	batchSize := len(data)

	// Create a simple feedforward network
	g = gorgonia.NewGraph()

	// Input layer
	x = gorgonia.NewMatrix(g,
		tensor.Float64,
		gorgonia.WithShape(batchSize, inputSize),
		gorgonia.WithName("x"),
	)

	// Hidden layer
	hiddenSize := 50
	w1 := gorgonia.NewMatrix(g,
		tensor.Float64,
		gorgonia.WithShape(inputSize, hiddenSize),
		gorgonia.WithName("w1"),
		gorgonia.WithInit(gorgonia.GlorotU(1.0)),
	)

	b1 := gorgonia.NewMatrix(g,
		tensor.Float64,
		gorgonia.WithShape(1, hiddenSize),
		gorgonia.WithName("b1"),
		gorgonia.WithInit(gorgonia.Zeroes()),
	)

	// Output layer
	outputSize := 10 // Assuming 10 possible output classes
	w2 := gorgonia.NewMatrix(g,
		tensor.Float64,
		gorgonia.WithShape(hiddenSize, outputSize),
		gorgonia.WithName("w2"),
		gorgonia.WithInit(gorgonia.GlorotU(1.0)),
	)

	b2 := gorgonia.NewMatrix(g,
		tensor.Float64,
		gorgonia.WithShape(1, outputSize),
		gorgonia.WithName("b2"),
		gorgonia.WithInit(gorgonia.Zeroes()),
	)

	// Define the computation
	var err error
	var hidden, output *gorgonia.Node

	// Forward pass
	if hidden, err = gorgonia.Mul(x, w1); err != nil {
		return fmt.Errorf("hidden = x*w1 error: %v", err)
	}
	if hidden, err = gorgonia.Add(hidden, b1); err != nil {
		return fmt.Errorf("hidden = hidden+b1 error: %v", err)
	}
	if hidden, err = gorgonia.Rectify(hidden); err != nil {
		return fmt.Errorf("hidden = rectify(hidden) error: %v", err)
	}

	if output, err = gorgonia.Mul(hidden, w2); err != nil {
		return fmt.Errorf("output = hidden*w2 error: %v", err)
	}
	if output, err = gorgonia.Add(output, b2); err != nil {
		return fmt.Errorf("output = output+b2 error: %v", err)
	}

	// Define symbolic y
	y = gorgonia.NewMatrix(g,
		tensor.Float64,
		gorgonia.WithShape(batchSize, outputSize),
		gorgonia.WithName("y"),
	)

	// Define loss function
	losses, err := gorgonia.Sub(output, y)
	if err != nil {
		return fmt.Errorf("losses = output-y error: %v", err)
	}

	square, err := gorgonia.Square(losses)
	if err != nil {
		return fmt.Errorf("square error: %v", err)
	}

	cost, err := gorgonia.Mean(square)
	if err != nil {
		return fmt.Errorf("cost = mean(square) error: %v", err)
	}

	// Create VM and Solver
	model = gorgonia.NewTapeMachine(g, gorgonia.BindDualValues(w1, w2))
	solver := gorgonia.NewRMSPropSolver(gorgonia.WithLearnRate(0.01))

	inputData := tensor.New(tensor.WithShape(batchSize, inputSize), tensor.WithBacking(convertToFloat64Slice(data, inputSize)))

	// Training loop
	for i := 0; i < 100; i++ { // Reduced number of iterations for testing
		if err := model.RunAll(); err != nil {
			log.Printf("Failed at iteration %d: %v", i, err)
			return err
		}

		// Create a new tensor node with the input data
		inputNode := gorgonia.NodeFromAny(g, inputData, gorgonia.WithName("input"))

		// Set the value of x to the input node
		if err := gorgonia.Let(x, inputNode); err != nil {
			return fmt.Errorf("failed to set x: %v", err)
		}

		if err := model.RunAll(); err != nil {
			return fmt.Errorf("failed to run: %v", err)
		}

		if err := solver.Step(gorgonia.NodesToValueGrads(gorgonia.Nodes{w1, w2})); err != nil {
			return fmt.Errorf("failed to solve: %v", err)
		}
		model.Reset() // Reset is required for CUDA-based graphs
	}

	fmt.Println("Neural network training completed")
	return nil
}

// Helper function to convert string data to float64 slice
func convertToFloat64Slice(data []string, inputSize int) []float64 {
	result := make([]float64, len(data)*inputSize)
	for i, text := range data {
		// Simple conversion: use ASCII values of characters
		// In a real scenario, you'd use proper text vectorization
		for j, char := range text {
			if j < inputSize {
				result[i*inputSize+j] = float64(char)
			} else {
				break
			}
		}
	}
	return result
}

func generateResponseFromNetwork(input string) string {
	// Convert input to float64 slice
	inputData := convertToFloat64Slice([]string{input}, 100) // Assuming input size of 100

	// Create a new tensor with the input data
	inputTensor := tensor.New(tensor.WithShape(1, 100), tensor.WithBacking(inputData))

	// Create a new node with the input tensor
	inputNode := gorgonia.NodeFromAny(g, inputTensor, gorgonia.WithName("input"))

	// Set the value of x to the input node
	if err := gorgonia.Let(x, inputNode); err != nil {
		log.Printf("Error setting input: %v", err)
		return "Error generating response"
	}

	// Run the model
	if err := model.RunAll(); err != nil {
		log.Printf("Error running model: %v", err)
		return "Error generating response"
	}

	// Get the output
	output, err := y.Value().Data().([]float64)
	if err {
		log.Printf("Error getting output: %v", err)
		return "Error generating response"
	}

	// Convert output to a response string (this is a simplistic approach)
	response := convertOutputToString(output)

	model.Reset()

	return response
}

func convertOutputToString(output []float64) string {
	// This is a very simplistic conversion. In a real scenario, you'd use a more sophisticated method.
	words := []string{"The", "quick", "brown", "fox", "jumps", "over", "the", "lazy", "dog"}
	var response []string
	for _, val := range output {
		index := int(val * float64(len(words)))
		if index >= len(words) {
			index = len(words) - 1
		}
		response = append(response, words[index])
	}
	return strings.Join(response, " ")
}
