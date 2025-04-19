package main

import (
	"context"
	"log"
	"time"

	pb "victord/daemon/grpc/gen/victord/daemon/grpc"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("No se pudo conectar: %v", err)
	}
	defer conn.Close()

	client := pb.NewVictorServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	req := &pb.CreateIndexRequest{
		IndexType: 1,
		Method:    1,
		Dims:      12,
		IndexName: "test-index",
	}

	res, err := client.CreateIndex(ctx, req)
	if err != nil {
		log.Fatalf("Error al llamar CreateIndex: %v", err)
	}

	log.Printf("Respuesta: %v", res)
}
