---
description: 'Claim a task: acquire the agentlocks file lock + update task frontmatter (status=in-progress, owner=me)'
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
/task-claim <task-id>
```

## Options

- `<task-id>`: The task ID, i.e. the filename (without `.md`) of `.shared/tasks/**/<id>.md`. For example `0088-new-sub-command`.

# Instructions

All replies and generated content should be in English.

Claim the specified task: atomically acquire the agentlocks file lock and update the task frontmatter. If the lock is already held by someone else, **do not** modify the file; report the conflict.

## Preflight checks

1. Confirm the `AGENTLOCKS_AGENT_ID` environment variable is set (recommended: your git username). If unset, prompt the user to configure it in their shell profile and ask whether to proceed using `$USER` as the identity.
2. Confirm the `agentlocks` command is available (`command -v agentlocks`). If unavailable, prompt `npm install -g agentlocks` (requires Node >= 22.18).

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

3. **Check the current state**:
   ```bash
   CUR_STATUS=$(task_read_field "$FILE" "status")
   CUR_OWNER=$(task_read_field "$FILE" "owner")
   LOCK_OWNER=$(task_current_owner "$FILE")
   ```
   - If `CUR_STATUS` is `done` or `cancelled`: report "Task is $CUR_STATUS, cannot be claimed" and stop.
   - If `CUR_OWNER` is non-empty and not equal to the current `task_agent_id`: report "Task frontmatter shows it has been claimed by $CUR_OWNER" and stop.
   - If `LOCK_OWNER` is non-empty and not equal to the current `task_agent_id`: report "agentlocks lock is held by $LOCK_OWNER, conflict" and stop.

4. **Acquire the lock** (critical atomic step):
   ```bash
   LOCK_ID=$(task_lock_acquire "$FILE" "claim task $TASK_ID")
   ```
   If this fails, report the agentlocks conflict information and stop. **Do not** modify the frontmatter.

5. **Update the frontmatter** (only after a successful lock):
   ```bash
   ME=$(task_agent_id)
   NOW=$(task_now_iso)
   task_write_field "$FILE" "status" "in-progress"
   task_write_field "$FILE" "owner" "$ME"
   task_write_field "$FILE" "locked_at" "$NOW"
   ```

6. **Output confirmation**: show claimer, task ID, lock_id, and TTL hint.

## Example output format

Success:
```
✅ Claimed task 0088-new-sub-command
   Claimed by: sinlovppt
   Lock ID:     lock_20260719T192035Z_3d500b1a
   TTL:         10 minutes (use /task-status to refresh, /task-release to release)

Task details: .shared/tasks/issure/0088-new-sub-command.md
```

Conflict:
```
❌ Cannot claim task 0088-new-sub-command
   Reason: agentlocks lock is held by other-dev
   Suggestion: wait for them to /task-release or for the lock to expire (default 10 minutes + 90s grace)
```

## Notes

- The lock TTL defaults to 10 minutes. For long-running tasks, periodically run `/task-status` (which triggers a refresh/renewal).
- Locks are advisory: they do not physically prevent others from writing the file; they rely on the team honoring the convention.
- The frontmatter `owner` / `status` / `locked_at` fields must only be modified via /task-* commands; do not hand-edit them.
