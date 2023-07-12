## softprops/action-gh-release

- [https://github.com/softprops/action-gh-release](https://github.com/softprops/action-gh-release)

`gh-prerelease.yml`

```yml
name: gh-prerelease

on:
  push:
    tags:
      - '*' # Push events to matching *, i.e. 1.0.0 v1.0, v20.15.10

permissions:
  contents: write
  discussions: write

jobs:
  gh-prerelease:
    name: gh-prerelease
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3

      - uses: softprops/action-gh-release@master
        name: Create Release
        if: startsWith(github.ref, 'refs/tags/')
        with:
          ## with permissions to create releases in the other repo
          token: "${{ secrets.GITHUB_TOKEN }}"
          prerelease: true
          # body: "this is pre-release"
          # body_path: ${{ github.workspace }}-CHANGELOG.txt
          # https://github.com/isaacs/node-glob
          # files: |
          #   **/*.tar.gz
          #   **/*.sha256
```