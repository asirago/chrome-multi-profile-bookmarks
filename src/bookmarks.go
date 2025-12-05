package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func getURLs(profileIndexes []int) map[string][]*bookmarkURL {
	urls := map[string][]*bookmarkURL{}

	for _, profileIndex := range profileIndexes {

		bookmarksFilePath := filepath.Join(bookmarksDir, getProfileName(profileIndex), "Bookmarks")
		b, err := os.ReadFile(bookmarksFilePath)
		if err != nil {
			log.Fatalf(
				"failed to read bookmarks file for %s: %v",
				getProfileName(profileIndex),
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
