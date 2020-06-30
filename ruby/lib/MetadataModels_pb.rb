# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/MetadataModels.proto

require 'google/protobuf'

require 'protoc/gateway/options/annotations_pb'
require 'google/protobuf/struct_pb'
require 'google/protobuf/timestamp_pb'
Google::Protobuf::DescriptorPool.generated_pool.build do
  add_file("proto/MetadataModels.proto", :syntax => :proto3) do
    add_message "InitMetadataDBRequest" do
      optional :Name, :string, 1
    end
    add_message "MetadataDBEntry" do
      optional :ID, :string, 1
      optional :Name, :string, 2
      optional :StructMessageRef, :string, 3
      optional :Created, :message, 4, "google.protobuf.Timestamp"
    end
    add_message "MetadataDBList" do
      repeated :entries, :message, 1, "MetadataDBEntry"
    end
    add_message "InitMetadataCollectionRequest" do
      optional :Name, :string, 1
      optional :MessageRef, :string, 2
      optional :MetadataDBEntryID, :string, 3
    end
    add_message "MetadataCollectionEntry" do
      optional :ID, :string, 1
      optional :Name, :string, 2
      optional :MessageRef, :string, 3
      optional :StructMessageRef, :string, 4
      optional :Created, :message, 5, "google.protobuf.Timestamp"
      optional :MetadatDBEntryID, :string, 6
    end
    add_message "MetadataDBCollectionList" do
      repeated :entries, :message, 1, "MetadataCollectionEntry"
    end
    add_message "InsertMetadataRequest" do
      optional :MetadataDBID, :string, 1
      optional :CollectionID, :string, 2
      repeated :StructMetadata, :message, 4, "StructMetadataMessage"
    end
    add_message "StructMetadataMessage" do
      optional :Metadata, :message, 1, "google.protobuf.Struct"
      map :Int64Indices, :string, :int64, 2
      map :UInt64Indices, :string, :uint64, 3
    end
    add_message "AddMetadataIndexRequest" do
      optional :MetadataDBID, :string, 1
      optional :CollectionID, :string, 2
      repeated :Indices, :message, 3, "Index"
    end
    add_message "QueryRequest" do
      optional :MetadataDBID, :string, 1
      optional :CollectionID, :string, 2
      optional :Query, :message, 4, "google.protobuf.Struct"
    end
    add_message "Index" do
      optional :Field, :string, 1
    end
    add_message "Field" do
      repeated :Data, :message, 1, "google.protobuf.Struct"
    end
  end
end

InitMetadataDBRequest = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("InitMetadataDBRequest").msgclass
MetadataDBEntry = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("MetadataDBEntry").msgclass
MetadataDBList = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("MetadataDBList").msgclass
InitMetadataCollectionRequest = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("InitMetadataCollectionRequest").msgclass
MetadataCollectionEntry = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("MetadataCollectionEntry").msgclass
MetadataDBCollectionList = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("MetadataDBCollectionList").msgclass
InsertMetadataRequest = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("InsertMetadataRequest").msgclass
StructMetadataMessage = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("StructMetadataMessage").msgclass
AddMetadataIndexRequest = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("AddMetadataIndexRequest").msgclass
QueryRequest = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("QueryRequest").msgclass
Index = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("Index").msgclass
Field = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("Field").msgclass
