package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

var bookmarksDir string
var maxResults = 20

func init() {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	bookmarksDir = filepath.Join(home, "Library/Application Support/Google/Chrome")
}

func main() {

	var profilesCSV string
	var folderFlag string
	var maxResultFlag int

	flag.StringVar(&profilesCSV, "profiles", "0,1", "comma-separated list of profiles")
	flag.StringVar(&folderFlag, "folder", "", "filter bookmarks by what folder they are in")
	flag.IntVar(&maxResultFlag, "maxResult", 20, "max amount of items")

	flag.Parse()
	args := flag.Args()

	if maxResultFlag != 20 {
		maxResults = maxResultFlag
	}

	profilesIndex, err := csvToIntArray(profilesCSV)
	if err != nil {
		log.Fatal(err)
	}

	urls := getURLs(profilesIndex)

	if folderFlag != "" {
		stringURLs := []string{}
		for _, bookmark := range urls[folderFlag] {
			stringURLs = append(stringURLs, bookmark.URL)
		}

		fmt.Printf("%s", strings.Join(stringURLs, "|"))
		return
	}

	items := []*Item{}

	if len(args) > 0 {
		scoredItems := []scoredItem{}

		for folderHash, bookmarks := range urls {
			folder := strings.Split(folderHash, "-")[0]

			for _, bookmark := range bookmarks {
				score := calculateScore(folder, bookmark.Name, bookmark.URL, args)
				if score <= 0 {
					continue
				}

				profileName := getProfileName(bookmark.ProfileIndex)
				item := createItem(bookmark, bookmarksDir, profileName)

				scoredItems = append(scoredItems, scoredItem{
					Item:  item,
					Score: score,
				})
			}
		}

		sort.Slice(scoredItems, func(i, j int) bool {
			return scoredItems[i].Score > scoredItems[j].Score
		})

		maxItems := min(len(scoredItems), maxResults)
		for i := range maxItems {
			items = append(items, scoredItems[i].Item)
		}
	}

	body, err := json.Marshal(map[string]any{
		"items": items,
	})
	if err != nil {
		log.Fatalf("failed to encode json: %v", err)
	}

	fmt.Print(string(body))
}
