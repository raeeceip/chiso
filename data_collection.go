package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type ContentType string

const (
	Journal    ContentType = "journal"
	BlogPost   ContentType = "blog_post"
	SocialPost ContentType = "social_post"
	Reference  ContentType = "reference"
)

type Content struct {
	Type      ContentType
	Text      string
	Date      time.Time
	Tags      []string
	SourceURL string // For references or social media posts
}

func collectTextData(directory string) ([]Content, error) {
	var allContent []Content

	// Walk through all subdirectories
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(info.Name(), ".txt") {
			content, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}

			// Parse filename for metadata
			fileName := info.Name()
			datePart := strings.Split(fileName, "_")[0]
			date, _ := time.Parse("2006-01-02", datePart)

			// Determine content type based on directory
			var contentType ContentType
			if strings.Contains(path, "journals") {
				contentType = Journal
			} else if strings.Contains(path, "blog_posts") {
				contentType = BlogPost
			} else if strings.Contains(path, "social_media") {
				contentType = SocialPost
			} else if strings.Contains(path, "references") {
				contentType = Reference
			}

			// Extract tags from content (assuming tags are in the format #tag)
			tags := extractTags(string(content))

			allContent = append(allContent, Content{
				Type: contentType,
				Text: string(content),
				Date: date,
				Tags: tags,
			})
		}

		return nil
	})

	return allContent, err
}

func extractTags(text string) []string {
	var tags []string
	words := strings.Fields(text)
	for _, word := range words {
		if strings.HasPrefix(word, "#") {
			tags = append(tags, strings.TrimPrefix(word, "#"))
		}
	}
	return tags
}
