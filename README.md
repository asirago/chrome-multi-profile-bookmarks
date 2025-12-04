# Chrome Multi-Profile Bookmarks

An [Alfred](https://www.alfredapp.com/) workflow for searching bookmarks across multiple Chrome profiles with intelligent fuzzy matching and bulk opening functionality.

## Why

Managing bookmarks across multiple Chrome profiles is cumbersome. Chrome's built-in bookmark search only
works within the active profile, forcing you to switch profiles to access different bookmark collections.
This workflow solves that by letting you search all your profiles at once from Alfred with natural, fast 
search that just works.

## Features

- **Multi-Profile Search** - Search bookmarks across all your Chrome profiles simultaneously
- **Smart Ranking** - Intelligent scoring algorithm that prioritizes:
  - Exact matches over partial matches
  - Name matches over folder and URL matches
  - Prefix matches over substring matches
- **Profile Icons** - Shows each profile's icon for easy identification
- **Folder Context** - Displays the bookmark's folder location
- **Bulk Open** - Hold ⌘ to open all bookmarks in a folder simultaneously


## Installation

1. [Download the latest release](https://github.com/asirago/chrome-multi-profile-bookmarks/releases)
2. Double-click the `.alfredworkflow` file
3. Configure your profiles in Alfred's workflow settings

## Usage

### Basic Search

Type your keyword followed by search terms:

```
bm github
```
The workflow searches across:
- Bookmark names
- Folder names
- URLs

Selecting a bookmark will open the URL in your respective profile.

**Bulk Open:** Hold <kbd>⌘</kbd> before selecting a bookmark to open all URLs in that folder simultaneously in your respective Chrome profile. This is useful for:
- Opening all your daily work tabs at once
- Launching a set of project-related links
- Quickly accessing all bookmarks in a specific category

## Configuration

### Set Profiles to Search From

Configure which Chrome profiles to search in the workflow settings:

**Default:** `0,1` (searches Default profile and Profile 1)

Enter a comma-separated list of profile numbers (e.g., `0,1,2,3`):
- `0` = Default profile
- `1` = Profile 1 (second profile)
- `2` = Profile 2 (third profile)
- And so on...

Leave as `0,1` to search bookmarks from the first two profiles.

### Chrome Profile Locations

The workflow follows the Chromium standard and expects bookmarks at:
```
~/Library/Application Support/Google/Chrome/Default/Bookmarks
~/Library/Application Support/Google/Chrome/Profile {N}/Bookmarks
```

Where `{N}` is the profile number (1, 2, 3, etc.)

## Scoring Algorithm

The workflow uses a weighted scoring system:

| Match Type | Name | Folder | URL |
|-----------|------|--------|-----|
| Exact match | 30 | 25 | - |
| Prefix match | 15 | 12 | 10 |
| Word start | 8 | 6 | 4 |
| Substring | 3 | 2 | 1 |

## Building from Source

### Requirements

- Go 1.22+
- macOS (for Alfred)

### Build

**Universal binary (Intel + Apple Silicon):**
```bash
# Build for both architectures
GOARCH=amd64 go build -o cmpbs-amd64
GOARCH=arm64 go build -o cmpbs-arm64

# Combine into universal binary
lipo -create -output cmpbs cmpbs-amd64 cmpbs-arm64

# Clean up
rm cmpbs-amd64 cmpbs-arm64
```
or 

```
make release
```

**Single architecture:**

`go build -o cmpbs` or `make build`

### Development

```bash
# Run with flags
./cmpbs --profiles "0,1" --folder "Work" search terms

# Get folder URLs (pipe-separated)
./cmpbs --folder "Work"
```

## Flags

| Flag | Description | Default |
|------|-------------|---------|
| `--profiles` | Comma-separated profile indices | `"0,1"` |
| `--folder` | Filter by folder name | `""` |


## Contributing

Issues and pull requests are welcome!

https://github.com/asirago/chrome-multi-profile-bookmarks/issues

## License

MIT
