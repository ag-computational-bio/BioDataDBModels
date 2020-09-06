# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/DatasetAPIModels.proto

require 'google/protobuf'

require 'google/protobuf/struct_pb'
require 'proto/CommonModels_pb'
require 'proto/DatasetEntryModels_pb'
require 'protoc/gateway/options/annotations_pb'
require 'google/protobuf/timestamp_pb'
Google::Protobuf::DescriptorPool.generated_pool.build do
  add_file("proto/DatasetAPIModels.proto", :syntax => :proto3) do
    add_message "CreateDatasetRequest" do
      optional :DatasetName, :string, 1
      optional :Datatype, :string, 2
      optional :ProjectID, :string, 3
    end
    add_message "DatasetList" do
      repeated :Datasets, :message, 1, "DatasetEntry"
    end
    add_message "CreateDatasetVersionRequest" do
      optional :DatasetID, :string, 1
      optional :Version, :message, 2, "Version"
      optional :AdditionalMetadataMessageRef, :string, 3
      optional :AdditionalObjectMetadataMessageRef, :string, 4
      optional :AdditionalMetadata, :message, 5, "google.protobuf.Struct"
      optional :ExpectedObjectCount, :int64, 6
    end
    add_message "DatasetVersionList" do
      repeated :DatasetVersions, :message, 2, "DatasetVersionEntry"
    end
    add_message "UpdateDatasetVersionObjectCountRequest" do
      optional :DatasetVersionID, :string, 1
      optional :Value, :int64, 2
      optional :DatasetID, :string, 3
    end
    add_message "CreateDatasetObjectGroupRequest" do
      optional :DatasetObjectAnchor, :string, 1
      optional :Name, :string, 2
      optional :Version, :message, 3, "Version"
      optional :DatasetID, :string, 4
      repeated :DatasetVersionID, :string, 5
      repeated :DatasetObjects, :message, 6, "CreateDatasetObjectRequest"
    end
    add_message "CreateDatasetObjectRequest" do
      optional :Filename, :string, 1
      optional :Filetype, :string, 2
      optional :Origin, :message, 3, "Origin"
      optional :Created, :message, 4, "google.protobuf.Timestamp"
      optional :AdditionalMetadata, :message, 5, "google.protobuf.Struct"
      optional :ContentLen, :int64, 8
      optional :UploadID, :string, 9
    end
    add_message "DatasetObjectGroupList" do
      repeated :DatasetObjectGroups, :message, 1, "DatasetObjectGroup"
    end
    add_message "DatasetObjectLinks" do
      optional :ID, :string, 1
      repeated :Entites, :message, 2, "DatasetObjectLink"
    end
    add_message "DatasetObjectLink" do
      optional :Link, :string, 1
      optional :ObjectID, :string, 2
      optional :Object, :message, 3, "DatasetObjectEntry"
    end
    add_message "UpdateFieldsRequest" do
      optional :ID, :string, 1
      map :UpdateStringFields, :string, :string, 2
      map :UpdateInt64Fields, :string, :int64, 3
    end
    add_message "UpdateCurrentDatasetVersionRequest" do
      optional :ID, :string, 1
      optional :UpdateTargetID, :string, 2
      optional :TargetResource, :enum, 3, "Resource"
      optional :UpdateStage, :enum, 4, "Stage"
    end
  end
end

CreateDatasetRequest = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("CreateDatasetRequest").msgclass
DatasetList = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("DatasetList").msgclass
CreateDatasetVersionRequest = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("CreateDatasetVersionRequest").msgclass
DatasetVersionList = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("DatasetVersionList").msgclass
UpdateDatasetVersionObjectCountRequest = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("UpdateDatasetVersionObjectCountRequest").msgclass
CreateDatasetObjectGroupRequest = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("CreateDatasetObjectGroupRequest").msgclass
CreateDatasetObjectRequest = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("CreateDatasetObjectRequest").msgclass
DatasetObjectGroupList = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("DatasetObjectGroupList").msgclass
DatasetObjectLinks = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("DatasetObjectLinks").msgclass
DatasetObjectLink = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("DatasetObjectLink").msgclass
UpdateFieldsRequest = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("UpdateFieldsRequest").msgclass
UpdateCurrentDatasetVersionRequest = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("UpdateCurrentDatasetVersionRequest").msgclass