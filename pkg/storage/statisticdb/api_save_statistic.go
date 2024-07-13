package statisticdb

import (
	"context"

	er "github.com/milara8888/collect_statistic/pkg/errors"
	"github.com/milara8888/collect_statistic/pkg/storage"
)

// сохранение заявки клиента
func (s *StatisticDB) SaveOrder(ctx context.Context, cl *storage.ClientSerializer, ho *storage.HistoryOrderSerializer) error {

	if cl.ClientName != ho.ClientName ||
		cl.ExchangeName != ho.ExchangeName ||
		cl.Label != ho.Label ||
		cl.Pair != ho.Pair {
		return er.ErrorMetchData
	}
	sql := `INSERT INTO  public.order_histories (client_name, exchange_name, "label",
				pair, side, "type", base_qty, price,
				algorithm_name_placed, lowest_sell_prc, highest_buy_prc, commission_quote_qty, time_placed)
				VALUES (@client_name, @exchange_name, @label, @pair, @side, @type, ROUND(@base_qty, 4), ROUND(@price, 4), @algorithm_name_placed,
				ROUND(@lowest_sell_prc, 4), ROUND(@highest_buy_prc, 4), ROUND(@commission_quote_qty, 4), @time_placed)`

	res := s.DB.WithContext(ctx).Exec(sql, map[string]any{
		"client_name":           ho.ClientName,
		"exchange_name":         ho.ExchangeName,
		"label":                 ho.Label,
		"pair":                  ho.Pair,
		"side":                  ho.Side,
		"type":                  ho.Type,
		"base_qty":              ho.BaseQty,
		"price":                 ho.Price,
		"algorithm_name_placed": ho.AlgorithmNamePlaced,
		"lowest_sell_prc":       ho.LowestSellPrc,
		"highest_buy_prc":       ho.HighestBuyPrc,
		"commission_quote_qty":  ho.CommissionQuoteQty,
		"time_placed":           ho.TimePlaced,
	})
	if res.Error != nil {
		return res.Error
	}
	return nil
}

// добавление записи в книгу звявок
func (s *StatisticDB) SaveOrderBook(ctx context.Context, name, pair string, da *storage.JSONB) error {

	var ob int

	//если запись есть, дополнить инфу
	sql := `select id from order_books where exchange = @name and pair = @pair`
	res := s.DB.WithContext(ctx).Raw(sql, map[string]any{
		"name": name,
		"pair": pair,
	}).Find(&ob)
	if res.Error != nil {
		return res.Error
	}

	x := storage.JSONB{}
	x.Scan(da)

	switch ob {
	//если объекта нет, создаю новый
	case 0:
		sql = `INSERT INTO order_books (exchange, pair, asks, bids)
				VALUES (@name, @pair, @asks, @bids)`

		res = s.DB.WithContext(ctx).Exec(sql, map[string]any{
			"name": name,
			"pair": pair,
			"asks": da,
			"bids": da,
		})
		if res.Error != nil {
			return res.Error
		}

	//если есть дополняю запись данными
	default:
		sql = `UPDATE order_books SET
				asks = asks || @asks ::jsonb,
				bids = bids || @bids ::jsonb
				WHERE exchange=@name and pair=@pair;`
		res = s.DB.WithContext(ctx).Exec(sql, map[string]any{
			"name": name,
			"pair": pair,
			"asks": da,
			"bids": da,
		})
		if res.Error != nil {
			return res.Error
		}
	}

	return nil
}
