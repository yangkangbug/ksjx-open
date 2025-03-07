package client

import (
	"context"
	"google.golang.org/grpc/credentials/insecure"
	"time"

	"google.golang.org/grpc"
	pb "ksjx-open/api/inventory" // 通过 protoc 生成的代码在这个目录下
)

var inventoryClient pb.InventoryServiceClient

// InitInventoryClient 负责初始化 gRPC 连接并创建 InventoryServiceClient 实例
func InitInventoryClient(grpcAddress string) error {
	transportCredentials := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.NewClient(grpcAddress, transportCredentials)
	if err != nil {
		return err
	}
	inventoryClient = pb.NewInventoryServiceClient(conn)
	return nil
}

// QueryInventory 使用已初始化的 gRPC 客户端调用远程库存查询服务
func QueryInventory(productID string) (int32, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &pb.InventoryRequest{ProductId: productID}

	resp, err := inventoryClient.QueryInventory(ctx, req)
	if err != nil {
		return 0, err
	}
	return resp.Stock, nil
}
