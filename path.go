package autoindex

import (
	"path/filepath"
	"strings"
)

const indexHtml = "index.html"

// Index joins the given directories and "index.html"
//
// Returns "index.html" only if no directory is given.
func Index(dirs ...string) string {
	dir := filepath.Join(dirs...)
	if dir == "" {
		return indexHtml
	}

	return filepath.Join(dir, indexHtml)
}

// rel returns the relative path of a directory by removing the root directory.
//
// It's useful if you want to use a specific directory as the index root.
//
// We don't use [filepath.Rel] because it clears the trailing slash.
func rel(dir string) string {
	return strings.TrimPrefix(dir, Opts.Root)
}
