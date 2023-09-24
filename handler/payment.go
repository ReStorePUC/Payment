package handler

import (
	"context"
	"errors"
	pb "github.com/ReStorePUC/protobucket/payment"
	"github.com/eduardo-mior/mercadopago-sdk-go"
)

type PaymentConfig struct {
	AccessToken string `yaml:"access_token"`
	RedirectURL string `yaml:"redirect_url"`
}

// PaymentServer is used to implement Payment.
type PaymentServer struct {
	pb.UnimplementedPaymentServer

	config *PaymentConfig
}

func NewPaymentServer(cfg *PaymentConfig) *PaymentServer {
	return &PaymentServer{
		config: cfg,
	}
}

func (s *PaymentServer) CreatePayment(ctx context.Context, req *pb.CreatePaymentRequest) (*pb.CreatePaymentResponse, error) {
	paymentItems := []mercadopago.Item{}
	for _, item := range req.Items {
		paymentItems = append(paymentItems, mercadopago.Item{
			Title:     item.Title,
			Quantity:  float64(item.Quantity),
			UnitPrice: float64(item.UnitPrice),
		})
	}

	response, mercadopagoErr, err := mercadopago.CreatePayment(mercadopago.PaymentRequest{
		Items: paymentItems,
		BackUrls: mercadopago.BackUrls{
			Success: s.config.RedirectURL,
		},
	}, s.config.AccessToken)

	if mercadopagoErr != nil {
		return nil, errors.New(mercadopagoErr.Error)
	}

	if err != nil {
		return nil, err
	}

	return &pb.CreatePaymentResponse{
		Id: response.ID,
	}, nil
}
