---
description: 'Interactively create a new task file under .shared/tasks/ (with id/title/status/priority/type/owner)'
allowed-tools:
  [
    'Bash(bash:*)',
    'Bash(source:*)',
    'Bash(find:*)',
    'Bash(mkdir:*)',
    'Bash(date:*)',
    'Read(.)',
    'Edit(.)',
    'Write(.)',
    'Question',
  ]
---

## Usage

```
/task-create
```

## Options

No arguments. The command collects task metadata via interactive Q&A.

# Instructions

All replies and generated content should be in English.

Interactively create a new task file, saved under `.shared/tasks/<category>/`. Each field is prompted to the user via the `question` tool (with defaults and selectable options). Once all questions are answered, generate the frontmatter file and write it to disk.

## Preflight checks

1. **Source the shared library** (used to reuse frontmatter validation logic):
   ```bash
   source .opencode/commands/lib/task_lib.sh
   ```
2. Confirm the `.shared/tasks/` directory exists; create it if it does not.

## Execution steps

Prompt the user for the following 6 fields in order. Each field is presented with the `question` tool, offering recommended options and custom input (`custom: true`).

### 1. Task id

Prompt for the task id (a number); it will be normalized to 4 digits (zero-padded if shorter; kept as-is if longer than 4 digits).

- Recommended option: scan existing `.shared/tasks/` for the maximum id and offer `max_id+1` as the recommendation.
- Validation: if the zero-padded id already exists in existing task filenames (prefix match `{id}-`), report an error and re-ask.

### 2. Task title

Prompt for the task title (a natural-language description).

- Conversion rule: convert to lowercase, keep only letters and digits, replace spaces with `-`, and strip other special characters. Example: `"Add new sub command"` -> `add-new-sub-command`.
- Validation: if the converted slug conflicts with an existing filename sharing the same id prefix (i.e. `{id}-{slug}.md` already exists), report an error and re-ask.

### 3. Task status

Prompt for the task status; options are `open` / `in-progress` / `done`, default `open`.

- Recommended options: `open` (Recommended) / `in-progress` / `done`.

### 4. Task priority

Prompt for the task priority; options are `low` / `medium` / `high`, default `medium`.

- Recommended options: `medium` (Recommended) / `high` / `low`.

### 5. Task type

Prompt for the task type (used both as the `category` field and the subdirectory name); options are:
`feat` / `fix` / `new` / `docs` / `style` / `refactor` / `test` / `ci` / `chore` / `revert`, default `feat`.

- Recommended options: `feat` (Recommended) / `fix` / `new` / `docs` / `refactor` / `test` / `ci` / `chore` / `style` / `revert`.
- **Users are allowed to add custom types** (`custom: true`; the user may enter a value outside the list). A custom type must be a valid directory name (lowercase letters, digits, `-`).
- The type value is the `category` field value and also the `.shared/tasks/<category>/` subdirectory name.

### 6. Task owner

Prompt for the task owner; default is empty.

- Recommended options: `(empty)` (Recommended) / current `AGENTLOCKS_AGENT_ID` (if set) / custom input.

## File generation

Once all questions are answered, generate the file with these steps:

1. **Assemble the filename**: `{id}-{type}-{slug}.md`, where `id` is the zero-padded 4-digit number, `type` is the task type, and `slug` is the title conversion result.
2. **Assemble the file path**: `.shared/tasks/{type}/{filename}`. Create the subdirectory if it does not exist.
3. **Get the current date**: `date +%Y-%m-%d`.
4. **Write the frontmatter + body**:

   ````
   ---
   id: {id}-{type}-{slug}
   title: {original task title}
   status: {task status}
   owner: "{task owner}"
   locked_at: ""
   priority: {task priority}
   category: {task type}
   created_at: "{YYYY-MM-DD}"
   ---

   TODO: fill in the task content
   ````

   Notes:
   - The `id` field = the filename (without `.md`); it must exactly match the filename.
   - `title` uses the user's original input (unconverted); quote it per YAML string rules if it contains special characters.
   - When `owner` is empty, write `""`; when non-empty, wrap it in double quotes.
   - `created_at` is a double-quoted `YYYY-MM-DD` string.

5. **Output confirmation**: show the created file path, the full frontmatter, and prompt to claim it next with `/task-claim <id>`.

## Example output format

```
✅ Created task file

  Path: .shared/tasks/feat/0089-add-new-feature.md

  ---
  id: 0089-feat-add-new-feature
  title: Add new feature
  status: open
  owner: ""
  locked_at: ""
  priority: medium
  category: feat
  created_at: "2026-07-20"
  ---

  TODO: fill in the task content

Next step: /task-claim 0089-feat-add-new-feature
```

## Notes

- If the zero-padded id already exists (a file with the `{id}-` prefix), **do not** overwrite it; report an error and re-ask for the id.
- The slug converted from the title must be non-empty; if the title consists entirely of special characters causing the slug to be empty, report an error and re-ask.
- Task types support custom values, but must be valid directory names (`^[a-z0-9][a-z0-9-]*$`).
- Do not auto-claim the task (do not call agentlocks); after creation the task is in the `open` unlocked state, and the user should manually run `/task-claim`.
- Do not modify other task files.
