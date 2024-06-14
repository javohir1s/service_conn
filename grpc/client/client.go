package client

import (
	"fmt"
	"go_catalog_service/config"
	"go_catalog_service/genproto/order_status_notes"
	"go_catalog_service/genproto/orders_service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceManagerI interface {
	OrderService() orders_service.OrderServiceClient
	OrderServiceStatus() order_status_notes.OrderStatusServiceClient
}

type grpcClients struct {
	orderService       orders_service.OrderServiceClient
	orderStatusService order_status_notes.OrderStatusServiceClient
}

func NewGrpcClients(cfg config.Config) (ServiceManagerI, error) {

	connOrderService, err := grpc.Dial(
		fmt.Sprintf("%s%s", cfg.OrderServicePort, cfg.OrderServiceHost),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(52428800), grpc.MaxCallSendMsgSize(52428800)),
	)
	if err != nil {
		return nil, err
	}

	connOrderStatusService, err := grpc.Dial(
		fmt.Sprintf("%s%s", cfg.OrderServicePort, cfg.OrderServiceHost),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(52428800), grpc.MaxCallSendMsgSize(52428800)),
	)
	if err != nil {
		return nil, err
	}

	return &grpcClients{
		orderService:       orders_service.NewOrderServiceClient(connOrderService),
		orderStatusService: order_status_notes.NewOrderStatusServiceClient(connOrderStatusService),
	}, nil
}

func (g *grpcClients) OrderService() orders_service.OrderServiceClient {
	return g.orderService
}

func (g *grpcClients) OrderServiceStatus() order_status_notes.OrderStatusServiceClient {
	return g.orderStatusService
}
