[![ci](https://github.com/convention-change/convention-change-log/actions/workflows/ci.yml/badge.svg)](https://github.com/convention-change/convention-change-log/actions/workflows/ci.yml)

[![go mod version](https://img.shields.io/github/go-mod/go-version/convention-change/convention-change-log?label=go.mod)](https://github.com/convention-change/convention-change-log)
[![GoDoc](https://godoc.org/github.com/convention-change/convention-change-log?status.png)](https://godoc.org/github.com/convention-change/convention-change-log)
[![goreportcard](https://goreportcard.com/badge/github.com/convention-change/convention-change-log)](https://goreportcard.com/report/github.com/convention-change/convention-change-log)

[![GitHub license](https://img.shields.io/github/license/convention-change/convention-change-log)](https://github.com/convention-change/convention-change-log)
[![GitHub latest SemVer tag)](https://img.shields.io/github/v/tag/convention-change/convention-change-log)](https://github.com/convention-change/convention-change-log/tags)
[![GitHub release)](https://img.shields.io/github/v/release/convention-change/convention-change-log)](https://github.com/convention-change/convention-change-log/releases)

[![en](https://img.shields.io/badge/lang-en-blue.svg)](https://github.com/github.com/convention-change/convention-change-log/blob/main/README.md)

## for what

- 约定式提交日志生成 规则文档 [https://www.conventionalcommits.org/zh-hans](https://www.conventionalcommits.org/zh-hans)
- [语义化版本 2.0.0 zh-CN](https://semver.org/lang/zh-CN/)
- 约定式提交日志生成配置文件在每个工程下通过 `.versionrc` 自定义

## 用法

### cli

[![GitHub release)](https://img.shields.io/github/v/release/convention-change/convention-change-log)](https://github.com/convention-change/convention-change-log/releases)

```bash
# install at $(GO_PATH)/bin
$ go install -v github.com/convention-change/convention-change-log/cmd/convention-change-log@latest
# install version @v1.9.1
$ go install -v github.com/convention-change/convention-change-log/cmd/convention-change-log@v1.9.1
````

- 或在这里下载二进制包 [release](https://github.com/convention-change/convention-change-log/releases)
  请配置到环境变量 `PATH` 中生效

- 请在 `git存储库根路径` 中使用本工具

```bash
## 初始化 配置
# 在 git 存储库根路径生成配置
$ convention-change-log init
# 将添加文件 `.versionrc`
$ convention-change-log init --dry-init
# 只会显示 配置文件

## 生成更改日志，这必须运行存储库根路径和项目必须由git管理
# 建议每次都执行 --dry-run 检查以防止错误
# with dry run
# 在 v1.7.0 后 默认 都开启 --dry-run
# 不设置 -r 版本，会根据 commit 中是否含有 类型 `feat`
# 如果 含有 类型 `feat` 的日志，则在最后一个版本上 次版本号 +1
# 如果 没有 类型 `feat` 的日志，则在最后一个版本上 修订号 +1
$ convention-change-log --dry-run
# --skip-worktree-check 将跳过检查 worktree (v1.8.1+)
$ convention-change-log --dry-run --skip-worktree-check

# 设置 -r 自定义发布版本
$ convention-change-log -r 0.1.0 --dry-run
# 更改 tag 前缀
$ convention-change-log -r 0.1.0 -t "" --dry-run

# 关闭 --dry-run 模式 在 本地 生成 日志 和 tag
$ convention-change-log --dry-run false
# 或使用 env CLI_DRY_RUN_DISABLE=true 也能关闭
$ convention-change-log --dry-run-disable
# 也可以这么写
$ convention-change-log -r 1.0.0 --dry-run-disable

# 更实用的自动推送到远程，v1.7.0 后，开启 --auto-push 会忽略 --dry-run
$ convention-change-log --auto-push

## read-latest 可以读取最新的更改日志
# 读取并输出到 stdout
$ convention-change-log read-latest
# 读取并输出到文件 `CHANGELOG.txt`
$ convention-change-log read-latest --read-latest-out
# 读取输出到自定义 文件
$ convention-change-log read-latest --read-latest-out --read-latest-out-path CHANGELOG-1.txt
```

### 最佳实践

- 设置 `convention-change-log` 别名 `ccl` 便于日常操作, 在 `bash` or `zsh` 配置中添加别名

```rc
alias gcmpt='git checkout $(git_main_branch) && git pull --verbose && git fetch --tags'
alias ccl='convention-change-log'
```

- powershell 在文件 `$PROFILE` 中添加

```ps1
## alias gcmpt 'git checkout $(git_main_branch) && git pull && git fetch --tags'
Function GitCheckOutMainAndFetchAllTagsFun {git checkout $(git config --get init.defaultBranch) ; git pull --verbose; git fetch --tags}
Set-Alias -Name gcmpt -Value GitCheckOutMainAndFetchAllTagsFun

## need install https://github.com/convention-change/convention-change-log
# alias ccl='convention-change-log'
Set-Alias -Name ccl -Value convention-change-log
```

- 请先在远程仓库，设置好 `分支保护规则` 和 `tag 保护规则`，防止误提交，或者删除 tag
- 生成 CHANGELOG.md 前，最好合并到 主分支, 上面的别名 `gcmpt` 就可以快速切到主分支，并同步到最新信息（包括 tag 同步）
- 检查当前分支的状态是可以进行钉版本后
- 执行 `ccl --dry-run` 或者 `ccl --dry-run -r 1.2.3` 指定目标版本，确认 预期生成的 change log 是否是正确的
- 确认没有错误后，执行 `ccl --auto-push`
- 更建议通过 `release` 分支，自动合并触发 CI 自动生产的方式，来触发生成 CHANGELOG.md，避免认为操作错误
