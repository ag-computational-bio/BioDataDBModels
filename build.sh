#!/bin/sh

protoc --proto_path=. --go_out=plugins=grpc:go --go_opt=paths=source_relative proto/DatabaseModels.proto proto/LoadModels.proto proto/APIs.proto
mv go/proto/APIs.pb.go go/api
mv go/proto/DatabaseModels.pb.go go/database
mv go/proto/LoadModels.pb.go go/loadmodels

rm -r go/proto