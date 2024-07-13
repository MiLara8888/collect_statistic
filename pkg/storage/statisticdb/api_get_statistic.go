package statisticdb

import (
	"context"
	"encoding/json"

	"github.com/milara8888/collect_statistic/pkg/storage"
)

// извлечение записи по имени и паре  из таблицы OrderBook
func (s *StatisticDB) GetOrderBook(ctx context.Context, name string, pair string) ([]*storage.DepthOrderSerializer, error) {

	var (
		res = []*storage.DepthOrderSerializer{}
		ob  = []*storage.JSON{}
	)

	sql := `select asks from order_books where exchange = @name and pair = @pair`

	row := s.DB.WithContext(ctx).Raw(sql, map[string]any{"name": name, "pair": pair}).Scan(&ob)
	if row.Error != nil {
		return nil, row.Error
	}

	for _, j := range ob {
		err := json.Unmarshal(*j, &res)
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}

//извлечение заявки по клиенту
func (s *StatisticDB) GetOrderHistory(ctx context.Context, client *storage.ClientSerializer, offset, limit int) ([]*storage.HistoryOrderSerializer, int, error) {
	var (
		res = []*storage.HistoryOrderSerializer{}
		cnt int
	)
	sql := `select *, count(*) over (PARTITION BY 1) as count
			from public.order_histories oh
			where oh.client_name = @name
			offset @offset limit @limit`

	rows, err := s.DB.WithContext(ctx).Raw(sql, map[string]any{"name": client.ClientName, "offset": offset, "limit": limit}).Rows()
	if err != nil {
		return nil, cnt, err
	}
	defer rows.Close()
	for rows.Next() {
		k := &storage.HistoryOrderSerializer{}
		err = rows.Scan(&k.ClientName, &k.ExchangeName, &k.Label, &k.Pair, &k.Side, &k.Type, &k.BaseQty,
			&k.Price, &k.AlgorithmNamePlaced, &k.LowestSellPrc, &k.HighestBuyPrc, &k.CommissionQuoteQty, &k.TimePlaced, &cnt)
		if err != nil {
			return nil, cnt, err
		}
		res = append(res, k)
	}
	return res, cnt, nil
}
