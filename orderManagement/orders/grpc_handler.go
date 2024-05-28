package main

import (
	"context"
	"log"

	pb "github.com/joshua468/commons/api"
	"google.golang.org/grpc"
)

type grpcHandler struct {
	pb.UnimplementedOrderServiceServer
	service OrdersService
}

func NewGRPCHandler(grpcServer *grpc.Server, service OrdersService) {
	handler := &grpcHandler{}
	pb.RegisterOrderServiceServer(grpcServer, handler)

}

func (h *grpcHandler) CreateOrder(ctx context.Context, p *pb.CreateOrderRequest) (*pb.Order, error) {

	log.Printf("New order received! Order %v", p)
	o := &pb.Order{
		ID: "42",
	}
	return o, nil
}
