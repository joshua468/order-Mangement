package main

import (
	"context"
)

type OrdersService interface {
	CreateOrder(context.Context) error
}

type OrderStore interface {
	create(context.Context) error
}

func main() {

}
