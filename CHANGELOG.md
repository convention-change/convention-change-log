# Changelog

All notable changes to this project will be documented in this file. See [convention-change-log](https://github.com/convention-change/convention-change-log) for commit guidelines.

## [1.8.3](https://github.com/convention-change/convention-change-log/compare/1.8.2...v1.8.3) (2024-09-29)

### 🐛 Bug Fixes

* monorepo-pkg-path replace package.json error ([78db6282](https://github.com/convention-change/convention-change-log/commit/78db628238f190b7e37561e50a4e0ba3c95208a3)), fix [#48](https://github.com/convention-change/convention-change-log/issues/48)

### 📝 Documentation

* update README to reflect v1.8.3 installation ([513c7cb6](https://github.com/convention-change/convention-change-log/commit/513c7cb62c3cf0b9cf16b82b567a3e96002dfc9d))

## [1.8.2](https://github.com/convention-change/convention-change-log/compare/1.8.1...v1.8.2) (2024-08-31)

### 🐛 Bug Fixes

* files with execution permissions, and pass build ID to goldie for change log generation ([c99897a7](https://github.com/convention-change/convention-change-log/commit/c99897a7d80672d3ba60bf47babbb9283da87ca7))

### ♻ Refactor

* use color.Print for single line outputs ([4f1b9f64](https://github.com/convention-change/convention-change-log/commit/4f1b9f64d5a6846fae3814ded67dd1bd7ec84b0f))

### 👷‍ Build System

* bump codecov/codecov-action from 4.4.1 to 4.5.0 ([9c00265e](https://github.com/convention-change/convention-change-log/commit/9c00265eeb9b2ef75de85c500d03f603cf140d5b))

* bump codecov/codecov-action from 4.4.0 to 4.4.1 ([758c23af](https://github.com/convention-change/convention-change-log/commit/758c23af532db982fcb759b978bd22c0a94f2b56))

## [1.8.1](https://github.com/convention-change/convention-change-log/compare/1.8.0...v1.8.1) (2024-05-18)

### 🐛 Bug Fixes

* add flag `--skip-worktree-check` worktree dirty check will error or want skip ([9bfd3c15](https://github.com/convention-change/convention-change-log/commit/9bfd3c152a578b064ed7945035ad6c8f3c2bc31e)), fix [#42](https://github.com/convention-change/convention-change-log/issues/42)

### 📝 Documentation

* add local repository dirty check (v1.8+) features ([91e9597b](https://github.com/convention-change/convention-change-log/commit/91e9597b3d8b9137d81304d4cf39493ae5b6267b))

### 👷‍ Build System

* github.com/sinlov-go/go-git-tools v1.13.0 ([36eb1663](https://github.com/convention-change/convention-change-log/commit/36eb1663c200dadb1509645ec47e77084336e3df))

## [1.8.0](https://github.com/convention-change/convention-change-log/compare/1.7.0...v1.8.0) (2024-05-17)

### 🐛 Bug Fixes

* show log not find old tag at debug ([bdd6c4a8](https://github.com/convention-change/convention-change-log/commit/bdd6c4a84d0c55f6092e02e04ea26a46963d93a6))

### ✨ Features

* local repo is dirty and check git submodule is dirty ([db9f3918](https://github.com/convention-change/convention-change-log/commit/db9f3918809f48c37fba73b43ced87fd49314627)), feat [#40](https://github.com/convention-change/convention-change-log/issues/40)

* `--git-info-scheme` and `.versionrc` `cover-git-info-scheme` to change log of git url ([cb02596a](https://github.com/convention-change/convention-change-log/commit/cb02596aa57733e9aaf10d975a51bd6fe85545b4))

* update full build pipline ([81f6c5c7](https://github.com/convention-change/convention-change-log/commit/81f6c5c7c3eca60d636519a51950a71773cb2c3d))

### 📝 Documentation

* fix u18n link ([4f508855](https://github.com/convention-change/convention-change-log/commit/4f508855f958d2decac3e78cfa556177b8926d5d))

* change badge to use 中文 ([64b65ebb](https://github.com/convention-change/convention-change-log/commit/64b65ebb3d9e54fb5c9fdd858728f6b5f75c7f53))

* try i18n ([ec2f2273](https://github.com/convention-change/convention-change-log/commit/ec2f22732b1459ccf8c014d250a8e680a2e51e7e))

* add usage of doc by i18n ([2cbc9d84](https://github.com/convention-change/convention-change-log/commit/2cbc9d84de3c60a117ec75f1a1a5d73878b17472))

### ♻ Refactor

* add `ChangeLogGenerator` to management change log ([e8426132](https://github.com/convention-change/convention-change-log/commit/e8426132931557f5e69a2f8863ed8eca4494ce35))

### 👷‍ Build System

* bump codecov/codecov-action from 4.3.0 to 4.4.0 ([71a4831b](https://github.com/convention-change/convention-change-log/commit/71a4831b428f4c1644134e0763227ae710cfbd13))

* bump golangci/golangci-lint-action from 5 to 6 ([0a3a34b3](https://github.com/convention-change/convention-change-log/commit/0a3a34b306dc28db13b92e2c8769c7db687010c2))

* bump codecov/codecov-action from 4.1.1 to 4.3.0 ([24170be0](https://github.com/convention-change/convention-change-log/commit/24170be032911f7f335151fa2756a94aa95f254f))

* bump golangci/golangci-lint-action from 4 to 5 ([e70b6ab6](https://github.com/convention-change/convention-change-log/commit/e70b6ab600e63fd8271b05fde55d75c33be8e9c7))

* bump convention-change/conventional-version-check ([1ae3c5ef](https://github.com/convention-change/convention-change-log/commit/1ae3c5ef8b3b5b44f79095a657d426cce3ba70ea))

* bump github.com/urfave/cli/v2 from 2.27.1 to 2.27.2 ([8bbbec69](https://github.com/convention-change/convention-change-log/commit/8bbbec69a47cb084e657dcf48c2d3a340fe0c077))

* bump github.com/go-git/go-git/v5 from 5.11.0 to 5.12.0 ([9916e040](https://github.com/convention-change/convention-change-log/commit/9916e0403c995c73017be3687ae8e53e7b98c713))

* bump github.com/sinlov-go/go-common-lib from 1.6.0 to 1.7.0 ([77c1c084](https://github.com/convention-change/convention-change-log/commit/77c1c084182b6f8a9723ca35d31dd412e00baba2))

* bump codecov/codecov-action from 4.1.0 to 4.1.1 ([2c6d2ac6](https://github.com/convention-change/convention-change-log/commit/2c6d2ac6a8e54d54c39c84cb003d700a2012c3e4))

## [1.7.0](https://github.com/convention-change/convention-change-log/compare/1.6.0...v1.7.0) (2024-03-20)

### BREAKING CHANGE:

* default in dry-run mode, disable by --dry-run-disable, --auto-push will ignore all

### ✨ Features

* default open dry-run mode, and disable by --dry-run-disable, --auto-push will ignore dry-run ([0f2be5c7](https://github.com/convention-change/convention-change-log/commit/0f2be5c7c8f125454c796305ca4004589b3109ce))

### 📝 Documentation

* change install version at README.md ([fb91939b](https://github.com/convention-change/convention-change-log/commit/fb91939bbfe77795562d6fd9469dd3e2966e109f))

* update usage at README.md ([40ac75bb](https://github.com/convention-change/convention-change-log/commit/40ac75bb3fff59ef4d4db20307e8a228caf22fc7))

## [1.6.0](https://github.com/convention-change/convention-change-log/compare/1.5.1...v1.6.0) (2024-03-02)

### ✨ Features

* change to github.com/sinlov-go/unittest-kit v1.1.0 and update test case template ([76b0e68e](https://github.com/convention-change/convention-change-log/commit/76b0e68ea2c8e54b5807a812c983eafe31ec3fee))

### 📝 Documentation

* update cli usage doc ([639bf9a8](https://github.com/convention-change/convention-change-log/commit/639bf9a80bc29c80c6ba81d86b2a3694f96a1d83))

### 👷‍ Build System

* github.com/stretchr/testify v1.9.0 ([0340016d](https://github.com/convention-change/convention-change-log/commit/0340016db2c3408f23b5fa5ff8706fa6af6a0e91))

* bump codecov/codecov-action from 4.0.0 to 4.1.0 ([58b4be16](https://github.com/convention-change/convention-change-log/commit/58b4be161fca33546e83160861ed7ca02254f0a2))

* bump golangci/golangci-lint-action from 3 to 4 ([7ffdd6a0](https://github.com/convention-change/convention-change-log/commit/7ffdd6a01a1e3b1b60a5d5dde5db968603cc94d9))

* bump codecov/codecov-action from 3.1.4 to 4.0.0 ([227a8680](https://github.com/convention-change/convention-change-log/commit/227a8680d16e55d15cd5f881c6292f74f980e856))

## [1.5.1](https://github.com/convention-change/convention-change-log/compare/1.5.0...v1.5.1) (2024-01-26)

### 📝 Documentation

* change default type docs not hidden ([5b4ca0c0](https://github.com/convention-change/convention-change-log/commit/5b4ca0c078f89518b225a023845fa8857dfe359e)), docs [#11](https://github.com/convention-change/convention-change-log/issues/11)

* update doc ([03a9d1f9](https://github.com/convention-change/convention-change-log/commit/03a9d1f9e1b4dfea10343e8d192621ea739ec162))

## [1.5.0](https://github.com/convention-change/convention-change-log/compare/1.4.2...v1.5.0) (2024-01-23)

### ✨ Features

* change to github.com/sinlov-go/go-git-tools v1.11.0 ([82ed979a](https://github.com/convention-change/convention-change-log/commit/82ed979adb834d11d5662bf41b68024afbce7cfb))

* let global command can try change by monorepo-pkg-path config at method monorepo-pkg-path ([cd157bba](https://github.com/convention-change/convention-change-log/commit/cd157bbaad82b52c95baa541db45f077272824f9))

* add convention.ConventionalChangeLogSpec.MonoRepoPkgPathList and can load by config file ([83d6b50a](https://github.com/convention-change/convention-change-log/commit/83d6b50aebdb19d655d00917f9050308ecbbe883))

### 👷‍ Build System

* update golang full action CI pipline ([500860a4](https://github.com/convention-change/convention-change-log/commit/500860a4935ca1278a0b4e9566ef43e54620195b))

* support actions/upload-artifact/tree/v4 and let action goversion ^1.19 ([1bd7ff4f](https://github.com/convention-change/convention-change-log/commit/1bd7ff4fbb7b5b64ed22280c2f794f7cc0b52eca))

* bump github.com/go-git/go-git/v5 from 5.10.0 to 5.11.0 ([08a2d841](https://github.com/convention-change/convention-change-log/commit/08a2d841bd7d34afe2fa70edc430c1ba5b981757))

* bump github.com/sinlov-go/go-git-tools from 1.9.1 to 1.10.0 ([e728183d](https://github.com/convention-change/convention-change-log/commit/e728183db9685c82816949ef0a557ff104d9532a))

* update to go 1.19+ to suppot new version of build ([bb977d0d](https://github.com/convention-change/convention-change-log/commit/bb977d0d6d6dd9db14b629e8a6634a154d99428f))

* bump github.com/urfave/cli/v2 from 2.25.7 to 2.27.1 ([45550366](https://github.com/convention-change/convention-change-log/commit/455503663af6fdc537051b85d80a5d6e6aa72851))

* bump actions/download-artifact from 3 to 4 ([5bfe2734](https://github.com/convention-change/convention-change-log/commit/5bfe2734600e8175cae1e720a787096048407c88))

* bump actions/setup-go from 4 to 5 ([b54a4fe7](https://github.com/convention-change/convention-change-log/commit/b54a4fe772810f90c22e74998ea57776fe5efc66))

* bump actions/upload-artifact from 3 to 4 ([3659695f](https://github.com/convention-change/convention-change-log/commit/3659695f08319093c67c24a81f495dfd607a977c))

* change golangci/golangci-lint-action use version latest ([fa79a999](https://github.com/convention-change/convention-change-log/commit/fa79a99931f03affc04729a2555bab165070474d))

* bump github.com/go-git/go-git/v5 from 5.9.0 to 5.10.0 ([48d8cff0](https://github.com/convention-change/convention-change-log/commit/48d8cff05e60c3c750dd242028abd81a85dcdd56))

* bump github.com/sinlov-go/go-common-lib from 1.4.0 to 1.5.0 ([630b77b2](https://github.com/convention-change/convention-change-log/commit/630b77b2337eed09a62ed7e11d96347724b48f27))

* bump actions/checkout from 3 to 4 ([7289c471](https://github.com/convention-change/convention-change-log/commit/7289c47102e36aa2686a09d0322293aae7df7551))

* bump github.com/go-git/go-git/v5 from 5.8.1 to 5.9.0 ([8fec5b90](https://github.com/convention-change/convention-change-log/commit/8fec5b90184b322b5087bc659955a793e1b361d5))

## [1.4.2](https://github.com/convention-change/convention-change-log/compare/1.4.1...v1.4.2) (2023-08-15)

### 🐛 Bug Fixes

* add more error log show at do git ([2249d141](https://github.com/convention-change/convention-change-log/commit/2249d14178809a29f8e8cdf4e419c4c1ff370ffc))

## [1.4.1](https://github.com/convention-change/convention-change-log/compare/1.4.0...v1.4.1) (2023-08-09)

### ♻ Refactor

* let out log more clear and show more fix info at flag --auto-push ([7fdfa0c9](https://github.com/convention-change/convention-change-log/commit/7fdfa0c9b98571bbeabb29fd593c88bccffbc430))

## [1.4.0](https://github.com/convention-change/convention-change-log/compare/1.3.1...v1.4.0) (2023-08-07)

### ✨ Features

* convention.NewCommitWithLogSpec generate head at word `:` support ([9540a3a4](https://github.com/convention-change/convention-change-log/commit/9540a3a43f4f16c627d4cdb03d4d2e6c550f8b15))

## [1.3.1](https://github.com/convention-change/convention-change-log/compare/1.3.0...v1.3.1) (2023-08-05)

### 🐛 Bug Fixes

* fix doc not set at subcommand at read-latest ([6780aacf](https://github.com/convention-change/convention-change-log/commit/6780aacfb493c1cee17efc88b3ddcd54570e5a2b))

## [1.3.0](https://github.com/convention-change/convention-change-log/compare/1.2.0...v1.3.0) (2023-08-05)

### ✨ Features

* add git command error can use command notice ([f02bb653](https://github.com/convention-change/convention-change-log/commit/f02bb653112152ad1e3f1fcf1973185919ca89b0))

## [1.2.0](https://github.com/convention-change/convention-change-log/compare/1.1.0...v1.2.0) (2023-08-01)

### BREAKING CHANGE:

* after this version each commit will not contain commit date information

### 🐛 Bug Fixes

* add convention.LoadConventionalChangeLogSpecByPath to fast load ConventionalChangeLogSpec ([b4fe99d0](https://github.com/convention-change/convention-change-log/commit/b4fe99d02cf1592e76ab636a6f2e372bb83562af))

### ✨ Features

* change doGit std print info ([96b29ae8](https://github.com/convention-change/convention-change-log/commit/96b29ae810b45abb43d8eb796fa54b11c2a1d294))

* change format of convention.NewCommitWithLogSpec not contain AddAuthorDate ([689a502e](https://github.com/convention-change/convention-change-log/commit/689a502eae264128f25ae77a6c90754b40a5ae69))

* use convention.LoadConventionalChangeLogSpecByPath to init ConventionalChangeLogSpec ([07ed5be0](https://github.com/convention-change/convention-change-log/commit/07ed5be0b3aea50a23bad29a58fd0f70ab8a15a7))

## [1.1.0](https://github.com/convention-change/convention-change-log/compare/v1.0.3...v1.1.0) (2023-07-31)

### 👷‍ Build System

* **gomod:** github.com/sinlov-go/go-git-tools v1.9.1 ([248572a](https://github.com/convention-change/convention-change-log/commit/248572a30b8460d8804e0b275476ba31f7b1f8cd))

## [1.0.3](https://github.com/convention-change/convention-change-log/compare/1.0.2...v1.0.3) (2023-07-30)

### 🐛 Bug Fixes

* change go-build-check-main for test (2023-07-30) [80f7d59c](https://github.com/convention-change/convention-change-log/commit/80f7d59c84ad0460344307e824de5442d286e243)

## [1.0.2](https://github.com/convention-change/convention-change-log/compare/1.0.1...v1.0.2) (2023-07-30)

### 🐛 Bug Fixes

* fix go-release-platform.yml build of archive (2023-07-30) [ed57f54b](https://github.com/convention-change/convention-change-log/commit/ed57f54b450c2d7588bfa7e4edb0e7bea93cdd69)

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
