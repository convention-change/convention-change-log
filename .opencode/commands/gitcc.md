---
description: 'Generate and submit code using the Conventional Commits specification'
allowed-tools: [
  'Bash(git status:*)',
  'Bash(git commit:*)',
  'Bash(git diff:*)',
  'Bash(git log:*)'
]
---

## Usage

```
/gitcc
/gitcc --verify
```

## Options

`--verify`: Run git hooks before commit if has lefthook , Husky and so on.

# Instructions

All replies and generated content should be in English.

**Each submission must generate a submission text for display**，Prompt how to commit，Each field is presented with the `directly submit` / `modify the content` / `not submit`, And perform the corresponding operation based on the selection.

If possible, please explain the details of the changes in the submission.

The submission information should briefly summarize the changes, followed by a blank line, and then a more detailed explanation.

The detailed explanation should be presented in the form of paragraphs or multiple paragraphs, with a `-` at the beginning of each paragraph.

If there are engineering custom rules, such as the `.commitlintrc.*` file, the reference configuration corrections generate content that meets the commit rules.

# Command: git Conventional Commit

Creates well-formatted commits with conventional commit messages and emoji.

Please help me generate a commit message that conforms to the Conventional Commits specification based on the current changes in the staging area and execute the commit.

#use-team-conventional-commits

**Current Git status and changes:**
!`git status`
!`git diff --cached`

#project-context

## Emoji Map

✨ feat | 🐛 fix | 📝 docs | 💄 style | ♻️ refactor | ⚡ perf | ✅ test | 🔧 chore | 🚀 ci | 🚨 warnings | 🔒️ security | 🚚 move | 🏗️ architecture | ➕ add-dep | ➖ remove-dep | 🌱 seed | 🧑‍💻 dx | 🏷️ types | 👔 business | 🚸 ux | 🩹 minor-fix | 🥅 errors | 🔥 remove | 🎨 structure | 🚑️ hotfix | 🎉 init | 🔖 release | 🚧 wip | 💚 ci-fix | 📌 pin-deps | 👷 ci-build | 📈 analytics | ✏️ typos | ⏪️ revert | 📄 license | 💥 breaking | 🍱 assets | ♿️ accessibility | 💡 comments | 🗃️ db | 🔊 logs | 🔇 remove-logs | 🙈 gitignore | 📸 snapshots | ⚗️ experiment | 🚩 flags | 💫 animations | ⚰️ dead-code | 🦺 validation | ✈️ offline

## Commit Format

The Conventional Commits specification is a lightweight convention on top of commit messages. It provides an easy set of rules for creating an explicit commit history; which makes it easier to write automated tools on top of. This convention dovetails with [SemVer](http://semver.org/), by describing the features, fixes, and breaking changes made in commit messages.

The commit message should be structured as follows:

---

```
<type>[optional scope]: <emoji> <description>

[optional body]

[optional footer(s)]
```

---

`<type>: <emoji> <description>`

**Types:**

- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation
- `style`: Formatting
- `refactor`: Code restructuring
- `perf`: Performance
- `test`: Tests
- `chore`: Build/tools

**Rules:**

- Imperative mood ("add" not "added")
- First line <72 chars
- Atomic commits (single purpose)
- Split unrelated changes

## Notes

- if has lefthook or other git hook tools, run handles pre-commit checks
- Only commit staged files if any exist
- Analyze diff for splitting suggestions
- **NEVER add signature to commits**
