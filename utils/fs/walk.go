package fs

import (
	"os"
	"path/filepath"
)

// walk dir, return files result, if function return true, append path to result
func Walk(dir string, function func(string, os.FileInfo) bool) []string {
	var result []string
	filepath.Walk("view", func(path string, info os.FileInfo, err error) error {
		if function(path, info) {
			result = append(result, path)
		}
		return nil
	})
	return result
}

// call Walk, append to result if  file name extension is ext
func WalkExt(dir, ext string) []string {
	return Walk(dir, func(path string, info os.FileInfo) bool {
		return filepath.Ext(path) == "."+ext
	})
}
