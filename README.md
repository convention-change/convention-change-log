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

## Features

- [ ] more perfect test case coverage
- [ ] more perfect benchmark case

## usage

### cli

```bash
# install at $(GO_PATH)/bin
$ go install -v github.com/convention-change/convention-change-log/cmd/convention-change-log@latest
# install version v1.0.0
$ go install -v github.com/convention-change/convention-change-log/cmd/convention-change-log@v1.0.0
````

- please use cli at `git repository root path`

```bash
# init config file at git repository root path
$ convention-change-log init

# 
```

# dev

## env

- minimum go version: go 1.18
- change `go 1.18`, `^1.18`, `1.18.10` to new go version

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
