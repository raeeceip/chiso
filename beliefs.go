package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type BeliefCategory int

const (
	FundamentalPrinciples BeliefCategory = iota
	KnowledgeAndTruth
	PersonalGrowth
	InterpersonalRelations
	SocietalResponsibility
	PersonalValues
	DecisionMaking
)

type Belief struct {
	Category    BeliefCategory
	Description string
	Source      string
	CreatedAt   time.Time
}

type BeliefSystemProcessor struct {
	Beliefs        []Belief
	BeliefFilePath string
	JournalPath    string
}

func NewBeliefSystemProcessor(beliefDirPath, journalPath string) *BeliefSystemProcessor {
	bsp := &BeliefSystemProcessor{
		BeliefFilePath: beliefDirPath,
		JournalPath:    journalPath,
	}
	bsp.LoadBeliefs()
	return bsp
}

func (bsp *BeliefSystemProcessor) LoadBeliefs() error {
	files, err := ioutil.ReadDir(bsp.BeliefFilePath)
	if err != nil {
		return fmt.Errorf("error reading belief directory: %v", err)
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".txt" {
			filePath := filepath.Join(bsp.BeliefFilePath, file.Name())
			beliefs, err := bsp.readBeliefFile(filePath)
			if err != nil {
				fmt.Printf("Error reading belief file %s: %v\n", file.Name(), err)
				continue
			}
			bsp.Beliefs = append(bsp.Beliefs, beliefs...)
		}
	}
	return nil
}

func (bsp *BeliefSystemProcessor) readBeliefFile(filePath string) ([]Belief, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var beliefs []Belief
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			continue
		}
		category := bsp.stringToCategory(strings.TrimSpace(parts[0]))
		description := strings.TrimSpace(parts[1])
		beliefs = append(beliefs, Belief{
			Category:    category,
			Description: description,
			Source:      filepath.Base(filePath),
			CreatedAt:   time.Now(),
		})
	}
	return beliefs, scanner.Err()
}

func (bsp *BeliefSystemProcessor) stringToCategory(s string) BeliefCategory {
	switch strings.ToLower(s) {
	case "fundamental principles":
		return FundamentalPrinciples
	case "knowledge and truth":
		return KnowledgeAndTruth
	case "personal growth":
		return PersonalGrowth
	case "interpersonal relations":
		return InterpersonalRelations
	case "societal responsibility":
		return SocietalResponsibility
	case "personal values":
		return PersonalValues
	case "decision making":
		return DecisionMaking
	default:
		return FundamentalPrinciples
	}
}

func (bsp *BeliefSystemProcessor) ProcessResponse(input string, response string) string {
	for _, belief := range bsp.Beliefs {
		if strings.Contains(strings.ToLower(input), strings.ToLower(belief.Description[:10])) {
			response = bsp.ApplyBelief(belief, response)
		}
	}
	return response
}

func (bsp *BeliefSystemProcessor) ApplyBelief(belief Belief, response string) string {
	switch belief.Category {
	case FundamentalPrinciples:
		return fmt.Sprintf("Considering our fundamental principle: %s, %s", belief.Description, response)
	case KnowledgeAndTruth:
		return fmt.Sprintf("Based on our commitment to truth and knowledge: %s, %s", belief.Description, response)
	case PersonalGrowth:
		return fmt.Sprintf("In the spirit of personal growth (%s), %s", belief.Description, response)
	case InterpersonalRelations:
		return fmt.Sprintf("Keeping in mind the importance of healthy relationships (%s), %s", belief.Description, response)
	case SocietalResponsibility:
		return fmt.Sprintf("Considering our responsibility to society (%s), %s", belief.Description, response)
	case PersonalValues:
		return fmt.Sprintf("In alignment with our personal values (%s), %s", belief.Description, response)
	case DecisionMaking:
		return fmt.Sprintf("Taking a thoughtful approach to decision making (%s), %s", belief.Description, response)
	default:
		return response
	}
}

func (bsp *BeliefSystemProcessor) UpdateBeliefsFromJournals() error {
	files, err := ioutil.ReadDir(bsp.JournalPath)
	if err != nil {
		return fmt.Errorf("error reading journal directory: %v", err)
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".txt" {
			filePath := filepath.Join(bsp.JournalPath, file.Name())
			err := bsp.processJournalFile(filePath)
			if err != nil {
				fmt.Printf("Error processing journal file %s: %v\n", file.Name(), err)
			}
		}
	}
	return nil
}

func (bsp *BeliefSystemProcessor) processJournalFile(filePath string) error {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	// This is a simple example. In a real implementation, you'd use more sophisticated NLP techniques.
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		if strings.Contains(strings.ToLower(line), "i believe") {
			belief := Belief{
				Category:    PersonalValues, // Default category, could be more sophisticated
				Description: strings.TrimSpace(strings.TrimPrefix(line, "I believe")),
				Source:      filepath.Base(filePath),
				CreatedAt:   time.Now(),
			}
			bsp.Beliefs = append(bsp.Beliefs, belief)
		}
	}
	return nil
}

func (bsp *BeliefSystemProcessor) SaveBeliefs() error {
	file, err := os.Create(filepath.Join(bsp.BeliefFilePath, "generated_beliefs.txt"))
	if err != nil {
		return err
	}
	defer file.Close()

	for _, belief := range bsp.Beliefs {
		_, err := fmt.Fprintf(file, "%s: %s\n", bsp.categoryToString(belief.Category), belief.Description)
		if err != nil {
			return err
		}
	}
	return nil
}

func (bsp *BeliefSystemProcessor) categoryToString(category BeliefCategory) string {
	switch category {
	case FundamentalPrinciples:
		return "Fundamental Principles"
	case KnowledgeAndTruth:
		return "Knowledge and Truth"
	case PersonalGrowth:
		return "Personal Growth"
	case InterpersonalRelations:
		return "Interpersonal Relations"
	case SocietalResponsibility:
		return "Societal Responsibility"
	case PersonalValues:
		return "Personal Values"
	case DecisionMaking:
		return "Decision Making"
	default:
		return "Uncategorized"
	}
}
