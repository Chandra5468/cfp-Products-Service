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
    protoc --proto_path=pkg/protobuf pkg/protobuf/product.proto --go_out=internal/services/common/genproto/products --go_opt=paths=source_relative --go-grpc_out=internal/services/common/genproto/products --go-grpc_opt=paths=source_relative