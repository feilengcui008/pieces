package service

import (
	"golang.org/x/net/context"
	pb "grpcapp/proto"
)

type OrderService struct {
}

func (s *OrderService) GetOrder(ctx context.Context, r *pb.OrderRequest) (w *pb.OrderResponse, err error) {
	w.OrderID = r.GetOrderID()
	return
}
