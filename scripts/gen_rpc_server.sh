#!/usr/bin/env bash
set -euo pipefail
shopt -s globstar

if ! [[ "$0" =~ scripts/gen_rpc_server.sh ]]; then
  echo "must be run from root dir"
  exit 255
fi

PROTO_ROOT="./rpc/idl/proto/rpc"
proto_files=()

function find_proto_files() {
    while IFS= read -r file; do
      filename=$(basename "$file" .proto)
      proto_files+=("$filename")
    done < <(find "$PROTO_ROOT" -type f -name "*.proto" | sort -u)
}
function gen_rpc_server_code() {
  local kitex_gen="./rpc/kitex_gen"
  if [ ! -d "$kitex_gen" ]; then
    echo "make sure run command: make gen_rpc_client_code ,first"
    exit 255
  fi
  for proto in "${proto_files[@]}"; do
    local proto_file="${proto}.proto"
    echo "service=$proto_file"
    echo "generating server code for service ${proto} to app/$proto"
    local workdir=app/"${proto}"
    pushd "${workdir:?}" > /dev/null || exit 1
    pwd
    cwgo server \
      --type RPC \
      --service "${proto}" \
      --pass "-use github.com/Cui-Guo-crushed-his-team/CuiGuoMall/rpc/kitex_gen" \
      --module github.com/Cui-Guo-crushed-his-team/CuiGuoMall/app/"${proto}" \
      --I ../../rpc/idl/proto/rpc \
      --idl ../../rpc/idl/proto/rpc/"${proto_file}"
    popd > /dev/null
  done
}
find_proto_files
echo "proto files has been found: ${proto_files[*]}"
gen_rpc_server_code
