[![ci](https://github.com/convention-change/convention-change-log/actions/workflows/ci.yml/badge.svg)](https://github.com/convention-change/convention-change-log/actions/workflows/ci.yml)
[![license](https://img.shields.io/github/license/convention-change/convention-change-log)](https://github.com/convention-change/convention-change-log)
[![go mod version](https://img.shields.io/github/go-mod/go-version/convention-change/convention-change-log?label=go.mod)](https://github.com/convention-change/convention-change-log)
[![GoDoc](https://godoc.org/github.com/convention-change/convention-change-log?status.png)](https://godoc.org/github.com/convention-change/convention-change-log/)
[![GoReportCard](https://goreportcard.com/badge/github.com/convention-change/convention-change-log)](https://goreportcard.com/report/github.com/convention-change/convention-change-log)
[![codecov](https://codecov.io/gh/convention-change/convention-change-log/branch/main/graph/badge.svg)](https://codecov.io/gh/convention-change/convention-change-log)
[![GitHub release)](https://img.shields.io/github/v/release/convention-change/convention-change-log)](https://github.com/convention-change/convention-change-log/releases)

## for what

- convention change log generate
- convention change log reader
- convention change log config by `.versionrc` file

## Features

- [x] subcommand `init` to init config file
  - init `.versionrc` file at git repository root path, as default config file
```json
{
  "types": [
    {"type": "feat", "section": "✨ Features", "hidden": false},
    {"type": "fix", "section": "🐛 Bug Fixes", "hidden": false},
    {"type": "docs", "section":"📝 Documentation", "hidden": true},
    {"type": "style", "section":"💄 Styles", "hidden": true},
    {"type": "refactor", "section":"♻ Refactor", "hidden": false},
    {"type": "perf", "section":"⚡ Performance Improvements", "hidden": false},
    {"type": "test", "section":"✅ Tests", "hidden": true},
    {"type": "build", "section":"👷‍ Build System", "hidden": false},
    {"type": "ci", "section":"🔧 Continuous Integration", "hidden": true},
    {"type": "chore", "section":"📦 Chores", "hidden": true},
    {"type": "revert", "section":"⏪ Reverts", "hidden": false}
  ],
  "tag-prefix": "v"
}
```

- [x] can read git root `.versionrc` for setting of change log generate
    - support change log item sort by `versionrc` config `{{ .types[ .sort ] }}`, and default sort will auto set by this kit
    - more settings see `init --more`
- [x] subcommand `read-latest` read the latest change log or write latest change to file
- [x] global flag
    - [x] `-r` or `--release-as` to set release version
      - when not set will auto generate release version
        - commit message contains `feat:` will update `MINOR` version
        - commit message not contains `feat:` will update `MAJOR` version
    - [x] `--dry-run` flag can see what change of new release
    - [x] `--auto-push` flag can auto push tag to remote
    - [x] `--tag-prefix` flag can change tag prefix, default will use `.versionrc` config `tag-prefix`
- generate from [conventional commits](https://www.conventionalcommits.org) for [semver.org](https://semver.org/)
  - [x] default will update `PATCH` version
  - [x] if the latest list has any `feat` message will update `MINOR` version
  - [x] if you want change release version please use global flag `-r`
- auto update version resource
  - [x] project has `package.json` file, will auto update `version` field
  - [x] project has `package-lock.json` file, will try use `npm install` to update `package-lock.json` file
  - [x] in `.versionrc` has `monorepo-pkg-path` field as string list, will auto update `package.json` file in `monorepo-pkg-path` path (v1.5.+)

```json
{
  "monorepo-pkg-path": [
    "pkg1",
    "pkg2"
  ]
}
```

- [ ] more perfect test case coverage

more use see `convention-change-log --help`

## usage

### cli

```bash
# install at $(GO_PATH)/bin
$ go install -v github.com/convention-change/convention-change-log/cmd/convention-change-log@latest
# install version v1.5.0
$ go install -v github.com/convention-change/convention-change-log/cmd/convention-change-log@v1.5.0
````

- or install by [release](https://github.com/convention-change/convention-change-log/releases) and add environment variables `PATH`

- please use cli at `git repository root path`

```bash
# init config file at git repository root path
$ convention-change-log init

# check release note by tag
$ convention-change-log --dry-run
# let release version as -r
$ convention-change-log -r 0.1.0 --dry-run
# change tag prefix
$ convention-change-log -r 0.1.0 -t "" --dry-run

# finish check then generate release note and tag
$ convention-change-log -r 1.0.0

# or add auto push to remote
$ convention-change-log --auto-push
```

# dev

## env

- minimum go version: go 1.19
- change `go 1.19`, `^1.19`, `1.19.13` to new go version

### libs

| lib                                 | version |
|:------------------------------------|:--------|
| https://github.com/stretchr/testify | v1.8.4  |
| https://github.com/sebdah/goldie    | v2.5.3  |

- more libs see [go.mod](https://github.com/convention-change/convention-change-log/blob/main/go.mod)

## Contributing

[![Contributor Covenant](https://img.shields.io/badge/contributor%20covenant-v1.4-ff69b4.svg)](.github/CONTRIBUTING_DOC/CODE_OF_CONDUCT.md)
[![GitHub contributors](https://img.shields.io/github/contributors/convention-change/convention-change-log)](https://github.com/convention-change/convention-change-log/graphs/contributors)

We welcome community contributions to this project.

Please read [Contributor Guide](.github/CONTRIBUTING_DOC/CONTRIBUTING.md) for more information on how to get started.

请阅读有关 [贡献者指南](.github/CONTRIBUTING_DOC/zh-CN/CONTRIBUTING.md) 以获取更多如何入门的信息

## local dev

```bash
# It needs to be executed after the first use or update of dependencies.
$ make init dep
```

- test code

```bash
$ make test testBenchmark
```

add main.go file and run

```bash
# run at env dev use cmd/main.go
$ make dev
```

- ci to fast check

```bash
# check style at local
$ make style

# run ci at local
$ make ci
```

## docker

```bash
# then test build as test/Dockerfile
$ make dockerTestRestartLatest
# clean test build
$ make dockerTestPruneLatest

# more info see
$ make helpDocker
```
