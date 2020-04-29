#!/bin/bash -e

declare -a arr=("solution" "workticket")

cd protoc
PROTOCOL_DIR=protocol/
rm -rf ${PROTOCOL_DIR}
mkdir ${PROTOCOL_DIR}

for i in "${arr[@]}"
do
protoc --proto_path=. --go_out=plugins=grpc,paths=source_relative:${PROTOCOL_DIR} ${i}/*
done

echo $(pwd)
