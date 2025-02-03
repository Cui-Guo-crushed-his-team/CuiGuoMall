#!/usr/bin/env bash
set -euo pipefail
shopt -s globstar

ROLE=${1:-}  # client or server
if [[ -z "$ROLE" || ! "$ROLE" =~ ^(client|server)$ ]]; then
  echo "usage: $0 <client|server>"
  exit 255
fi

if ! [[ "$0" =~ scripts/gen_rpc.sh ]]; then
  echo "must be run from root dir"
  exit 255
fi

PROTO_ROOT="./idl/proto/rpc"
proto_files=()

function find_proto_files() {
    while IFS= read -r file; do
      filename=$(basename "$file" .proto)
      proto_files+=("$filename")
    done < <(find "$PROTO_ROOT" -type f -name "*.proto" | sort -u)
}

function gen_client() {
  local kitex_gen="./rpc/kitex_gen"
  [ -d "$kitex_gen" ] && rm -rf "${kitex_gen:?}"

  for proto in "${proto_files[@]}"; do
    mkdir -p "$kitex_gen/rpc/$proto"
    echo "generating client code for ${proto}.proto"
    pushd rpc >/dev/null || exit 1
    cwgo client \
      --type RPC \
      --service "${proto}" \
      --module github.com/Cui-Guo-crushed-his-team/CuiGuoMall/rpc \
      --I ../idl/proto/rpc \
      --idl "../idl/proto/rpc/${proto}.proto"
    popd >/dev/null
  done
}

function gen_server() {
  local kitex_gen="./rpc/kitex_gen"
  [ ! -d "$kitex_gen" ] && { echo "Run client generation first"; exit 1; }

  for proto in "${proto_files[@]}"; do
    local workdir="app/${proto}"
    echo "generating server code for ${proto}.proto"
    pushd "$workdir" >/dev/null || exit 1
    cwgo server \
      --type RPC \
      --service "${proto}" \
      --pass "-use github.com/Cui-Guo-crushed-his-team/CuiGuoMall/rpc/kitex_gen" \
      --module "github.com/Cui-Guo-crushed-his-team/CuiGuoMall/app/${proto}" \
      --I ../../idl/proto/rpc \
      --idl "../../idl/proto/rpc/${proto}.proto"
    popd >/dev/null
  done
}

find_proto_files
echo "found proto files: ${proto_files[*]}"

case "$ROLE" in
  client) gen_client ;;
  server) gen_server ;;
esac