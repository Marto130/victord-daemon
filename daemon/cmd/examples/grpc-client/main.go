package main

import (
	"context"
	"log"
	"time"
	"victord/daemon/transport/grpc/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	conn, err := grpc.NewClient(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewVictorServiceClient(conn)

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
