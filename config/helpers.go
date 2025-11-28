package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// FileExists checks if a file exists and is not a directory before we
// try using it to prevent further errors.
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// DirectoryExists checks if a directory exists.
func DirectoryExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

// GetAbsolutePath resolves the given path to an absolute path.
// Returns an error if the path cannot be resolved.
func GetAbsolutePath(path string) (string, error) {
	if filepath.IsAbs(path) {
		return path, nil
	}

	absPath, err := filepath.Abs(path)
	if err != nil {
		return "", fmt.Errorf("failed to get absolute path for %s: %w", path, err)
	}

	return absPath, nil
}

// SanitizeString removes potentially harmful characters from a string.
func SanitizeString(input string) string {
	// Example: Remove characters that could be used in shell injection.
	replacer := strings.NewReplacer(
		";", "",
		"&", "",
		"|", "",
		">", "",
		"<", "",
		"`", "",
		"$", "",
		"(", "",
		")", "",
		"\\", "",
	)
	return replacer.Replace(input)
}

// ReadEnvVarOrDefault reads an environment variable and returns its value.
// If the environment variable is not set, it returns the provided default value.
func ReadEnvVarOrDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}