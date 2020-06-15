# Generated by the protocol buffer compiler.  DO NOT EDIT!
# Source: proto/APIs.proto for package ''

require 'grpc'
require 'proto/APIs_pb'

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
    rpc :CreateLoadLinkSet, CreateLoadLinkSetRequest, UploadLinks
  end

  Stub = Service.rpc_stub_class
end
module DatasetService
  # Dataset management service
  class Service

    include GRPC::GenericService

    self.marshal_class_method = :encode
    self.unmarshal_class_method = :decode
    self.service_name = 'DatasetService'

    # Creates a new dataset
    rpc :CreateNewDataset, CreateDatasetRequest, DatasetEntry
    # Creates a new dataset version based on an existing dataset
    rpc :CreateNewDatasetVersion, NewDatasetVersionRequest, DatasetVersionEntry
    # Lists all datasets
    rpc :Datasets, Empty, DatasetEntry
    # Lists Versions of a dataset
    rpc :DatasetVersion, ID, DatasetVersionEntry
    # Lists all entities of a dataset
    rpc :DatasetVersionEntities, ID, DatasetObjectList
    # Creates preauth download links for all entities of a dataset
    rpc :DatasetVersionEntitiesLinks, ID, DatasetObjectLinks
  end

  Stub = Service.rpc_stub_class
end
