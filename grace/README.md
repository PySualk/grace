# GRACE (Git Repository Architecture Compliance Enforcer)

GRACE is a CLI tool to enforce a configurable directory and file structure in git repositories. It reads a YAML config, validates the repo structure, and outputs compliance results. Designed for local, CI, and containerized use.

## Features
- Configurable structure via YAML
- Ignore patterns support
- Clear CLI output
- Container-ready

## Usage
1. Place your structure config in `repo-structure.yaml`.
2. Run GRACE to validate your repository.

## Build & Run
```sh
go build -o grace ./main.go
./grace --config repo-structure.yaml
```

## Container
```sh
docker build -t grace .
docker run --rm -v $(pwd):/repo grace --config /repo/repo-structure.yaml
```

## CI
GRACE is tested and linted automatically via GitHub Actions. See `.github/workflows/grace.yml`.
