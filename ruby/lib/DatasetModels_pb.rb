# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/DatasetModels.proto

require 'google/protobuf'

require 'google/protobuf/timestamp_pb'
require 'proto/CommonModels_pb'
require 'protoc/gateway/options/annotations_pb'
Google::Protobuf::DescriptorPool.generated_pool.build do
  add_file("proto/DatasetModels.proto", :syntax => :proto3) do
    add_message "DatasetEntry" do
      optional :ID, :string, 1
      optional :Metadata, :message, 2, "DatasetEntryMetadata"
    end
    add_message "DatasetEntryMetadata" do
      optional :Datasetname, :string, 1
      optional :Datasettype, :string, 2
      optional :OwnerID, :string, 3
      optional :isPublic, :bool, 4
    end
    add_message "DatasetObjectEntry" do
      optional :ID, :string, 1
      optional :DatasetID, :string, 2
      optional :Metadata, :message, 3, "DatasetObjectMetaData"
      optional :Location, :message, 4, "Location"
    end
    add_message "DatasetObjectMetaData" do
      optional :Filename, :string, 1
      optional :Filetype, :string, 2
      optional :Name, :string, 3
      optional :Version, :string, 4
      optional :Origin, :message, 5, "Origin"
      optional :ContentLen, :int64, 6
    end
    add_message "DatasetObjectLinks" do
      optional :ID, :string, 1
      repeated :Entites, :message, 2, "DatasetObjectLink"
    end
    add_message "DatasetObjectLink" do
      optional :Link, :string, 1
      optional :ObjectID, :string, 2
    end
    add_message "CreateDatasetRequest" do
      optional :DatasetName, :string, 1
      optional :Datatype, :string, 2
    end
    add_message "DatasetList" do
      repeated :Datasets, :message, 1, "DatasetEntry"
    end
    add_message "DatasetVersionList" do
      repeated :DatasetVersions, :message, 2, "DatasetVersionEntry"
    end
    add_message "DatasetObjectList" do
      repeated :Datasets, :message, 1, "DatasetEntry"
    end
    add_message "DatasetVersionEntry" do
      optional :ID, :string, 1
      optional :Metadata, :message, 2, "DatasetVersionMetadata"
    end
    add_message "DatasetVersionMetadata" do
      optional :DatasetID, :string, 1
      optional :DatasetName, :string, 2
      optional :Version, :string, 3
      optional :Created, :message, 4, "google.protobuf.Timestamp"
    end
    add_message "NewDatasetVersionRequest" do
      optional :DatasetID, :string, 1
      optional :Version, :string, 2
    end
  end
end

DatasetEntry = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("DatasetEntry").msgclass
DatasetEntryMetadata = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("DatasetEntryMetadata").msgclass
DatasetObjectEntry = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("DatasetObjectEntry").msgclass
DatasetObjectMetaData = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("DatasetObjectMetaData").msgclass
DatasetObjectLinks = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("DatasetObjectLinks").msgclass
DatasetObjectLink = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("DatasetObjectLink").msgclass
CreateDatasetRequest = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("CreateDatasetRequest").msgclass
DatasetList = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("DatasetList").msgclass
DatasetVersionList = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("DatasetVersionList").msgclass
DatasetObjectList = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("DatasetObjectList").msgclass
DatasetVersionEntry = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("DatasetVersionEntry").msgclass
DatasetVersionMetadata = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("DatasetVersionMetadata").msgclass
NewDatasetVersionRequest = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("NewDatasetVersionRequest").msgclass
