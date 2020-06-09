#!/bin/sh

protoc --proto_path=. --go_out=plugins=grpc:go --go_opt=paths=source_relative proto/CommonModels.proto proto/DatabaseModels.proto proto/LoadModels.proto proto/APIs.proto
mv go/proto/APIs.pb.go go/api
mv go/proto/DatabaseModels.pb.go go/databasemodels
mv go/proto/LoadModels.pb.go go/loadmodels
mv go/proto/CommonModels.pb.go go/commonmodels

protoc -I. --grpc-gateway_out=logtostderr=true,paths=source_relative:./go/gateway proto/APIs.proto

rm -r go/proto