#!/bin/sh

set -e

readonly COMMIT="$(git rev-parse HEAD)"
readonly DATE="$(date -u +"%Y-%m-%dT%H:%M:%SZ")"
readonly LDFLAGS="-X code.dwrz.net/src/pkg/build.Commit=${COMMIT} \
-X code.dwrz.net/src/pkg/build.Hostname=${HOSTNAME} \
-X code.dwrz.net/src/pkg/build.Time=${DATE}"

# Disable the race detector unless RACE is set to true in the environment.
RACE="${RACE:-false}"

all() {
  clean

  for f in cmd/*; do
    if [ -d "$f" ]; then
      cmd=$(basename "$f")

      build "$cmd"
    fi
  done
}

build() {
  local cmd="$1"

  go build \
     -ldflags "${LDFLAGS}" \
     -o bin/"${cmd}" \
     -race="${RACE}" \
     ./cmd/"${cmd}"/*.go
  if [[ "$?" != 0 ]]; then
    printf "\nFailed to build %s.\n" "${cmd}"
    exit 1
  fi
}

clean() {
  go clean
  rm -rf bin/*
}

gotest() {
  clean
  all
  go test ./...
}


lint() {
  goimports -e -w -local="code.dwrz.net/src" ./cmd/ ./pkg/
  go fmt ./...
  go vet ./...
  govulncheck ...
}

run() {
  local cmd="$1"

  build "${cmd}"

  ./bin/"${cmd}" "${@:2}"
}

main() {
  local action="$1";
  if [ ! -z "${action}" ]; then
    shift;
  fi

  case "${action}" in
    "all"|"") all ;;
    "build") build "$@" ;;
    "clean") clean ;;
    "lint") lint ;;
    "run") run "$@" ;;
    "test") gotest ;;
    *) all ;;
  esac
}

main "$@"
