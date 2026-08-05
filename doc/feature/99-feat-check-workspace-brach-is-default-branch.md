# Feature: Check workspace branch is default branch

> Since v1.14.0

## Overview

This feature checks whether the current git workspace branch matches the expected default branch before running changelog generation or initialization. This prevents accidentally creating releases from feature branches.

## Configuration

### `.versionrc` key `branch-check`

Add a `branch-check` key to your `.versionrc` file with an array of glob patterns:

```json
{
  "branch-check": ["main"]
}
```

- **Default**: `["main"]` if not set
- **Matching**: Uses [doublestar v4](https://github.com/bmatcuk/doublestar) glob pattern matching
- Supports patterns like `main`, `release-*`, `{main,master}`, etc.

### Global flag `--skip-branch`

Skip the branch check entirely:

```bash
$ convention-change-log --skip-branch --dry-run
```

Environment variable: `CLI_SKIP_BRANCH_CHECK=true`

## Behavior

| Mode | Branch matches | Branch does NOT match |
|------|---------------|----------------------|
| `--dry-run` (default) | Proceeds normally | Prints warning, continues |
| `--dry-run-disable` | Proceeds normally | Exits with error |
| `--skip-branch` | Skips check | Skips check |

## Affected commands

- Main generate command (default action)
- `init` subcommand

## Examples

### Custom branch patterns

```json
{
  "branch-check": ["main", "release-*"]
}
```

### Skip branch check

```bash
$ convention-change-log --skip-branch
$ convention-change-log init --skip-branch
```

### Dry-run shows warning on wrong branch

```bash
$ git checkout feature/some-work
$ convention-change-log --dry-run
# WARNING: current branch "feature/some-work" does not match branch-check patterns [main], this may not be the intended release branch
```
