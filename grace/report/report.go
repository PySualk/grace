package report

import "fmt"

func PrintIssues(issues []string) {
       if len(issues) == 0 {
	       fmt.Println("Repository structure is valid.")
       } else {
	       fmt.Println("Repository structure issues found:")
	       for _, issue := range issues {
		       fmt.Println("-", issue)
	       }
       }
}
