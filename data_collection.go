package main

import (
	"io/ioutil"
	"path/filepath"
)

func collectTextData(directory string) ([]string, error) {
	var data []string

	files, err := ioutil.ReadDir(directory)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".txt" {
			content, err := ioutil.ReadFile(filepath.Join(directory, file.Name()))
			if err != nil {
				return nil, err
			}
			data = append(data, string(content))
		}
	}

	return data, nil
}
