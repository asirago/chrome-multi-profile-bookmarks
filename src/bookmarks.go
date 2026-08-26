package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func resolveBookmarksFile(profileDir string) (string, error) {
	candidates := []string{"AccountBookmarks", "Bookmarks"}

	for _, filename := range candidates {
		candidate := filepath.Join(profileDir, filename)
		_, err := os.Stat(candidate)
		if err == nil {
			return candidate, nil
		}
		if !errors.Is(err, os.ErrNotExist) {
			return "", fmt.Errorf("failed to access %s: %w", candidate, err)
		}
	}

	return "", fmt.Errorf(
		"no bookmarks file found in %s (tried AccountBookmarks and Bookmarks)",
		profileDir,
	)
}

func getURLs(profileIndexes []int) map[string][]*bookmarkURL {
	urls := map[string][]*bookmarkURL{}

	for _, profileIndex := range profileIndexes {
		profileName := getProfileName(profileIndex)
		profileDir := filepath.Join(bookmarksDir, profileName)
		bookmarksFilePath, err := resolveBookmarksFile(profileDir)
		if err != nil {
			log.Fatalf("failed to locate bookmarks file for %s: %v", profileName, err)
		}

		b, err := os.ReadFile(bookmarksFilePath)
		if err != nil {
			log.Fatalf(
				"failed to read bookmarks file for %s: %v",
				profileName,
				err,
			)
		}

		bookmarksFile := bookmarksFile{}

		err = json.Unmarshal(b, &bookmarksFile)
		if err != nil {
			log.Fatalf("failed to unmarshal bookmarks json: %v", err)
		}

		for _, rootNode := range bookmarksFile.Roots {
			collectBookmarks(rootNode, rootNode, urls, profileIndex)
		}

	}

	return urls
}

func collectBookmarks(
	parent, node *bookmarkNode,
	urls map[string][]*bookmarkURL,
	profileIndex int,
) {
	if node == nil {
		return
	}

	if node.URL != "" {
		hash := fmt.Sprintf("%s-%s", parent.Name, parent.DateAdded)
		url := bookmarkURL{
			ProfileIndex: profileIndex,
			Folder:       parent.Name,
			FolderHash:   hash,
			Name:         node.Name,
			URL:          node.URL,
		}

		urls[hash] = append(urls[hash], &url)
	}

	for _, child := range node.Children {
		collectBookmarks(node, child, urls, profileIndex)
	}
}
