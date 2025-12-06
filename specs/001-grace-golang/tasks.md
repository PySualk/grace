# Tasks: GRACE (Git Repository Architecture Compliance Enforcer)

## Phase 1: Setup
- [X] T001 Create `grace/` project directory in repository root
- [X] T002 Create initial `go.mod` in `grace/`
- [X] T003 Create `.dockerignore` and `Dockerfile` in `grace/`
- [X] T004 [P] Add README for GRACE in `grace/README.md`

## Phase 2: Foundational
- [X] T005 Implement CLI entrypoint in `grace/cmd/root.go`
- [X] T006 Implement YAML config loader in `grace/config/config.go`
- [X] T007 Implement structure validation logic in `grace/checker/checker.go`
- [X] T008 Implement output formatting in `grace/report/report.go`
- [X] T009 [P] Add unit test scaffolding in `tests/unit/`
- [X] T010 [P] Add integration test scaffolding in `tests/integration/`

## Phase 3: User Story 1 (Configurable Structure Enforcement)
- [X] T011 [US1] Parse YAML config and build expected structure tree in `grace/config/config.go`
- [X] T012 [US1] Traverse repo and compare to expected structure in `grace/checker/checker.go`
- [X] T013 [US1] Report missing/extra files and directories in `grace/report/report.go`
- [X] T014 [P] [US1] Add unit tests for config parsing in `tests/unit/config_test.go`
- [X] T015 [P] [US1] Add unit tests for structure validation in `tests/unit/checker_test.go`
- [X] T016 [P] [US1] Add integration test for CLI in `tests/integration/cli_test.go`

## Phase 4: User Story 2 (Ignore Patterns)
- [X] T017 [US2] Implement ignore pattern support in `grace/config/config.go`
- [X] T018 [US2] Integrate ignore logic in structure validation in `grace/checker/checker.go`
- [X] T019 [P] [US2] Add unit tests for ignore patterns in `tests/unit/config_test.go`

## Phase 6: Polish & Cross-Cutting
- [X] T023 Add CLI help, version, and usage docs in `grace/cmd/root.go`
- [X] T024 Add code linting and formatting config in `grace/`
- [X] T025 Add CI workflow for build/test in `.github/workflows/grace.yml`
- [X] T026 Add container build/test instructions in `grace/README.md`

## Dependencies
- Phase 1 and 2 must complete before any user story phases
- US1 (Configurable Structure Enforcement) is required for US2
- US2 (Ignore Patterns) can be developed after US1
- Polish phase can run in parallel with final user story tasks

## Parallel Execution Examples
- T004, T009, T010 can run in parallel during setup
- T014, T015, T016 can run in parallel after US1 implementation
- T019 can run in parallel after US2 implementation
- T023, T024, T025, T026 can run in parallel during polish

## Implementation Strategy
- MVP: Complete Phase 1, 2, and US1 (T001-T016)
- Incremental delivery: Add ignore patterns (US2), then polish

## Independent Test Criteria
- US1: Given a valid YAML config, GRACE reports all missing/extra files and directories
- US2: Given ignore patterns, GRACE excludes specified files/dirs from validation

## Task Format Validation
All tasks use strict checklist format: `- [ ] T### [P] [US#] Description with file path`
