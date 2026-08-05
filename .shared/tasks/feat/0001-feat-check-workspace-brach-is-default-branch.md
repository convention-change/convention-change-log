---
id: 0001-feat-check-workspace-brach-is-default-branch
title: check workspace brach is default branch
status: in-progress
owner: "sinlovppt"
locked_at: ""
priority: medium
category: feat
created_at: "2026-08-05"
---

## feature description

- let generate workspace branch is default branch

### Suggested solution

after v1.14+ , must check branch for `main` branch
- add flat `--skip-branch` at global flag, to pass check branch
- add config `check branch with match list`
  - config setting at `.versionrc` key `branch-check`, will with array of string, if not set, will check `main` branch only
  - match `branch-check` setting use golang package `github.com/bmatcuk/doublestar/v4`
  - default config setting is `["main"]` if not set, will check `main` branch only, and will effect at sub command `init`
  - `--dry-run` mode will check `branch-check` setting, if not match, will print warning message, but not exit with error code

## do task

- update code at `cmd/kit/command/`
- this change will record at `doc/feature/99-feat-check-workspace-brach-is-default-branch.md` file
- must update usage guide at
  - `README.md`
  - `README.zh-Hans.md`

## check task

- check as this project setting
  - go unit test
  - go lint
  - go vet
- check new feature with `--dry-run` mode, and check warning message
- check new feature with `--skip-branch` mode, and check warning message
