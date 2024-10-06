#!/bin/bash

protoc --proto_path=../proto --go_out=../proto --go_opt=paths=source_relative ../proto/blur.proto ../proto/detection.proto
protoc --proto_path=../proto --go-grpc_out=../proto --go-grpc_opt=paths=source_relative ../proto/blur.proto