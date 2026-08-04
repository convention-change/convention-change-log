---
description: 'List all tasks under .shared/tasks/ and their claim status (owner / status)'
allowed-tools:
  [
    'Bash(bash:*)',
    'Bash(find:*)',
    'Bash(awk:*)',
    'Bash(agentlocks:*)',
    'Read(.)',
  ]
---

## Usage

```
/task-list
```

## Options

No arguments. Lists all tasks.

# Instructions

All replies and generated content should be in English.

List all tasks under the `.shared/tasks/` directory (excluding `README.md`) and show each task's claim status.

## Execution steps

Follow these steps (run each step with the bash tool):

1. **Source the shared library**:
   ```bash
   source .opencode/commands/lib/task_lib.sh
   ```

2. **Iterate the task files** and print a table. For each `.md` file (excluding README.md), read the frontmatter fields `id` / `title` / `status` / `owner` / `priority`, and query agentlocks whether someone currently holds the lock:
   ```bash
   find .shared/tasks -type f -name '*.md' ! -name 'README.md' | sort | while read -r f; do
     id=$(task_read_field "$f" "id")
     title=$(task_read_field "$f" "title")
     status=$(task_read_field "$f" "status")
     owner=$(task_read_field "$f" "owner")
     priority=$(task_read_field "$f" "priority")
     lock_owner=$(task_current_owner "$f")
     # If frontmatter owner is empty but agentlocks shows a holder, show the lock holder
     display_owner="${owner:-$lock_owner}"
     printf '| %-30s | %-10s | %-12s | %-15s |\n' "$id" "$status" "$priority" "${display_owner:-(unclaimed)}"
   done
   ```

3. **Aggregate the output**: present it as a Markdown table with columns: `Task ID` `Status` `Priority` `Claimed by`. Add a line before the table stating the total number of tasks, and a line after the table prompting to use `/task-claim <id>` to claim a task.

## Example output format

```
3 tasks total:

| Task ID                       | Status     | Priority     | Claimed by      |
|-------------------------------|------------|--------------|-----------------|
| 0088-new-sub-command          | open       | medium       | (unclaimed)     |
| 0092-fix-template-lang        | in-progress| high         | sinlovppt       |

Claim a task: /task-claim <id>
```

## Notes

- If `.shared/tasks/` does not exist or has no tasks, output "No tasks yet" and point to `.shared/tasks/README.md` for the task definition spec.
- The `status` field reflects the value in the frontmatter; `Claimed by` prefers the frontmatter `owner`, falling back to the live agentlocks lock holder if empty.
- Do not modify any files.
