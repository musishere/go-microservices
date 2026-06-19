package server

import (
	"context"

	"github.com/hashicorp/go-hclog"
	currency "github.com/musishere/grpc/protos"
)

type Currency struct {
	currency.UnimplementedCurrencyServer
	log hclog.Logger
}

func NewCurrency(l hclog.Logger) *Currency {
	return &Currency{
		log: l,
	}
}

func (c *Currency) GetRate(ctx context.Context, in *currency.RateRequest) (*currency.RateResponse, error) {
	c.log.Info("Handle GetRate", "base", in.GetBase(), "destination", in.Destination)
	return &currency.RateResponse{
		Rate: 30.5,
	}, nil
}
