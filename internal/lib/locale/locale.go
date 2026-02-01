package locale

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

var (
	// messages stores all loaded locale messages
	messages = make(map[string]map[string]string)
	mu       sync.RWMutex

	// Default language
	defaultLang = "en"
)

// Load loads all locale files from the locales directory
func Load() error {
	localesDir := "locales"

	// Read all JSON files in locales directory
	files, err := filepath.Glob(filepath.Join(localesDir, "*.json"))
	if err != nil {
		return fmt.Errorf("failed to read locales directory: %w", err)
	}

	if len(files) == 0 {
		return fmt.Errorf("no locale files found in %s", localesDir)
	}

	mu.Lock()
	defer mu.Unlock()

	for _, file := range files {
		// Extract language code from filename (e.g., "en" from "en.json")
		lang := filepath.Base(file)
		lang = lang[:len(lang)-5] // Remove .json extension

		// Read file
		data, err := os.ReadFile(file)
		if err != nil {
			return fmt.Errorf("failed to read locale file %s: %w", file, err)
		}

		// Parse JSON
		var localeMessages map[string]string
		if err := json.Unmarshal(data, &localeMessages); err != nil {
			return fmt.Errorf("failed to parse locale file %s: %w", file, err)
		}

		messages[lang] = localeMessages
	}

	return nil
}

// Get retrieves a message for a given key and language
func Get(lang, key string) string {
	mu.RLock()
	defer mu.RUnlock()

	// Try to get message for requested language
	if langMessages, ok := messages[lang]; ok {
		if message, ok := langMessages[key]; ok {
			return message
		}
	}

	// Fallback to default language
	if lang != defaultLang {
		if langMessages, ok := messages[defaultLang]; ok {
			if message, ok := langMessages[key]; ok {
				return message
			}
		}
	}

	// If still not found, return the key itself
	return key
}

// SetDefaultLanguage sets the default fallback language
func SetDefaultLanguage(lang string) {
	mu.Lock()
	defer mu.Unlock()
	defaultLang = lang
}

// GetAvailableLanguages returns list of loaded languages
func GetAvailableLanguages() []string {
	mu.RLock()
	defer mu.RUnlock()

	langs := make([]string, 0, len(messages))
	for lang := range messages {
		langs = append(langs, lang)
	}
	return langs
}
