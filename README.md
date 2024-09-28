[![ci](https://github.com/convention-change/convention-change-log/actions/workflows/ci.yml/badge.svg)](https://github.com/convention-change/convention-change-log/actions/workflows/ci.yml)

[![go mod version](https://img.shields.io/github/go-mod/go-version/convention-change/convention-change-log?label=go.mod)](https://github.com/convention-change/convention-change-log)
[![GoDoc](https://godoc.org/github.com/convention-change/convention-change-log?status.png)](https://godoc.org/github.com/convention-change/convention-change-log)
[![goreportcard](https://goreportcard.com/badge/github.com/convention-change/convention-change-log)](https://goreportcard.com/report/github.com/convention-change/convention-change-log)

[![GitHub license](https://img.shields.io/github/license/convention-change/convention-change-log)](https://github.com/convention-change/convention-change-log)
[![GitHub latest SemVer tag)](https://img.shields.io/github/v/tag/convention-change/convention-change-log)](https://github.com/convention-change/convention-change-log/tags)
[![GitHub release)](https://img.shields.io/github/v/release/convention-change/convention-change-log)](https://github.com/convention-change/convention-change-log/releases)

[![en](https://img.shields.io/badge/lang-en-blue.svg)](https://github.com/convention-change/convention-change-log/blob/main/README.md)
[![zh](https://img.shields.io/badge/lang-%E4%B8%AD%E6%96%87-red)](https://github.com/convention-change/convention-change-log/blob/main/README.zh-Hans.md)

## Features

- [x] subcommand `init` to init config file
    - init `.versionrc` file at git repository root path, as default config file

```json
{
  "types": [
    {
      "type": "feat",
      "section": "✨ Features",
      "hidden": false
    },
    {
      "type": "fix",
      "section": "🐛 Bug Fixes",
      "hidden": false
    },
    {
      "type": "docs",
      "section": "📝 Documentation",
      "hidden": false
    },
    {
      "type": "style",
      "section": "💄 Styles",
      "hidden": true
    },
    {
      "type": "refactor",
      "section": "♻ Refactor",
      "hidden": false
    },
    {
      "type": "perf",
      "section": "⚡ Performance Improvements",
      "hidden": false
    },
    {
      "type": "test",
      "section": "✅ Tests",
      "hidden": true
    },
    {
      "type": "build",
      "section": "👷‍ Build System",
      "hidden": false
    },
    {
      "type": "ci",
      "section": "🔧 Continuous Integration",
      "hidden": true
    },
    {
      "type": "chore",
      "section": "📦 Chores",
      "hidden": true
    },
    {
      "type": "revert",
      "section": "⏪ Reverts",
      "hidden": false
    }
  ],
  "tag-prefix": "v"
}
```

- [x] can read git root `.versionrc` for setting of change log generate
    - support change log item sort by `versionrc` config `{{ .types[ .sort ] }}`, and default sort will auto set by this
      kit
    - more settings use `init --more` to generate `.versionrc` file
- [x] subcommand `read-latest` read the latest change log or write latest change to file
    - `--read-latest-file` read change log file path (default: "CHANGELOG.md")
    - `--read-latest-out` flag can open output to file, not settings will not output
        - `--read-latest-out-path` write last change file path (default: "CHANGELOG.txt")
- [x] global flag
    - [x] `--dry-run` flag can see what change of new release
    - [x] `-r` or `--release-as` to set release version
        - when not set will auto generate release version
            - commit message contains `feat:` will update `MINOR` version
            - commit message not contains `feat:` will update `MAJOR` version
    - [x] `--auto-push` flag can auto push tag to remote
    - [x] `--tag-prefix` flag can change tag prefix, default will use `.versionrc` config `tag-prefix`
- generate from [conventional commits](https://www.conventionalcommits.org) for [semver.org](https://semver.org/)
    - [x] default will update `PATCH` version
    - [x] if the latest list has any `feat` message will update `MINOR` version
    - [x] if you want change release version please use global flag `-r`
- auto update version resource
    - [x] project has `package.json` file, will auto update `version` field
    - [x] project has `package-lock.json` file, will try use `npm install` to update `package-lock.json` file
    - [x] in `.versionrc` has `monorepo-pkg-path` field as string list, will auto update `package.json` file
      in `monorepo-pkg-path` path (v1.5.+)

```json
{
  "monorepo-pkg-path": [
    "pkg1",
    "pkg2"
  ]
}
```

- [x] git url scheme default is `https` can change.(v1.8+)
    - use cli flag `--git-info-scheme` to change git info scheme, only support: https, http
    - in `.versionrc` has `cover-git-info-scheme` field as string, will change remote for example fill in `http`
- [x] check worktree is dirty (v1.8+)
    - add flag `--skip-worktree-check` will skip check (v1.8.1+)
    - check repository is dirty like `git status --porcelain`
    - if repository has submodule, will check, like `git submodule status --recursive`

more use see `convention-change-log --help`

## for what

- convention change log generate from [https://www.conventionalcommits.org/](https://www.conventionalcommits.org/)
- for [semver 2.0.0](https://semver.org)
- The convention commit log generation configuration file is customized under each project through `.versionrc`

## usage

### cli

[![GitHub release)](https://img.shields.io/github/v/release/convention-change/convention-change-log)](https://github.com/convention-change/convention-change-log/releases)

```bash
# install at $(GO_PATH)/bin
$ go install -v github.com/convention-change/convention-change-log/cmd/convention-change-log@latest
# install version v1.8.1
$ go install -v github.com/convention-change/convention-change-log/cmd/convention-change-log@v1.8.3
````

- or install by [release](https://github.com/convention-change/convention-change-log/releases) and add environment
  variables `PATH`

- please use this tool in `git repository root path`

```bash
## initialize configuration
# generate configuration in git repository root path
$ convention-change-log init
# files will be added `.versionrc`

## To generate the change log, this must run the repository root path and the project must be managed by git.
# it is recommended to be implemented every time --dry-run check to prevent errors
# with dry run
# after v1.7.0, it is turned on by default --dry-run
# If the flag `-r` is not set, the type `feat` will be determined according to whether the commit contains'
# if there is a log of type `feat`，then increment the MINOR version
# if there is no log of type `feat`, then increment the PATCH version
$ convention-change-log --dry-run
# --skip-worktree-check will skip check worktree (v1.8.1+)
$ convention-change-log --dry-run --skip-worktree-check

# flat -r to set custom release version
$ convention-change-log -r 0.1.0 --dry-run
# change tag prefix
$ convention-change-log -r 0.1.0 -t "" --dry-run

# disable --dry-run mode to generate logs and tags locally
$ convention-change-log --dry-run false
# or using env CLI_DRY_RUN_DISABLE=true can also be closed
$ convention-change-log --dry-run-disable
# you can also write like this
$ convention-change-log -r 1.0.0 --dry-run-disable

# more practical automatic push to remote, after v1.7.0, open --auto-push will ignore --dry-run
$ convention-change-log --auto-push

## read-latest can read the latest change log
# read and output to stdout
$ convention-change-log read-latest
# read and output to file `CHANGELOG.txt`
$ convention-change-log read-latest --read-latest-out
# read output to a custom file
$ convention-change-log read-latest --read-latest-out --read-latest-out-path CHANGELOG-1.txt
```

### best practices

- set `convention-change-log` to as alias `ccl`, bash or zsh as

```rc
alias gcmpt='git checkout $(git_main_branch) && git pull --verbose && git fetch --tags'
alias ccl='convention-change-log'
```

- powershell set in `$PROFILE` file

```ps1
## alias gcmpt 'git checkout $(git_main_branch) && git pull && git fetch --tags'
Function GitCheckOutMainAndFetchAllTagsFun {git checkout $(git config --get init.defaultBranch) ; git pull --verbose; git fetch --tags}
Set-Alias -Name gcmpt -Value GitCheckOutMainAndFetchAllTagsFun

## need install https://github.com/convention-change/convention-change-log
# alias ccl='convention-change-log'
Set-Alias -Name ccl -Value convention-change-log
```

- Please set up the `branch protection rule` and `tag protection rule` in the remote warehouse first to prevent false
  submission or delete tag
- Before generating CHANGELOG.md, it is best to merge into the main branch, the alias above `gcmpt` You can quickly
  switch to the main branch and synchronize to the latest information (including tag synchronization)
- Check the status of the current branch is available after the nail version
- execution `ccl --dry-run` or `ccl --dry-run -r 1.2.3` Specify the target version and verify that the expected change
  log is correct
- after confirming that there are no errors execute `ccl --auto-push`
- It is recommended to use the `release` branch to automatically merge and trigger CI automatic production to trigger
  the generation of CHANGELOG.md to avoid thinking that the operation is wrong.

# dev

## Contributing

[![Contributor Covenant](https://img.shields.io/badge/contributor%20covenant-v1.4-ff69b4.svg)](.github/CONTRIBUTING_DOC/CODE_OF_CONDUCT.md)
[![GitHub contributors](https://img.shields.io/github/contributors/convention-change/convention-change-log)](https://github.com/convention-change/convention-change-log/graphs/contributors)

We welcome community contributions to this project.

Please read [Contributor Guide](.github/CONTRIBUTING_DOC/CONTRIBUTING.md) for more information on how to get started.

请阅读有关 [贡献者指南](.github/CONTRIBUTING_DOC/zh-CN/CONTRIBUTING.md) 以获取更多如何入门的信息

developer see doc at [doc/dev.md](doc/dev.md)