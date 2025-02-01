#!/usr/bin/env bash
set -euo pipefail
shopt -s globstar

if ! [[ "$0" =~ scripts/gen_rpc_client.sh ]]; then
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
function gen_rpc_client_code() {
  local kitex_gen="./rpc/kitex_gen"
  if [ -d "$kitex_gen" ]; then
    echo "found existing $kitex_gen, cleaning all files under it"
    rm -rf ${kitex_gen:?}
  fi

  for proto in "${proto_files[@]}"; do
    local proto_file="${proto}.proto"
    echo "service=$proto_file"

    if [ -d "$kitex_gen/rpc/$proto" ]; then
        rm -rf "${kitex_gen:?}"/"${proto:?}"/*
    else
      mkdir -p "$kitex_gen/rpc/$proto"
    fi
    echo "generating code for service ${proto} to $kitex_gen/rpc/$proto"
    pushd rpc > /dev/null || exit 1
    cwgo client \
      --type RPC \
      --service "${proto}" \
      --module github.com/Cui-Guo-crushed-his-team/CuiGuoMall/rpc \
      --I ../idl/proto/rpc \
      --idl ../idl/proto/rpc/"${proto_file}"
    if [ -d "$kitex_gen/rpc/$proto" ]; then
      echo "gen ${proto} service rpc code finish"
    fi
    popd > /dev/null
  done
}
find_proto_files
echo "proto files has been found: ${proto_files[*]}"
gen_rpc_client_code
