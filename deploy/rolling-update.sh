#!/usr/bin/env bash
set -euo pipefail

if [[ $# -ne 1 || ! -f "$1" ]]; then
  echo "usage: $0 /path/to/new/Danmu" >&2
  exit 2
fi

new_binary=$(readlink -f "$1")

wait_healthy() {
  local port=$1
  local deadline=$((SECONDS + 90))
  while (( SECONDS < deadline )); do
    if curl -fsS --max-time 3 "http://127.0.0.1:${port}/healthz" | grep -Eq '"status":"ok".*"connected":true'; then
      return 0
    fi
    sleep 2
  done
  return 1
}

deploy_instance() {
  local instance=$1
  local port=$2
  local directory="/home/ecs-user/${instance}"
  local binary="${directory}/Danmu"
  local previous="${directory}/Danmu.previous"

  echo "Updating ${instance}..."
  sudo cp -p "$binary" "$previous"
  sudo install -o ecs-user -g ecs-user -m 0755 "$new_binary" "${binary}.next"
  sudo mv -f "${binary}.next" "$binary"
  sudo systemctl restart "danmu@${instance}.service"

  if wait_healthy "$port"; then
    echo "${instance} is healthy."
    return 0
  fi

  echo "${instance} failed health check; rolling it back." >&2
  sudo install -o ecs-user -g ecs-user -m 0755 "$previous" "${binary}.next"
  sudo mv -f "${binary}.next" "$binary"
  sudo systemctl restart "danmu@${instance}.service"
  wait_healthy "$port" || echo "WARNING: ${instance} is still unhealthy after rollback." >&2
  return 1
}

# Both instances record normally. During upgrades, update and verify the backup
# first so the main instance stays on the known-good version until verification.
deploy_instance DanmuBackup 10217
deploy_instance Danmu 10216
echo "Rolling update completed: backup first, then main."
