# AGENTS.md

Repo-specific guidance for OpenCode agents working on `convention-change-log`,
a Go CLI that generates changelogs from Conventional Commits. Keep this file
short and high-signal; verify against the code if anything looks stale.

## Toolchain

- **Go 1.25** required (`go.mod`: `go 1.25.0`, toolchain `go1.25.11`). CI matrix
  tests `^1.25` and `1.25.11`.
- `package.json` at the repo root is **not a Node project**. It is Go-embedded
  metadata (`//go:embed package.json` in `resource.go`) consumed by the CLI for
  name/version/author. The release flow rewrites its `version` field. The real
  JS package is `.opencode/package.json` (OpenCode tooling only, CI-ignored).
- `golangci-lint` **v2** is the active linter. Config is `.golangci-v2.yaml`
  (used by both `make style` and CI). `.golangci.yaml` is the legacy v1 config
  — do not edit it expecting changes to take effect. v2 config has `tests: false`
  (test files are not linted), `modules-download-mode: readonly`, and
  `new-from-merge-base: main` (only new issues vs `main` are reported).

## Commands (Makefile is the source of truth)

```bash
make init dep      # first-time setup: check Go env, verify/download/tidy modules
make test          # fast unit tests
make style         # go mod verify + tidy + fmt + golangci-lint v2 (local style check)
make ci            # style + go vet + test + run.help  (local CI replica)
make ci.all        # ci + benchmark + coverage show
make buildMain     # build ./build/convention-change-log
make dev / dev.help  # build then run with CI_DEBUG=true / -h
make clean         # remove build/, logs/, coverage files
```

- To run a single Go test, match CI's invocation: `go test -tags test -run <Name> ./<pkg>/...`
- Benchmarks and coverage targets always pass `-tags test`; `make test` itself
  does not, but CI runs `go test -v -tags test ./...`. Prefer `-tags test`.
- Integration test coverage (binary-instrumented) is separate:
  `make test.go.integration.run` (uses `GOCOVERDIR`, not the unit-test profile).

## Build tags

`cmd/convention-change-log/main.go` and `example/TestMain.go` carry
`//go:build !test`. CI builds and tests with `-tags test`, so the main binary
source is excluded under the test tag. When running `go build`/`go test` by hand
to match CI, pass `-tags test`. `make` targets already handle this where needed.

## Tests and golden files

- Test stack: `stretchr/testify` + `sebdah/goldie/v2` (golden-file assertions).
  `goldie.New(t, goldie.WithDiffEngine(...))` is the common pattern; no custom
  fixture dir, so golden files land next to the test as `testdata/*.golden`.
- Golden/testdata files live under `<pkg>/testdata/` (36 tracked `.golden`
  files across `changelog/`, `convention/`, `example_test/`).
- **Critical gitignore trap:** `.gitignore` has `test[Dd]ata/` and `*.golden`
  patterns. Existing files are tracked only because git keeps tracking files
  after they were added. **New** `testdata/` or `*.golden` files are silently
  ignored — `git add -f <file>` is required, or they will not be committed and
  tests will fail for others. Always check `git status --ignored` after creating
  fixtures.
- To regenerate goldens, goldie respects its update flag; confirm the new output
  is actually tracked (see trap above) before committing.

## Package layout

- `cmd/convention-change-log/main.go` — CLI entry (build tag `!test`).
- `cmd/kit/cli/` — builds the `urfave/cli/v2` app (`NewCliApp`).
- `cmd/kit/command/` — command + flag logic: `global.go`, `flag.go`,
  `change_log_generator*.go`, and subcommand packages `subcommand_init/`,
  `subcommand_read_latest/`.
- `convention/` — **core domain**: commit parsing, changelog spec
  (`changeLogSpec.go`), template rendering, git repo info, types. Pure logic.
- `changelog/` — markdown node generation (`markdown.go`) and reader
  (`reader.go`); the output layer over `convention/`.
- `internal/` — private: `log/`, `pkg_kit/`, `tools/`, `urfave_cli/`.
- `constant/version.go` — version + copyright-year constants.
- `resource.go` — `//go:embed` of `package.json` and
  `resource/versionrc-beauty.json`.
- `z-MakefileUtils/` — included Makefile fragments; excluded from lint.

## Core behavior (version bump rules)

Default SemVer bump from commit types since the last tag:

- any `feat:` → **MINOR**
- otherwise → **PATCH**
- `BREAKING CHANGE:` or `BREAKING-CHANGE:` footer (or `!` after type) → **MAJOR**
- `-r`/`--release-as` overrides; no `feat` means PATCH.

This is the behavior the tool implements and also dogfoods for its own releases.

## Config and dogfooding

- `.versionrc` is the tool's own release config **and** the format it generates
  for users. `tag-prefix: "v"`, `monorepo-pkg-path: []`. Editing it changes this
  repo's own changelog generation.
- Conventional Commits are **required** for this repo's history (the repo
  consumes its own output). Commit format used here:
  `<type>[scope]: <emoji> <description>`. The `/gitcc` OpenCode command
  (`.opencode/commands/gitcc.md`) generates compliant messages.
- `.opencode/` and `.shared/` are **ignored by CI** (`ci.yml` `paths-ignore`).
  They are local tooling/config, not part of the build or test surface.

## Debug / env toggles

- `CLI_VERBOSE=true` — verbose CLI output (`make run.debug` / `make dev` set this).
- `CI_DEBUG=true` — debug mode for `make dev`.
- `CLI_DRY_RUN_DISABLE` — disable dry-run (the tool defaults to dry-run since v1.7).
- `CLI_LOG_LEVEL`, `CLI_CONFIG_TIMEOUT_SECOND`, `CLI_GIT_INFO_SCHEME`,
  `CLI_SKIP_WORKTREE_CHECK` — other runtime toggles.

## Release / CI flow

- Push a tag → `go-release-platform.yml` cross-builds
  (linux/darwin/windows × amd64/arm64, `CGO_ENABLED=0`, static) →
  `deploy-tag.yml` creates a GitHub Release with archive + `.sha256` assets.
- `version.yml` uses `convention-change/conventional-version-check@v1.5.0` to
  derive version/changelog for the release body.
- CI triggers on branches: `release-*`, `FE-*`, `*-feature-*`,
  `*-enhancement-*`, `BF-*`, `*-bug-*`, `PU-*`, `DOC-*`, `*-documentation-*`,
  `*-hotfix-*`, and on all tags. `main` builds but release artifacts only ship
  from tags.
