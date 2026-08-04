---
description: 'View task and lock status: all currently held agentlocks locks + a task list summary, and refresh your own locks'
allowed-tools:
  [
    'Bash(bash:*)',
    'Bash(source:*)',
    'Bash(agentlocks:*)',
    'Read(.)',
  ]
---

## Usage

```
/task-status
```

## Options

No arguments.

# Instructions

All replies and generated content should be in English.

Show the current agentlocks lock state + locks I hold + refresh (renew) locks that are about to expire.

## Execution steps

1. **Source the shared library**:
   ```bash
   source .opencode/commands/lib/task_lib.sh
   ```

2. **Show the current identity**:
   ```bash
   echo "Current developer: $(task_agent_id)"
   ```

3. **Show all active locks** (raw):
   ```bash
   agentlocks status 2>/dev/null
   ```
   Parse and present as a table: `lock_id` `resource` `owner` `remaining TTL / status`.

4. **Refresh locks I hold**:
   ```bash
   agentlocks refresh --mine 2>/dev/null
   ```
   Report which locks were refreshed.

5. **Clean up expired locks** (preview only; do not actually prune unless the user confirms):
   ```bash
   agentlocks prune --dry-run 2>/dev/null
   ```
   If there are expired locks, prompt the user that they can run `agentlocks prune` manually to clean up.

6. **Show tasks I have claimed** (iterate `.shared/tasks/` and find entries where owner == me):
   ```bash
   find .shared/tasks -type f -name '*.md' ! -name 'README.md' | while read -r f; do
     owner=$(task_read_field "$f" "owner")
     me=$(task_agent_id)
     if [[ "$owner" == "$me" ]]; then
       id=$(task_read_field "$f" "id")
       status=$(task_read_field "$f" "status")
       printf '  - %s [%s]\n' "$id" "$status"
     fi
   done
   ```

## Example output format

```
Current developer: sinlovppt

## Active locks (2)

| lock_id                          | resource                                      | owner      | status |
|----------------------------------|-----------------------------------------------|------------|--------|
| lock_20260719T192035Z_3d500b1a   | .shared/tasks/issure/0088-new-sub-command.md  | sinlovppt  | held   |

## Tasks I have claimed

  - 0088-new-sub-command [in-progress]

✅ Refreshed 1 lock (TTL reset to 10 minutes)
```

## Notes

- `/task-status` automatically refreshes the locks you hold, acting as a "heartbeat renewal." Run this command periodically during long tasks to prevent locks from expiring and being reclaimed by others.
- If you find one of your locks has been reclaimed by someone else (status shows reclaimable, or the owner changed), it means the lock expired and was taken over; you need to `/task-claim` again.
- Do not modify any files (except agentlocks internal lock state).
