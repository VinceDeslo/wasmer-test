#!/bin/bash
cd ../cmd/wasm && GOOS=js GOARCH=wasm go build -o  ../../assets/main.wasm
