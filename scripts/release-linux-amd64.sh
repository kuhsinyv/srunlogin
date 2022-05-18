#!/usr/bin/env bash

DIR=releases/srunlogin-linux-amd64-"$1"
EXEC=srunlogin-linux-amd64-"$1".exe

echo "Building ${EXEC}"

if [ ! -d DIR ]; then
  mkdir DIR
fi

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
go build \
-o "${DIR}"/"${EXEC}" \
-ldflags="-s -w" \
cmd/srunlogin/srunlogin.go

upx -9 "${DIR}"/"${EXEC}"

cp config.yaml "${DIR}"/

zip -r "${DIR}".zip "${DIR}"

rm -r "${DIR}"