package main

import (
	"fmt"
	"path/filepath"
)

type bookmarksFile struct {
	Checksum string                   `json:"checksum"`
	Roots    map[string]*bookmarkNode `json:"roots"`
}

type bookmarkNode struct {
	Children  []*bookmarkNode `json:"children"`
	DateAdded string          `json:"date_added"`
	Name      string          `json:"name"`
	Type      string          `json:"type"`
	URL       string          `json:"url"`
}

type bookmarkURL struct {
	ProfileIndex int
	Folder       string
	FolderHash   string
	Name         string
	URL          string
}

type Item struct {
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Arg      string `json:"arg"`
	Icon     struct {
		Path string `json:"path"`
	} `json:"icon"`
}

type scoredItem struct {
	Item  *Item
	Score int
}

func createItem(bookmark *bookmarkURL, bookmarksDir, profileName string) *Item {
	item := &Item{
		Title:    bookmark.Name,
		Subtitle: fmt.Sprintf("[%s] %s", bookmark.Folder, bookmark.URL),
		Arg:      fmt.Sprintf("%s|%s|%s", bookmark.URL, bookmark.FolderHash, profileName),
	}
	item.Icon.Path = filepath.Join(
		bookmarksDir,
		profileName,
		"Google Profile Picture.png",
	)
	return item
}
