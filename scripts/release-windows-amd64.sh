#!/usr/bin/env bash

DIR=releases/srunlogin-windows-amd64-"$1"
EXEC=srunlogin-windows-amd64-"$1".exe

if [ ! -d DIR ]; then
  mkdir DIR
fi

CGO_ENABLED=0 GOOS=windows GOARCH=amd64 \
go build \
-o "${DIR}"/"${EXEC}" \
-ldflags="-s -w" \
cmd/srunlogin/srunlogin.go

upx -9 "${DIR}"/"${EXEC}"

cp config.yaml "${DIR}"/

zip -r "${DIR}".zip "${DIR}"

rm -r "${DIR}"