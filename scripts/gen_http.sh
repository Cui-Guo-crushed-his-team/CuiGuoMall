#!/usr/bin/env bash
set -euo pipefail
shopt -s globstar

ROLE=${1:-}
if [[ -z "$ROLE" || ! "$ROLE" =~ ^(client|server)$ ]]; then
  echo "usage: $0 <client|server>"
  exit 255
fi

if ! [[ "$0" =~ scripts/gen_http.sh ]]; then
  echo "must be run from root dir"
  exit 255
fi

PROTO_ROOT="./idl/proto/http"
proto_files=()

function find_http_proto_files() {
    while IFS= read -r file; do
      filename=$(basename "$file")
      if [[ "$filename" == "api.proto" || "$filename" == "common.proto" ]]; then
        continue
      fi
      filename=$(basename "$file" .proto)
      proto_files+=("$filename")
    done < <(find "$PROTO_ROOT" -type f -name "*.proto" | sort -u)
}

function gen_http_client() {
  local hertz_gen="./client"
  [ -d "$hertz_gen" ] && rm -rf "${hertz_gen:?}"

  for proto in "${proto_files[@]}"; do
    mkdir -p "$hertz_gen/http/$proto"
    echo "generating HTTP client code for ${proto}.proto"
    pushd http >/dev/null || exit 1
    cwgo client \
      --type HTTP \
      --service "client" \
      --module "github.com/Cui-Guo-crushed-his-team/CuiGuoMall/client" \
      --I ../idl/proto \
      --idl "../idl/proto/http/${proto}.proto"
    popd >/dev/null
  done
}

function gen_http_server() {
  local workdir="app/gateway"
  [ ! -d "$workdir" ] && mkdir -p "$workdir"

  for proto in "${proto_files[@]}"; do
    echo "generating HTTP server code for ${proto}.proto"
    pushd "$workdir" >/dev/null || exit 1
    cwgo server \
      --type HTTP \
      --service "gateway" \
      --module "github.com/Cui-Guo-crushed-his-team/CuiGuoMall/app/gateway" \
      --I ../../idl/proto \
      --idl "../../idl/proto/http/${proto}.proto"
    popd >/dev/null
  done
}

find_http_proto_files
echo "found HTTP proto files: ${proto_files[*]}"

case "$ROLE" in
  client) gen_http_client ;;
  server) gen_http_server ;;
esac