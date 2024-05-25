package main

type service struct {
	store OrdersService
}

func NewService(store OrdersService) *service {
	return &service{store}
}

func (s *service) CreateOrder() {
	return nil
}
