func TestIgnorePatterns(t *testing.T) {
       ignore := []string{"*.md", "tests/"}
       shouldIgnore := func(path string) bool {
	       for _, pattern := range ignore {
		       match, _ := filepath.Match(pattern, filepath.Base(path))
		       if match {
			       return true
		       }
	       }
	       return false
       }
       if !shouldIgnore("README.md") {
	       t.Errorf("README.md should be ignored")
       }
       if !shouldIgnore("tests/") {
	       t.Errorf("tests/ should be ignored")
       }
       if shouldIgnore("src/main.go") {
	       t.Errorf("src/main.go should not be ignored")
       }
}
package unit

import (
	"testing"
	"github.com/PySualk/grace/config"
)

func TestLoadConfig(t *testing.T) {
	sample := []byte(`root:
  - README.md
  - src/
  - tests/
`)
	var structure map[string][]string
	err := yaml.Unmarshal(sample, &structure)
	if err != nil {
		t.Fatalf("Failed to parse YAML: %v", err)
	}
	if len(structure["root"]) != 3 {
		t.Errorf("Expected 3 items in root, got %d", len(structure["root"]))
	}
}
