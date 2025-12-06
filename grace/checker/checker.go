package checker

import (
	"os"
	"path/filepath"
)

func ValidateStructure(structure map[string][]string, basePath string, ignorePatterns []string) ([]string, error) {
       var issues []string
       shouldIgnore := func(path string) bool {
	       for _, pattern := range ignorePatterns {
		       match, _ := filepath.Match(pattern, filepath.Base(path))
		       if match {
			       return true
		       }
	       }
	       return false
       }
       for key, items := range structure {
	       dirPath := basePath
	       if key != "root" {
		       dirPath = basePath + "/" + key
	       }
	       for _, item := range items {
		       itemPath := dirPath + "/" + item
		       if shouldIgnore(itemPath) {
			       continue
		       }
		       if len(item) > 0 && item[len(item)-1] == '/' {
			       if stat, err := os.Stat(itemPath); err != nil || !stat.IsDir() {
				       issues = append(issues, "Missing directory: "+itemPath)
			       }
		       } else {
			       if stat, err := os.Stat(itemPath); err != nil || stat.IsDir() {
				       issues = append(issues, "Missing file: "+itemPath)
			       }
		       }
	       }
       }
       return issues, nil
}
