# Changelog

All notable changes to this project will be documented in this file. See [convention-change-log](https://github.com/convention-change/convention-change-log) for commit guidelines.

## [1.0.1](https://github.com/convention-change/convention-change-log/compare/1.0.0...v1.0.1) (2023-07-30)

### 🐛 Bug Fixes

* fix go-release-platform.yml config of build (2023-07-30) [9be3443d](https://github.com/convention-change/convention-change-log/commit/9be3443d6fb921e069cb9c6e9d721a20aacd9b1f)

## 1.0.0 (2023-07-30)

### BREAKING CHANGE:

* change convention.GitRepositoryInfo to convention.GitRepositoryHttpInfo

* move package from github.com/sinlov-go/convention-change-log to

### 🐛 Bug Fixes

* fix -r set version error when not find CHANGELOG.md (2023-07-30) [5ce6ec42](https://github.com/convention-change/convention-change-log/commit/5ce6ec42730ba64c9ea04fddeb3b5f5d088f6671)

* changelog.NewReader support HistoryFirst* method to get info of Changelog (2023-07-19) [c060add4](https://github.com/convention-change/convention-change-log/commit/c060add4aa775e69809121ff7fd7a7353401bb60)

### ✨ Features

* change default init config for command init (2023-07-30) [9f147f57](https://github.com/convention-change/convention-change-log/commit/9f147f57b5c1376bab4a1d8a16d30d0d5a479f40)

* let --dry-run only show changelog content (2023-07-30) [27cc198b](https://github.com/convention-change/convention-change-log/commit/27cc198bdb2c03da011b59f49f5bb7bd2e470244)

* support auto update package.json update by version (2023-07-30) [c6693b7b](https://github.com/convention-change/convention-change-log/commit/c6693b7b860ffb98aa31123713312ae58696626b)

* add default version (2023-07-28) [4bcd5c5f](https://github.com/convention-change/convention-change-log/commit/4bcd5c5f95b02ee137b1830b4584115edd213004)

* let tools can auto generate by https [7b29e1d3](https://github.com/convention-change/convention-change-log/commit/7b29e1d38f351bd7a9baa63b5beac587837c7623)

* let local git auto add and git (2023-07-22) [4c4c9bb3](https://github.com/convention-change/convention-change-log/commit/4c4c9bb30bc2de04f85af22438c4d7f6842cf16d)

* let cli global can find last tag and append history changelog (2023-07-21) [daf8c953](https://github.com/convention-change/convention-change-log/commit/daf8c9537d0aa8056f007e81b34248adcd50ed18)

* add ConventionalChangeLogSpec.CoverHttpHost (2023-07-21) [c9c98d7b](https://github.com/convention-change/convention-change-log/commit/c9c98d7b984ba00ef21a43400d58be2f829254a0)

* change global cli git clone by sinlov-go/git (2023-07-21) [103f2886](https://github.com/convention-change/convention-change-log/commit/103f2886875a2ad740a819e2b9b38dc533fbde4c)

* support patch changelog read-latest and let --read-latest-out not contain title (2023-07-21) [e4c7337f](https://github.com/convention-change/convention-change-log/commit/e4c7337f82e7f229bfb2c5df4c4d64895502ed5d)

* generateMarkdownNodes support markdown Issure link and BREAKING CHANGE (2023-07-19) [76e180ac](https://github.com/convention-change/convention-change-log/commit/76e180acd639296fc6ad5f834b9a1622fa8e07fb)

* add GenerateMarkdownNodes append breaking change (2023-07-19) [fcbdfbcf](https://github.com/convention-change/convention-change-log/commit/fcbdfbcfb1feb709ec7607cdde31c34ba3718d21)

* subcommand read-latest can read ChangeLog and write to out text (2023-07-19) [30769a5b](https://github.com/convention-change/convention-change-log/commit/30769a5b4a2df2d8ff27525ba213b8e78898bf6f)

* changelog.NewReader can read ChangeLog by path and get markdown node (2023-07-19) [83f55a74](https://github.com/convention-change/convention-change-log/commit/83f55a7454cbdc8ee4126927d9ffbc29bb17cf60)

* add exit_cli, change  NewCommitWithLogSpec func (2023-07-17) [b99bb424](https://github.com/convention-change/convention-change-log/commit/b99bb424e039d3f04b4deeafcd0c89b2de490b07)

* change GitRootPath at each command, add can get by subcommand (2023-07-17) [6b37ee4a](https://github.com/convention-change/convention-change-log/commit/6b37ee4afc7183d7c2f26e0e8de3be2591e42d88)

* add GitRepositoryInfo to changelog generate (2023-07-17) [ba094f2f](https://github.com/convention-change/convention-change-log/commit/ba094f2f9ed37d2fa68f9a35899217525dc61e8e)

* init new --more to init config more, remove depends github.com/sinlov/gh-conventional-kit (2023-07-17) [5eba75ef](https://github.com/convention-change/convention-change-log/commit/5eba75ef933e12c0d22de2cfd28ec6d374a072f8)

* change BreakingChanges.IssueReferencesNum to IssueReferencesId and add template_render (2023-07-17) [5b97a1e5](https://github.com/convention-change/convention-change-log/commit/5b97a1e5808dc5d8101773e1be4f8abcf37af406)

* appendMarkdownCommitLink can render by commitUrlFormat (2023-07-16) [294ab02b](https://github.com/convention-change/convention-change-log/commit/294ab02bade3e534e1fbf89bde0ba4403eb8f3fc)

* add ConventionalChangeLogSpec full config and test case (2023-07-16) [3958f4b9](https://github.com/convention-change/convention-change-log/commit/3958f4b9f7a2cef7a706c65571d3af1226f75c74)

* add convention.BreakingChanges support and update test case (2023-07-16) [f9bf552f](https://github.com/convention-change/convention-change-log/commit/f9bf552f67605477e97699906a5543d77d862af9)

* let init not check gt remote info force (2023-07-15) [47fed652](https://github.com/convention-change/convention-change-log/commit/47fed6523882654e5e9422e09175de10003c6e49)

* add cli global flags and convention-change-log (2023-07-15) [8a6e6748](https://github.com/convention-change/convention-change-log/commit/8a6e6748ecf0f82897a3a453c7a13484c92abbcf)

* let convention.NewCommitWithLogSpec can append commit hash link (2023-07-14) [24cc9294](https://github.com/convention-change/convention-change-log/commit/24cc9294fd1c811af72cf283cf9de4b6295dac25)

* add ConventionalChangeLogDesc.VersionNotesUrl to support link of version (2023-07-14) [46708d75](https://github.com/convention-change/convention-change-log/commit/46708d750379f9408c1dedab6c56a469afdf6780)

* add ConventionalChangeLogSpec.TagPrefix support (2023-07-14) [e0785f69](https://github.com/convention-change/convention-change-log/commit/e0785f6921ea2d3698ae82f8c077b6a04a89f685)

* add LoadConventionalChangeLogSpecByData to load by json can auto add sort by default (2023-07-12) [6a834e65](https://github.com/convention-change/convention-change-log/commit/6a834e65106b788fb4dccc65ed19fd9672c8e6c4)

* basic convention and changelog (2023-07-12) [0a6f8778](https://github.com/convention-change/convention-change-log/commit/0a6f8778f72b8ec3d3f2c43fc30fa87838c4808c)

### ♻ Refactor

* move clone-url and add show more run info (2023-07-26) [f605e74d](https://github.com/convention-change/convention-change-log/commit/f605e74dbf5535391d45987e2423c4758c879538)

* move package to github.com/convention-change/convention-change-log (2023-07-16) [0db2257b](https://github.com/convention-change/convention-change-log/commit/0db2257b55c02802700120b4b6ed0e8b9c5da058)

* change convention.Types json format (2023-07-15) [ba0b515c](https://github.com/convention-change/convention-change-log/commit/ba0b515c2c9969bf02378bfe722e816941c545ee)

### 👷‍ Build System

* change changelog/reader_test.go at windows test check (2023-07-19) [8a262b31](https://github.com/convention-change/convention-change-log/commit/8a262b313cd7c1c89398eb36136dd96f79871e6b)

* bump github.com/urfave/cli/v2 from 2.23.7 to 2.25.7 (2023-07-16) [8669f168](https://github.com/convention-change/convention-change-log/commit/8669f168e59694b4d0a592b4037cbe43dbff5385)

* bump github.com/gookit/color from 1.5.3 to 1.5.4 (2023-07-21) [8be9c620](https://github.com/convention-change/convention-change-log/commit/8be9c620d4f51cd056b83fab19ad1a06e5c0f9ba)

* bump github.com/go-git/go-git/v5 from 5.7.0 to 5.8.0 (2023-07-24) [102624a5](https://github.com/convention-change/convention-change-log/commit/102624a50b4ba7786d422d29471575601a16010a)

* bump github.com/go-git/go-git/v5 from 5.8.0 to 5.8.1 (2023-07-27) [80f5f2bb](https://github.com/convention-change/convention-change-log/commit/80f5f2bbc231931e72621f2738474f7c9daa0c62)
