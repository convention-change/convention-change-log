# `.shared/tasks/` - Shared tasks directory

This directory holds task definitions for **multi-developer collaborative claiming**. Each task is one Markdown file, paired with the lightweight file lock from [Agentlocks](https://github.com/simke9445/agentlocks) to coordinate who-is-working-on-what under mutual exclusion.

## Directory layout

```
.shared/tasks/
├── README.md                      ← This file (spec; tracked in git)
├── issure/                        ← Subdirectory per category (issue / feature / hotfix ...)
│   ├── 0088-new-sub-command.md    ← Task definition file (tracked; shared with the team)
│   └── ...
└── <other-category>/
    └── ...
```

## Task definition file format

Every `.md` file **must** begin with YAML frontmatter describing the task metadata; the body is the task detail.

```markdown
---
id: 0088-new-sub-command            # Required; matches the filename (without extension); globally unique
title: New agent subcommand         # Required; short title
status: open                        # Required; open | in-progress | done | cancelled
owner: ""                           # Claimer's AGENTLOCKS_AGENT_ID; empty means unclaimed
locked_at: ""                       # Claim time (ISO 8601); empty means unclaimed
priority: medium                    # Optional; high | medium | low
category: issure                    # Optional; matches the containing subdirectory name
created_at: "2026-07-20"            # Optional; creation date
---

## Task description

(Body: steps, acceptance criteria, related issue, etc.)
```

> The `status` / `owner` / `locked_at` fields are maintained automatically by the
> `/task-claim`, `/task-release`, and `/task-done` commands; **do not hand-edit them**.
> All other fields are maintained manually.

## Lock mechanism (Agentlocks)

When claiming a task, `/task-claim <id>` will:

1. Call `agentlocks acquire '.shared/tasks/<category>/<id>.md' --reason '...'`
   to obtain an advisory lock on that task file (default TTL 10 minutes, max 30 minutes).
2. On a successful lock, change the frontmatter `status` to `in-progress`, `owner` to the
   current `AGENTLOCKS_AGENT_ID`, and `locked_at` to the current time.
3. On a lock conflict (already held by someone else), the command reports an error and
   shows the current owner; it does **not** modify the file.

Lock state is recorded under `.agentlocks/locks/` (already in `.gitignore`; **not tracked**).

## Per-developer prerequisites

Set your identity in your shell profile (`~/.zshrc` / `~/.bashrc`):

```bash
export AGENTLOCKS_AGENT_ID="sinlovppt"   # Recommended: your git username or a unique abbreviation
```

If unset, Agentlocks falls back to `$USER`, then to the `dev` prefix
(multiple people sharing `dev` will overwrite each other's locks; **make sure each person sets their own**).

## Daily workflow

```bash
# 1. List all tasks and their claim status
/task-list

# 2. Claim a task (auto-lock + update frontmatter)
/task-claim 0088-new-sub-command

# 3. Show the locks currently held
/task-status

# 4. Release the lock when done (the task stays in-progress, ready for the follow-up commit)
/task-release 0088-new-sub-command

# 5. Mark the task done (auto-releases the lock + sets status=done)
/task-done 0088-new-sub-command
```

## Lock expiry and conflicts

- Agentlocks locks have a TTL (default 10 minutes). For long-running tasks, periodically run `/task-status` to trigger
  `agentlocks refresh` and renew (the commands have built-in auto-refresh).
- After a lock expires it enters a grace window (90s); after that, others can use `--reclaim` to forcibly take it over.
- Locks are **advisory**: they do not physically prevent you from writing the file; they rely on the team honoring the convention.
- Locks **only coordinate the same local worktree**; they do not work across machines. For distributed collaboration, use Git branches.

## Conventions

- Task filename: `<id>-<kebab-case-description>.md`, where `id` is typically the issue number.
- One task per file; claiming = locking that file.
- Do not hand-edit `status` / `owner` / `locked_at`; use the commands.
- Before committing a change to a task definition, confirm no one holds the lock for that task (`/task-status`).
