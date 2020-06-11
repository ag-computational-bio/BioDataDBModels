#!/bin/sh


# Builds the go stubs from the protobuf files and places them in their directories
# Temporary solution until ci/cd is enabled 
protoc --proto_path=. --go_out=plugins=grpc:go --go_opt=paths=source_relative proto/CommonModels.proto proto/DatabaseModels.proto proto/LoadModels.proto proto/APIs.proto
mv go/proto/APIs.pb.go go/api
mv go/proto/DatabaseModels.pb.go go/databasemodels
mv go/proto/LoadModels.pb.go go/loadmodels
mv go/proto/CommonModels.pb.go go/commonmodels

protoc -I. --grpc-gateway_out=logtostderr=true,paths=source_relative:./go/api proto/APIs.proto
mv ./go/api/proto/APIs.pb.gw.go ./go/api
rm -r ./go/api/proto

protoc -I. --swagger_out=logtostderr=true,allow_merge=true:./go/swaggerhandler proto/APIs.proto proto/Swagger.proto

go generate go/swaggerhandler/SwaggerGen.go

rm -r go/proto