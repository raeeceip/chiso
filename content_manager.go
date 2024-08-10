package main

import (
	"sort"
	"strings"
	"time"
)

type ContentManager struct {
	Contents []Content
}

func NewContentManager(contents []Content) *ContentManager {
	return &ContentManager{Contents: contents}
}

func (cm *ContentManager) FilterByType(contentType ContentType) []Content {
	var filtered []Content
	for _, content := range cm.Contents {
		if content.Type == contentType {
			filtered = append(filtered, content)
		}
	}
	return filtered
}

func (cm *ContentManager) FilterByDateRange(start, end time.Time) []Content {
	var filtered []Content
	for _, content := range cm.Contents {
		if (content.Date.After(start) || content.Date.Equal(start)) && 
		   (content.Date.Before(end) || content.Date.Equal(end)) {
			filtered = append(filtered, content)
		}
	}
	return filtered
}

func (cm *ContentManager) SearchByKeyword(keyword string) []Content {
	var results []Content
	for _, content := range cm.Contents {
		if strings.Contains(strings.ToLower(content.Text), strings.ToLower(keyword)) {
			results = append(results, content)
		}
	}
	return results
}

func (cm *ContentManager) FilterByTags(tags []string) []Content {
	var filtered []Content
	for _, content := range cm.Contents {
		if containsAllTags(content.Tags, tags) {
			filtered = append(filtered, content)
		}
	}
	return filtered
}

func containsAllTags(contentTags, searchTags []string) bool {
	for _, searchTag := range searchTags {
		found := false
		for _, contentTag := range contentTags {
			if strings.EqualFold(searchTag, contentTag) {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func (cm *ContentManager) SortByDate() {
	sort.Slice(cm.Contents, func(i, j int) bool {
		return cm.Contents[i].Date.After(cm.Contents[j].Date)
	})
}
