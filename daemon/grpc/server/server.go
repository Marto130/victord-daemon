package server

import (
	"context"
	"log"
	"net"
	"victord/daemon/grpc/pb"

	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedVictorServiceServer
}

func (*Server) CreateIndex(ctx context.Context, req *pb.CreateIndexRequest) (*pb.CreateIndexResponse, error) {
	log.Printf("CreateIndex: index_name=%s", req.IndexName)

	resp := &pb.CreateIndexResponse{
		Status:  "success",
		Message: "Index created successfully",
		Results: []*pb.CreateIndexResult{
			{
				IndexName: req.IndexName,
				Id:        "abc-123",
				Dims:      128,
				IndexType: "HNSW",
				Method:    "IVF",
			},
		},
	}
	return resp, nil
}

func StartGRPCServer() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	srv := grpc.NewServer()
	pb.RegisterVictorServiceServer(srv, &Server{})

	log.Println("Starting gRPC server on :50051")
	if err := srv.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

	log.Println("gRPC server stopped")
	listener.Close()
	log.Println("Listener closed")

}
