# Chrome Multi-Profile Bookmarks for Alfred

An Alfred workflow for searching bookmarks across multiple Chrome profiles with intelligent fuzzy matching and bulk opening functionality.

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

## How to Use

Simply open Alfred and start typing your bookmark keyword followed by search terms:
```
bm github
bm react docs
bm work project
```

The workflow searches across bookmark names, folder names, and URLs.

**Selecting a bookmark** opens the URL in the correct Chrome profile.

**Bulk Open:** Hold <kbd>⌘</kbd> before selecting to open all URLs in that folder at once. Perfect for:
- Opening your daily work tabs
- Launching project-related links
- Accessing bookmark categories quickly

## Configuration

### Set Profiles to Search From

In the workflow configuration, set which Chrome profiles to search:

**Profiles to search:** `0,1`

Comma-separated list of profile numbers to search (e.g., `0,1,2`). Use `0` for the Default profile, `1` for the second profile, etc. Leave as `0,1` to search bookmarks of the first two profiles.

### Chrome Profile Locations

Bookmarks are expected at the standard Chromium locations:
```
~/Library/Application Support/Google/Chrome/Default/Bookmarks
~/Library/Application Support/Google/Chrome/Profile {N}/Bookmarks
```

## Support

If you need help or want to report a bug, please open an issue:

https://github.com/asirago/chrome-multi-profile-bookmarks/issues

## License

MIT

---

Made with ☕ by [asirago](https://github.com/asirago)
