package unit

import (
	"testing"
	"github.com/your-org/grace/checker"
)

func TestValidateStructure(t *testing.T) {
       structure := map[string][]string{
	       "root": {"README.md", "src/", "tests/"},
       }
       issues, err := checker.ValidateStructure(structure, ".")
       if err != nil {
	       t.Fatalf("Validation failed: %v", err)
       }
       // This test expects files/dirs to be missing in a clean test env
       if len(issues) == 0 {
	       t.Errorf("Expected missing items, got none")
       }
}
