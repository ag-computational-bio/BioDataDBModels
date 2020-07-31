# Generated by the protocol buffer compiler.  DO NOT EDIT!
# Source: proto/ProjectAPI.proto for package ''

require 'grpc'
require 'proto/ProjectAPI_pb'

module ProjectAPI
  class Service

    include GRPC::GenericService

    self.marshal_class_method = :encode
    self.unmarshal_class_method = :decode
    self.service_name = 'ProjectAPI'

    rpc :CreateProject, CreateProjectRequest, Project
    rpc :AddUserToProject, AddUserToProjectRequest, Project
    rpc :GetProjectDatasets, ID, DatasetList
    rpc :GetUserProjects, Empty, ProjectEntryList
  end

  Stub = Service.rpc_stub_class
end
