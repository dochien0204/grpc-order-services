package services

import (
	"context"
	"fmt"
	"net/http"
	"order_svc/pkg/client"
	"order_svc/pkg/db"
	"order_svc/pkg/models"
	"order_svc/pkg/pb"
)

type Server struct {
	H               db.Handler
	ProductServices client.ProductServicesClient
}

func (s *Server) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	//check product exists
	product, err := s.ProductServices.FindOne(req.ProductId)

	if err != nil {
		fmt.Println("Hello")
		return &pb.CreateOrderResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, nil
	} else if product.Status >= http.StatusNotFound {
		return &pb.CreateOrderResponse{
			Status: product.Status,
			Error:  product.Error,
		}, nil
	}

	order := models.Order{
		Quantity:  req.Quantity,
		ProductId: product.Data.Id,
	}

	s.H.DB.Create(&order)

	return &pb.CreateOrderResponse{
		Status: http.StatusOK,
		Id:     order.Id,
	}, nil
}
