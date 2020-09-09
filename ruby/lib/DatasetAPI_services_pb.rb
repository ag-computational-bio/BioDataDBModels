# Generated by the protocol buffer compiler.  DO NOT EDIT!
# Source: proto/DatasetAPI.proto for package ''

require 'grpc'
require 'proto/DatasetAPI_pb'

module DatasetService
  # Dataset management service
  # Manages all dataset related services
  # All data objects are associated with one data dataset
  # Dataset versions group these data objects, which makes them reusable
  class Service

    include GRPC::GenericService

    self.marshal_class_method = :encode
    self.unmarshal_class_method = :decode
    self.service_name = 'DatasetService'

    # CreateNewDataset Creates a new dataset and associates it with a dataset
    rpc :CreateNewDataset, ::CreateDatasetRequest, ::DatasetEntry
    # Datasets Lists all datasets of a user
    rpc :Datasets, ::ID, ::DatasetList
    # Lists Versions of a dataset
    rpc :DatasetVersions, ::ID, ::DatasetVersionList
    # Updates a field of a dataset
    rpc :UpdateDatasetField, ::UpdateFieldsRequest, ::DatasetEntry
    # Updates the current dataset version of a dataset
    rpc :UpdateCurrentDatasetVersion, ::UpdateCurrentDatasetVersionRequest, ::DatasetEntry
    # DeleteDataset Delete a dataset
    # Datasets can only be deleted if 
    rpc :DeleteDataset, ::ID, ::Empty
    # ---------------------------------------------------------------------------------------
    # Dataset version calls
    #
    # Creates a new dataset version which is linked to an exisiting dataset
    rpc :CreateNewDatasetVersion, ::CreateDatasetVersionRequest, ::DatasetVersionEntry
    rpc :GetDatasetVersion, ::ID, ::DatasetVersionEntry
    rpc :UpdateDatasetVersionField, ::UpdateFieldsRequest, ::DatasetEntry
    rpc :AddObjectGroupToDatasetVersion, ::AddDatasetObjectGroupToDatasetVersionRequest, ::Empty
    # Deletes a dataset version
    # This should not delete the underlaying dataset objects
    rpc :DeleteDatasetVersion, ::ID, ::Empty
    # DatasetVersionObjectGroups Lists all objects groups that are part of the given dataset version
    rpc :DatasetVersionObjectGroups, ::ID, ::DatasetObjectGroupList
  end

  Stub = Service.rpc_stub_class
end
# Dataset calls
module ObjectsService
  class Service

    include GRPC::GenericService

    self.marshal_class_method = :encode
    self.unmarshal_class_method = :decode
    self.service_name = 'ObjectsService'

    # CreateDatsetObjectGroup Creates a new dataset object group in the database
    # Will also create all related dataset objects
    rpc :CreateDatsetObjectGroup, ::CreateDatasetObjectGroupRequest, ::DatasetObjectGroup
    # DatasetVersionObjectGroups Lists all objects groups that are part of the given dataset version
    rpc :DatasetVersionObjectGroups, ::ID, ::DatasetObjectGroupList
    # GetDatasetObjectGroup The dataset object group with the given ID
    # Will only return a dataset object group without its affiliated objects
    rpc :GetDatasetObjectGroup, ::ID, ::DatasetObjectGroup
    # DeleteDatasetObjectGroup Deletes the given dataset group and all associated dataset objects
    # Can only be used when all linked dataset versions have been deleted
    rpc :DeleteDatasetObjectGroup, ::ID, ::Empty
  end

  Stub = Service.rpc_stub_class
end
# Object group calls
