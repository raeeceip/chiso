package quine

import (
	"fmt"
	"io/ioutil"
	"runtime"
)

func Quine() (string, error) {
	// Get the file name and line number of the current function
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return "", fmt.Errorf("unable to get caller information")
	}

	// Read the entire file
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("error reading file: %v", err)
	}

	// Convert the content to a string and return it
	return string(content), nil
}

func main() {
	content, err := Quine()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println("Quine output:")
	fmt.Println(content)

	// Optionally, write the content to a file
	err = ioutil.WriteFile("quine_output.txt", []byte(content), 0644)
	if err != nil {
		fmt.Printf("Error writing to file: %v\n", err)
	} else {
		fmt.Println("Quine output written to quine_output.txt")
	}
}
