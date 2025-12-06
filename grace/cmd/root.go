package cmd

import (
	"flag"
	"fmt"
	"os"
	"github.com/your-org/grace/config"
	"github.com/your-org/grace/checker"
	"github.com/your-org/grace/report"
)

func Execute() {
       help := flag.Bool("help", false, "Show help")
       version := flag.Bool("version", false, "Show version")
       configPath := flag.String("config", "repo-structure.yaml", "Path to YAML config file")
       basePath := flag.String("base", ".", "Base path to check")
       flag.Parse()

       if *help {
	       fmt.Println("GRACE - Git Repository Architecture Compliance Enforcer")
	       fmt.Println("Usage: grace --config <path> [--base <dir>] [--help] [--version]")
	       fmt.Println("Options:")
	       fmt.Println("  --config   Path to YAML config file")
	       fmt.Println("  --base     Base path to check (default: .)")
	       fmt.Println("  --help     Show help")
	       fmt.Println("  --version  Show version")
	       os.Exit(0)
       }
       if *version {
	       fmt.Println("GRACE version 1.0.0")
	       os.Exit(0)
       }

       structure, ignore, err := config.LoadConfig(*configPath)
       if err != nil {
	       fmt.Println("Error loading config:", err)
	       os.Exit(1)
       }

       issues, err := checker.ValidateStructure(structure, *basePath, ignore.Patterns)
       if err != nil {
	       fmt.Println("Error validating structure:", err)
	       os.Exit(1)
       }

       report.PrintIssues(issues)
       if len(issues) > 0 {
	       os.Exit(1)
       }
       os.Exit(0)
}
