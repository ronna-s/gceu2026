#!/usr/bin/env bash

set -euo pipefail

SCRIPT_DIR="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd)"
IMAGE_NAME="gceu-l3-performance-tools"
BINARY_PATH="./mutex-contention.test"

usage() {
  cat <<'EOF'
Usage: ./pprof.sh <cpu|mutex> [port]

Examples:
  ./pprof.sh cpu
  ./pprof.sh mutex
  ./pprof.sh cpu 9090
EOF
}

if [[ $# -lt 1 || $# -gt 2 ]]; then
  usage
  exit 1
fi

PROFILE_KIND="$1"
PORT="${2:-8080}"

case "$PROFILE_KIND" in
  cpu)
    PROFILE_PATH="./cpu.out"
    ;;
  mutex)
    PROFILE_PATH="./mutex.out"
    ;;
  *)
    echo "Unknown profile: $PROFILE_KIND" >&2
    usage
    exit 1
    ;;
esac

if [[ ! -f "$SCRIPT_DIR/${BINARY_PATH#./}" ]]; then
  echo "Missing test binary: $SCRIPT_DIR/${BINARY_PATH#./}" >&2
  exit 1
fi

if [[ ! -f "$SCRIPT_DIR/${PROFILE_PATH#./}" ]]; then
  echo "Missing profile: $SCRIPT_DIR/${PROFILE_PATH#./}" >&2
  exit 1
fi

docker build -t "$IMAGE_NAME" -f "$SCRIPT_DIR/Dockerfile.tools" "$SCRIPT_DIR"

docker run --rm -it \
  -p "$PORT:8080" \
  -v "$SCRIPT_DIR:/work" \
  -w /work \
  "$IMAGE_NAME" \
  go tool pprof -http=0.0.0.0:8080 -no_browser "$BINARY_PATH" "$PROFILE_PATH"
