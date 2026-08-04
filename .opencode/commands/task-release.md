---
description: 'Release a task lock: release the agentlocks lock + revert the task to open (owner cleared)'
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
/task-release <task-id>
```

## Options

- `<task-id>`: The task ID.

# Instructions

All replies and generated content should be in English.

Release the specified task: release the agentlocks lock and restore the task frontmatter to `open` (clearing owner / locked_at). Use this for "I claimed it but I'm not working on it for now - let someone else take it."

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
   CUR_OWNER=$(task_read_field "$FILE" "owner")
   ME=$(task_agent_id)
   LOCK_ID=$(task_current_lock_id "$FILE")
   ```
   - If `CUR_OWNER` is non-empty and not equal to `$ME`: report "Task is claimed by $CUR_OWNER, you are not authorized to release it" and stop.
   - If `CUR_OWNER` is empty but `LOCK_ID` is non-empty (frontmatter and lock are out of sync): warn, and ask whether to still release the current agentlocks lock. Only release if `LOCK_OWNER` (via `task_current_owner`) equals `$ME` or is empty.

4. **Release the agentlocks lock**:
   ```bash
   if [[ -n "$LOCK_ID" ]]; then
     task_lock_release "$LOCK_ID"
   fi
   ```
   Also run `agentlocks release --mine` as a fallback to clean up any lingering locks owned by this machine.

5. **Update the frontmatter** (revert to open):
   ```bash
   task_write_field "$FILE" "status" "open"
   task_write_field "$FILE" "owner" ""
   task_write_field "$FILE" "locked_at" ""
   ```

6. **Output confirmation**.

## Example output format

Success:
```
🔓 Released task 0088-new-sub-command
   Task reverted to open, others may claim it.
```

Not authorized:
```
❌ Cannot release task 0088-new-sub-command
   Reason: Task is claimed by other-dev, only the claimer can release it.
```

## Notes

- `/task-release` reverts the task to `open` (not done; claimable by others).
- If the task is actually complete, use `/task-done` (which sets status=done and releases the lock); do not use this command.
- After releasing, it is recommended to `git add .shared/tasks/... && git commit` the frontmatter change so the team can see the task is claimable again.
