package util

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// NormalizePath removes extraneous spaces, new lines, and carriage returns from the given path.
//   - If the path starts with a tilde (~) character,
//     it replaces the tilde with the user's home directory.
//   - It then cleans the path to remove extra separators and dots.
//   - Finally, it converts the path to its absolute form if possible.
//   - If not, it returns the cleaned path.
func NormalizePath(path string) string {
	// remove extraneous spaces
	path = strings.TrimSpace(path)            // extra space at start and end
	path = strings.ReplaceAll(path, "\n", "") // new lines
	path = strings.ReplaceAll(path, "\r", "") // carriage returns

	// replace instances of the tilde user home directory shorthand
	if strings.HasPrefix(path, "~") {
		home, err := os.UserHomeDir()
		if err == nil {
			path = strings.Replace(path, "~", home, 1)
		}
	}

	// clean the path to remove extra separators and dots
	cleaned := filepath.Clean(path)

	// convert the path to its absolute form
	if normalized, err := filepath.Abs(cleaned); err == nil {
		return normalized
	}

	// use the cleaned path if we can't get an absolute
	return cleaned
}

func ProjectRootDir() string {
	_, caller, _, _ := runtime.Caller(0)
	return NormalizePath(filepath.Join(filepath.Dir(caller), ".."))
}

func ProjectDir(elem ...string) string {
	return NormalizePath(filepath.Join(ProjectRootDir(), filepath.Join(elem...)))
}
