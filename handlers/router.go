package handlers

import (
	"github.com/gorilla/mux"
	"github.com/izaakdale/service-dynamo/dao"
)

type ServiceOptions struct {
	DBClient dao.Client
}

type Service struct {
	Router   *mux.Router
	DBClient dao.Client
}

func New(opts ServiceOptions) *Service {
	return &Service{
		mux.NewRouter(),
		opts.DBClient,
	}
}
