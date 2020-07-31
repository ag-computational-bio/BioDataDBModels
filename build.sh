#!/bin/sh


# Builds the go stubs from the protobuf files and places them in their directories
# Temporary solution until ci/cd is enabled 
protoc --proto_path=. --go_out=plugins=grpc:go --go_opt=paths=source_relative proto/*.proto
mv go/proto/APIs.pb.go go/api
mv go/proto/ProjectAPI.pb.go go/api

mv go/proto/ProjectModels.pb.go go/projectmodels
mv go/proto/ProjectEntryModels.pb.go go/projectmodels


mv go/proto/DatasetModels.pb.go go/datasetmodels
mv go/proto/LoadModels.pb.go go/loadmodels
mv go/proto/CommonModels.pb.go go/commonmodels
mv go/proto/MetadataModels.pb.go go/metadatamodels
mv go/proto/AuthModels.pb.go go/authmodels
mv go/proto/DatasetEntryModels.pb.go go/datasetentrymodels


protoc -I. --grpc-gateway_out=logtostderr=true,paths=source_relative:./go/api proto/APIs.proto proto/*API.proto
mv ./go/api/proto/APIs.pb.gw.go ./go/api
mv ./go/api/proto/ProjectAPI.pb.gw.go ./go/api
rm -r ./go/api/proto

protoc -I. --swagger_out=logtostderr=true,allow_merge=true:./go/swaggerhandler proto/APIs.proto proto/*API.proto

go generate go/swaggerhandler/SwaggerGen.go

grpc_tools_ruby_protoc -I. --ruby_out=ruby --grpc_out=ruby proto/*.proto
mv ruby/proto/* ruby/lib

rm -r ruby/proto
rm -r go/proto