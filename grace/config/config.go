package config

import (
       "gopkg.in/yaml.v3"
       "os"
)

type RepoStructure map[string][]string

type IgnoreConfig struct {
       Patterns []string `yaml:"ignore"`
}

func LoadConfig(path string) (RepoStructure, *IgnoreConfig, error) {
       var structure RepoStructure
       var ignore IgnoreConfig
       file, err := os.ReadFile(path)
       if err != nil {
              return nil, nil, err
       }
       if err := yaml.Unmarshal(file, &structure); err != nil {
              return nil, nil, err
       }
       if err := yaml.Unmarshal(file, &ignore); err != nil {
              // ignore section is optional
              ignore.Patterns = nil
       }
       return structure, &ignore, nil
}
