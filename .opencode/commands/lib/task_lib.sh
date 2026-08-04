#!/usr/bin/env bash
# .opencode/commands/lib/task_lib.sh
# Shared helpers for /task-* commands. Safe to source from any command markdown.
#
# Env:
#   AGENTLOCKS_AGENT_ID  - developer identity (fallback: $USER, then "dev")
#   AGENTLOCKS_BIN       - agentlocks binary path (default: agentlocks)
#
# Usage:
#   source "$(dirname "${BASH_SOURCE[0]}")/lib/task_lib.sh"
#   task_find_file <task_id>            -> echoes path, exits 1 if not found
#   task_read_field <file> <field>      -> echoes field value from frontmatter
#   task_write_field <file> <field> <val>
#   task_lock_acquire <file> <reason>   -> echoes lock_id, exits 1 on conflict
#   task_lock_release <lock_id>
#   task_lock_refresh <lock_id>
#   task_current_owner <file>           -> echoes agentlocks owner of the file's lock, or empty
#   task_now_iso                        -> echoes ISO 8601 timestamp
#   task_agent_id                       -> echoes current AGENTLOCKS_AGENT_ID

set -euo pipefail

AGENTLOCKS_BIN="${AGENTLOCKS_BIN:-agentlocks}"
SHARED_TASKS_DIR="${SHARED_TASKS_DIR:-.shared/tasks}"

task_agent_id() {
  if [[ -n "${AGENTLOCKS_AGENT_ID:-}" ]]; then
    echo "$AGENTLOCKS_AGENT_ID"
  elif [[ -n "${USER:-}" ]]; then
    echo "$USER"
  else
    echo "dev"
  fi
}

task_now_iso() {
  date -u +"%Y-%m-%dT%H:%M:%SZ"
}

# Find a task file by id (searches all subdirs of .shared/tasks/). Exits 1 if not found.
task_find_file() {
  local id="$1"
  local found=""
  # Try exact filename match across subdirectories
  while IFS= read -r -d '' f; do
    local base
    base="$(basename "$f" .md)"
    if [[ "$base" == "$id" ]]; then
      found="$f"
      break
    fi
  done < <(find "$SHARED_TASKS_DIR" -type f -name '*.md' ! -name 'README.md' -print0 2>/dev/null)
  if [[ -z "$found" ]]; then
    echo "Error: task '$id' not found under $SHARED_TASKS_DIR/" >&2
    return 1
  fi
  echo "$found"
}

# Read a single YAML frontmatter field from a task file. Echoes value (may be empty/quoted).
task_read_field() {
  local file="$1" field="$2"
  # Extract frontmatter (between first and second ---), then match `field: value`
  awk -v f="$field" '
    /^---$/ { fm++; next }
    fm==1 {
      if (index($0, f ":") == 1) {
        val = substr($0, length(f) + 2)
        sub(/^[ \t]+/, "", val)
        gsub(/^["'"'"']|["'"'"']$/, "", val)
        print val
        exit
      }
    }
    fm>=2 { exit }
  ' "$file"
}

# Write/update a single YAML frontmatter field in place. Creates the field if missing.
task_write_field() {
  local file="$1" field="$2" value="$3"
  local tmp
  tmp="$(mktemp)"
  # Quote value if it contains spaces or special chars; empty -> '""'
  local quoted
  if [[ -z "$value" ]]; then
    quoted='""'
  elif [[ "$value" =~ [[:space:]] || "$value" =~ [[:punct:]] ]]; then
    quoted="\"$value\""
  else
    quoted="$value"
  fi
  awk -v f="$field" -v v="$quoted" '
    BEGIN { in_fm=0; replaced=0 }
    /^---$/ {
      in_fm++
      if (in_fm==2 && !replaced) {
        print f ": " v
        replaced=1
      }
      print
      next
    }
    {
      if (in_fm==1 && index($0, f ":") == 1) {
        if (!replaced) { print f ": " v; replaced=1 }
      } else {
        print
      }
    }
  ' "$file" > "$tmp" && mv "$tmp" "$file"
}

# Acquire an agentlocks lock on a task file. Echoes lock_id on success.
# Exits non-zero on conflict or agentlocks failure.
# Note: agentlocks prints node module warnings to stderr; we discard stderr and
# only capture stdout (the clean lock id). On failure we re-run capturing stderr
# so the caller can surface the real error.
task_lock_acquire() {
  local file="$1" reason="$2"
  local id
  if ! id="$($AGENTLOCKS_BIN acquire "$file" --reason "$reason" --id-only 2>/dev/null)"; then
    local err
    err="$($AGENTLOCKS_BIN acquire "$file" --reason "$reason" --id-only 2>&1 >/dev/null)"
    echo "Error: agentlocks acquire failed:" >&2
    echo "$err" >&2
    return 1
  fi
  echo "$id" | tr -d '[:space:]'
}

# Release a lock by id. Silent on success.
task_lock_release() {
  local lock_id="$1"
  $AGENTLOCKS_BIN release "$lock_id" >/dev/null 2>&1 || true
}

# Refresh a lock's TTL. Silent on success.
task_lock_refresh() {
  local lock_id="$1"
  $AGENTLOCKS_BIN refresh "$lock_id" >/dev/null 2>&1 || true
}

# Echoes the agentlocks owner of any active lock covering `file`, or empty if none.
# Parses status --json via node (guaranteed available since agentlocks needs Node >= 22.18).
task_current_owner() {
  local file="$1"
  local json
  json="$($AGENTLOCKS_BIN status --json 2>/dev/null)" || return 0
  [[ -z "$json" ]] && return 0
  echo "$json" | node -e '
    let input = "";
    process.stdin.on("data", (c) => (input += c));
    process.stdin.on("end", () => {
      try {
        const data = JSON.parse(input);
        const file = process.argv[1];
        const lock = (data.locks || []).find((l) =>
          (l.resources || []).some((r) => r.value === file)
        );
        process.stdout.write(lock && lock.owner ? (lock.owner.agent_id || "") : "");
      } catch (e) {
        process.exit(0);
      }
    });
  ' "$file"
}

# Echoes the lock_id of any active lock covering `file`, or empty if none.
task_current_lock_id() {
  local file="$1"
  local json
  json="$($AGENTLOCKS_BIN status --json 2>/dev/null)" || return 0
  [[ -z "$json" ]] && return 0
  echo "$json" | node -e '
    let input = "";
    process.stdin.on("data", (c) => (input += c));
    process.stdin.on("end", () => {
      try {
        const data = JSON.parse(input);
        const file = process.argv[1];
        const lock = (data.locks || []).find((l) =>
          (l.resources || []).some((r) => r.value === file)
        );
        process.stdout.write(lock ? (lock.lock_id || "") : "");
      } catch (e) {
        process.exit(0);
      }
    });
  ' "$file"
}
