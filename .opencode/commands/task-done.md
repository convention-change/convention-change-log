---
description: 'Mark a task done: status=done + release the agentlocks lock'
allowed-tools:
  [
    'Bash(bash:*)',
    'Bash(source:*)',
    'Bash(agentlocks:*)',
    'Bash(mktemp:*)',
    'Bash(mv:*)',
    'Read(.)',
    'Edit(.)',
  ]
---

## Usage

```
/task-done <task-id>
```

## Options

- `<task-id>`: The task ID.

# Instructions

All replies and generated content should be in English.

Mark the specified task as done (`status=done`) and release its agentlocks lock. Use this only when the task is actually complete.

## Execution steps

Arguments: `$ARGUMENTS` (the first token is used as task-id)

1. **Source the shared library**:
   ```bash
   source .opencode/commands/lib/task_lib.sh
   ```

2. **Find the task file**:
   ```bash
   FILE=$(task_find_file "$TASK_ID") || exit 1
   ```

3. **Check ownership**:
   ```bash
   CUR_STATUS=$(task_read_field "$FILE" "status")
   CUR_OWNER=$(task_read_field "$FILE" "owner")
   ME=$(task_agent_id)
   LOCK_ID=$(task_current_lock_id "$FILE")
   ```
   - If `CUR_STATUS` is already `done`: report "Task is already done, no action needed" and stop.
   - If `CUR_OWNER` is non-empty and not equal to `$ME`: report "Task is claimed by $CUR_OWNER, only the claimer can mark it done" and stop.

4. **Release the agentlocks lock**:
   ```bash
   if [[ -n "$LOCK_ID" ]]; then
     task_lock_release "$LOCK_ID"
   fi
   agentlocks release --mine 2>/dev/null || true
   ```

5. **Update the frontmatter**:
   ```bash
   task_write_field "$FILE" "status" "done"
   task_write_field "$FILE" "owner" ""
   task_write_field "$FILE" "locked_at" ""
   ```

6. **Output confirmation**, and prompt to commit the frontmatter change.

## Example output format

Success:
```
✅ Task 0088-new-sub-command marked done, lock released.
   Please commit the change: git add .shared/tasks/issure/0088-new-sub-command.md && git commit
```

## Notes

- `/task-done` is a terminal action: once `status` becomes `done`, you can no longer `/task-claim` it. To reopen, manually edit the frontmatter to `status: open`.
- After completion, commit the frontmatter change so the team can see the task completion in the repository history.
- If the task is not actually finished but only paused, use `/task-release` instead of this command.
