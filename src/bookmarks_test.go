package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestResolveBookmarksFilePrefersAccountBookmarks(t *testing.T) {
	profileDir := t.TempDir()
	accountBookmarks := filepath.Join(profileDir, "AccountBookmarks")
	legacyBookmarks := filepath.Join(profileDir, "Bookmarks")

	for _, path := range []string{accountBookmarks, legacyBookmarks} {
		if err := os.WriteFile(path, []byte("{}"), 0o600); err != nil {
			t.Fatalf("write test bookmark file: %v", err)
		}
	}

	got, err := resolveBookmarksFile(profileDir)
	if err != nil {
		t.Fatalf("resolveBookmarksFile() error = %v", err)
	}
	if got != accountBookmarks {
		t.Fatalf("resolveBookmarksFile() = %q, want %q", got, accountBookmarks)
	}
}

func TestResolveBookmarksFileFallsBackToBookmarks(t *testing.T) {
	profileDir := t.TempDir()
	legacyBookmarks := filepath.Join(profileDir, "Bookmarks")
	if err := os.WriteFile(legacyBookmarks, []byte("{}"), 0o600); err != nil {
		t.Fatalf("write test bookmark file: %v", err)
	}

	got, err := resolveBookmarksFile(profileDir)
	if err != nil {
		t.Fatalf("resolveBookmarksFile() error = %v", err)
	}
	if got != legacyBookmarks {
		t.Fatalf("resolveBookmarksFile() = %q, want %q", got, legacyBookmarks)
	}
}

func TestResolveBookmarksFileReturnsErrorWhenNeitherExists(t *testing.T) {
	if _, err := resolveBookmarksFile(t.TempDir()); err == nil {
		t.Fatal("resolveBookmarksFile() error = nil, want an error")
	}
}
