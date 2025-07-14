package grpcserver

import (
	"database/sql"
	"log"
	"net"

	grpchandler "github.com/Chandra5468/cfp-Products-Service/internal/handlers/grpcHandler"
	"github.com/Chandra5468/cfp-Products-Service/internal/services/database/postgresql/products"
	"google.golang.org/grpc"
)

type grpcServer struct {
	addr string
	db   *sql.DB
}

func NewGrpcServer(addr string, db *sql.DB) *grpcServer {
	return &grpcServer{
		addr: addr,
		db:   db,
	}
}

func (s *grpcServer) Run() {
	lis, err := net.Listen("tcp", s.addr)

	if err != nil {
		log.Fatalf("failed to listen on tcp for grpc : %s", err.Error())
	}

	grpcServer := grpc.NewServer()

	// register our grpc services
	store := products.NewStore(s.db)
	grpchandler.NewGRPCProductsService(grpcServer, store)
	log.Println("Grpc server started at ", s.addr)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatal()
	}

}

// greeterServer implements the GreeterServer interface generated from the .proto file.
// type greeterServer struct {
// 	pb.UnimplementedGreeterServer // Embed to provide forward compatibility
// }

// // server side below

// // SayHello implements the SayHello RPC method.
// func (s *greeterServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
// 	// Business logic: Create a greeting message
// 	message := "Hello, " + req.Name + "!"
// 	return &pb.HelloReply{Message: message}, nil
// }

// func main() {
// 	// Create a TCP listener on port 50051
// 	lis, err := net.Listen("tcp", ":50051")
// 	if err != nil {
// 		log.Fatalf("failed to listen: %v", err)
// 	}

// 	// Create a new gRPC server
// 	s := grpc.NewServer()

// 	// Register the Greeter service implementation
// 	pb.RegisterGreeterServer(s, &greeterServer{})

// 	// Start the server
// 	log.Printf("gRPC server listening on :50051")
// 	if err := s.Serve(lis); err != nil {
// 		log.Fatalf("failed to serve: %v", err)
// 	}
// }

//--------------------------
// Client side below
// package main

// import (
//     "context"
//     "log"
//     "time"

//     pb "github.com/Chandra5468/cfp-Products-Service/internal/services/common/genproto/products"
//     "google.golang.org/grpc"
//     "google.golang.org/grpc/credentials/insecure"
// )

// func main() {
//     // Connect to the gRPC server
//     conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
//     if err != nil {
//         log.Fatalf("failed to connect: %v", err)
//     }
//     defer conn.Close()

//     // Create a Greeter client
//     client := pb.NewGreeterClient(conn)

//     // Set a context with timeout
//     ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//     defer cancel()

//     // Call the SayHello RPC
//     resp, err := client.SayHello(ctx, &pb.HelloRequest{Name: "World"})
//     if err != nil {
//         log.Fatalf("failed to call SayHello: %v", err)
//     }

//     // Print the response
//     log.Printf("Response from server: %s", resp.Message)
// }
