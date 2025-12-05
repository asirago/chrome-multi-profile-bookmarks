package main

import (
	"fmt"
	"strconv"
	"strings"
)

// getProfileName gets the Profile name from the index
func getProfileName(profileIndex int) string {
	if profileIndex == 0 {
		return "Default"
	}

	return fmt.Sprintf("Profile %d", profileIndex)
}

// csvToIntArray gets int array from strings that comma-separeted values of integers
func csvToIntArray(s string) ([]int, error) {
	strs := strings.Split(s, ",")
	profilesArr := []int{}

	if s == "" {
		return []int{0, 1}, nil
	}

	for _, str := range strs {
		i, err := strconv.Atoi(str)
		if err != nil {
			return []int{}, fmt.Errorf("failed to get profiles from csv: %v", err)
		}

		profilesArr = append(profilesArr, i)
	}

	return profilesArr, nil
}

// calculateScore computes a relevance score for a bookmark based on query terms
func calculateScore(folder, name, url string, queries []string) int {
	score := 0
	folderLower := strings.ToLower(folder)
	nameLower := strings.ToLower(name)
	urlLower := strings.ToLower(url)

	for _, query := range queries {
		query = strings.TrimSpace(query)
		if query == "" {
			continue
		}
		queryLower := strings.ToLower(query)

		// Name Match
		if nameLower == queryLower {
			score += 30
		} else if strings.HasPrefix(nameLower, queryLower) {
			score += 15
		} else if isWordStart(nameLower, queryLower) {
			score += 8
		} else if strings.Contains(nameLower, queryLower) {
			score += 3
		}

		// Folder Match
		if folderLower == queryLower {
			score += 25
		} else if strings.HasPrefix(folderLower, queryLower) {
			score += 12
		} else if isWordStart(folderLower, queryLower) {
			score += 6
		} else if strings.Contains(folderLower, queryLower) {
			score += 2
		}

		// URL Match
		if strings.HasPrefix(urlLower, "https://"+queryLower) ||
			strings.HasPrefix(urlLower, "https://www."+queryLower) {
			score += 10
		} else if isWordStart(urlLower, queryLower) {
			score += 4
		} else if strings.Contains(urlLower, queryLower) {
			score += 1
		}
	}

	return score
}

// isWordStart checks if the query matches at a word start
func isWordStart(text, query string) bool {
	separators := []string{" ", "-", "_", "/", "."}
	for _, sep := range separators {
		if strings.Contains(text, sep+query) {
			return true
		}
	}
	return false
}
