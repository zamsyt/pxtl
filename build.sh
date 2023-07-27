#!/usr/bin/env bash
cd cmd/pxtl
for os in "darwin" "linux" "windows"; do for arch in "amd64" "arm64"; do
    if [[ $os == "windows" ]]; then ext=".exe"; fi
    env GOOS=$os GOARCH=$arch go build -o "../../bin/pxtl_${os}_${arch}${ext}"
done; done