# summingservice

-- initial set for dev environment:
-- install the protoc compiler: brew install protobuf
-- ensure dep is installed: brew install dep
-- under the root of project: dep init
   Thiswill create a folder called “vendor” along with “Gopkg.lock” and “Gopkg.toml” file. 
   These two files are important to manage different dependencies of our project.

-- Install Go bindings & Generate Stubs
--This command will download go bindings into the “vendor” folder: dep ensure --add google.golang.org/grpc github.com/golang/protobuf/protoc-gen-go

-- Install the protoc plugin for Go
go get -u google.golang.org/grpc
go get -u github.com/golang/protobuf/proto
go get -u github.com/golang/protobuf/protoc-gen-go

-- Project folder structure:

build/  
-- contains sh script for generating new service go files from proto files whenever they are modified.

cmd/
    client/  
    -- configure a gRPC client to consume the services from remote host over protobuf
        cli/  
        -- a command line interface for client's use
    server/  
    -- configure a gRPC server and start it up to serve the services over protobuf

datastore/  
-- the data storeage to hold the data state by the service and for the service. currently, only an inmomory store is implemented for testing purpose. the real data store should be implemented as a backing service to the gRPC server.

gRPC/       
    proto-file/  
    domain/  
    service/  
    impl/

-- everything in here is internal to the services bounded only to this summing context
-- contains proto files for describing services and domain objects for this microservice.
-- protoc generated domain objects go file
-- protoc generated service go file
-- expose service endpoints, containing all business logics for the implementation of the services

vendor/      
Gopkg.lock
Gopkg.toml
-- outcome from executing dep init and dep ensure commands.

Dockerfile    
k8s-deployment.yml
-- define the docker image for this microservice in the CI/CD pipeline
-- k8 descriptions for deploy and create service into kubernettes 

Makefile
-- a handy script for bring all steps together, test, build, pack, deploy. this can be build into CI/CD pipline in the cloud


# Run the gRPC server and then client app for testing the flow on local dev box:

From your project’s root directory:

To run the gRPC server, execute:

$ export SERVICE_TCP_PORT=9000
$ go run ./cmd/gRPC/server/main.go

To run client (in a separate terminal window), execute:

$ export GRPC_SERVER_HOST=localhost:9000
$ go run ./cmd/gRPC/client/main.go


# Following 12factor app best practices
1. built on one codebase following TDD style, and is full tracked from: https://github.com/ydsxiong/summingservice.git. 

2. Dep is being used for dependencies management for developing and building the microservice

3. configs such as service port, etc. required for running the app are all stored in the environment

4. data source can be implemented as an external backing service according to store interface inside the datastore package. 

5. all build, deploy and make script files are ready for CI/CD pipeline integration to automate the process from source code into production.

6. no state is maintained inside this microservice, data/state communication is done via external backing service, like a db

7. required port binding to the service is done by build and deployment scripts

8. this sum service is small and single-purposed, and stateless, that would support auto-scaling in production

9. once deployed, this service is designed to be started and stopped fast with no side-effects left behind.

10. this service is containerized for a smooth CI/CD pipeline deployment.

11. the service handle is backed by middlewares for logging, instrumenting, etc. and the app is logged wherever applicable.

#  the service extensibility 
This service is designed to communicate with external data store via a designated interface: datastore/store.go.
And this interface can be implemented in various ways without impacting the service internal definitions and behaviours, e.g.
a backing database implementation to handle the data persistence, or a backing event store implementation for handling events or publishing further messages from those events, all can be easily added into the datastore package with no change required at all for the service internal core package gRPC.


# service accessibility
When deployed into a internal network, this service can be mapped to a public port and IP address by the host environment like kubernetes in the cloud.

