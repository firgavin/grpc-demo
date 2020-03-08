#!/bin/bash

parent_dir="$(go env GOPATH)"/src
echo "$parent_dir"
proto_path=(
  "proto"
)

for entry in "${proto_path[@]}"; do
  protoc -I "$parent_dir" "$entry"/*.proto --proto_path="$entry" --go_out=plugins=grpc:"$parent_dir"
done
