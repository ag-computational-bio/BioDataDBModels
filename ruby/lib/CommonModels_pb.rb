# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/CommonModels.proto

require 'google/protobuf'

require 'protoc/gateway/options/annotations_pb'
Google::Protobuf::DescriptorPool.generated_pool.build do
  add_file("proto/CommonModels.proto", :syntax => :proto3) do
    add_message "Empty" do
    end
    add_message "Location" do
      optional :Bucket, :string, 1
      optional :Key, :string, 2
      optional :URL, :string, 3
    end
    add_message "Origin" do
      optional :Link, :string, 1
      optional :ObjectStorageLocatio, :message, 2, "Location"
      optional :OriginType, :enum, 4, "Origin.OriginTypeEnum"
    end
    add_enum "Origin.OriginTypeEnum" do
      value :ObjectStorage, 0
      value :OriginLink, 1
    end
    add_message "ID" do
      optional :ID, :string, 1
    end
  end
end

Empty = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("Empty").msgclass
Location = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("Location").msgclass
Origin = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("Origin").msgclass
Origin::OriginTypeEnum = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("Origin.OriginTypeEnum").enummodule
ID = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("ID").msgclass