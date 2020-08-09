# Generated by the protocol buffer compiler.  DO NOT EDIT!
# Source: proto/LoadAPI.proto for package ''

require 'grpc'
require 'proto/LoadAPI_pb'

module LoadService
  # This service is used to create datasets and load their entities into an object storage
  # and create corresponding metadata objects to store the location of the entities in object storage.
  # This interface should be usable from within the platform and from outside
  class Service

    include GRPC::GenericService

    self.marshal_class_method = :encode
    self.unmarshal_class_method = :decode
    self.service_name = 'LoadService'

    # Creates a list of upload links to place dataset entities in object storage
    # and adds corresponding metadata objects
    rpc :CreateUploadLink, CreateUploadLinkRequest, CreateUploadLinkResponse
    rpc :InitMultipartUpload, InitMultipartUploadRequest, InitMultiPartUploadResponse
    rpc :GetMultipartUploadLinkPart, GetMultipartUploadLinkPartRequest, GetMultipartUploadLinkPartResponse
    rpc :FinishMultipartUpload, FinishMultipartUploadRequest, Empty
  end

  Stub = Service.rpc_stub_class
end
