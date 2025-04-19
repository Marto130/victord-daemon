package server

import (
	"context"
	"log"
	"net"
	gen "victord/daemon/grpc/gen/victord/daemon/grpc"

	"google.golang.org/grpc"
)

type Server struct {
	gen.UnimplementedVictorServiceServer
}

func (*Server) CreateIndex(ctx context.Context, req *gen.CreateIndexRequest) (*gen.CreateIndexResponse, error) {
	log.Printf("CreateIndex recibido: index_name=%s", req.IndexName)

	resp := &gen.CreateIndexResponse{
		Status:  "success",
		Message: "Index created successfully",
		Results: []*gen.CreateIndexResult{
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
	s := grpc.NewServer()
	gen.RegisterVictorServiceServer(s, &Server{})
	log.Println("Starting gRPC server on :50051")
	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
	log.Println("gRPC server stopped")
	listener.Close()
	log.Println("Listener closed")
}
