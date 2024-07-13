package storage

import (
	"context"
)

type IStatisticDB interface {
	// сохранение заявки клиента
	SaveOrder(ctx context.Context, cl *ClientSerializer, ho *HistoryOrderSerializer) error

	// извлечение записи по имени и паре  из таблицы OrderBook
	GetOrderBook(ctx context.Context, name string, pair string) ([]*DepthOrderSerializer, error)

	//извлечение заявки по клиенту
	GetOrderHistory(ctx context.Context, client *ClientSerializer, offset, limit int) ([]*HistoryOrderSerializer, int, error)

	//добавление записи в книгу звявок
	SaveOrderBook(ctx context.Context, name string, pair string, da *JSONB) error

	Migrate(ctx context.Context) error

	Close(ctx context.Context) error
}
