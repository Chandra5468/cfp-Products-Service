# cfp-Products-Service
Products microservice under cfp(cakefactoryproject). Golang and Psql are used here.

# PSQL Table schema used 

CREATE TABLE products (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL,
    description TEXT,
    price DECIMAL(10, 2) NOT NULL,
    quantity INTEGER NOT NULL CHECK (quantity >= 0),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);



## GRPC TRIALS :
1. Proto file trial 1 :
        syntax = "proto3";

        option go_package = "github.com/Chandra5468/cfp-Products-Service/protobuf";

        service ProductService {
            rpc GetProduct (GetProductRequest) returns (GetProductResponse) {}
        }

2.  syntax = "proto3";

    package "github.com/Chandra5468/cfp-Products-Service/protobuf";

    service ProductService {
        rpc GetProduct (GetProductRequest) returns (GetProductResponse) {}
    }

    message GetProductRequest {
        bytes id = 1
        string name = 2
    }

3. GRPC go proto buf generation files
    protoc --proto_path=pkg/genproto/products pkg/genproto/products/product.proto --go_out=pkg/genproto/products --go_opt=paths=source_relative --go-grpc_out=pkg/genproto/products --go-grpc_opt=paths=source_relative

4. The pb.go file contains the Go structs for the Protobuf messages defined in your .proto file. 
        These are used by  both the client and server to serialize and deserialize data.
   The _grpc.pb.go file contains the gRPC service interface (for the server to implement) and the client stubs (for the client to call the service).

5. Folder structure
    // Folder structure
    cfp-Products-Service/
    ├── internal/
    │   ├── services/
    │   │   ├── common/
    │   │   │   ├── genproto/
    │   │   │   │   ├── products/
    │   │   │   │   │   ├── products.proto
    │   │   │   │   │   ├── products.pb.go
    │   │   │   │   │   ├── products_grpc.pb.go
    │   │   │   │   │   ├── go.mod
    │   │   │   │   │   ├── go.sum
    ├── server/
    │   ├── main.go
    │   ├── go.mod
    │   ├── go.sum
    ├── client/
    │   ├── main.go
    │   ├── go.mod
    │   ├── go.sum