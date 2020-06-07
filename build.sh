#!/bin/sh

protoc --proto_path=proto --go_out=plugins=grpc:go --go_opt=paths=source_relative proto/ApiModel.proto