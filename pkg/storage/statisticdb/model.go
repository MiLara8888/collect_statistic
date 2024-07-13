package statisticdb

import (
	"time"

	"github.com/milara8888/collect_statistic/pkg/storage"
)



type Table any

func Tables() []Table {
	return []Table{
		&OrderBook{},
		&OrderHistory{},
		// &DepthOrder{},
	}
}

//Книга заказов
//TODO индексы?
type OrderBook struct {
	ID       int64  `gorm:"primaryKey;index:,unique ;notnull" json:"id"`

	//TODOподразумеваем ли мы уникальность по этим двум полям??
	Exchange string `gorm:"size:1000 ;notnull" json:"exchange"`
	// отношение цен двух валют, входящих в данную пару, на валютном рынке
	Pair string `gorm:"size:1000 ;notnull" json:"pair"`

	// аски — это цены продажи, устанавливаемые теми, кто владеет активом и хочет его продать
	Asks []storage.JSONB `gorm:"type:jsonb"`

	//представляют собой предложения в базовой валюте за единицу торгового актива
	Bids []storage.JSONB `gorm:"type:jsonb"`
}

// заказы клиентов, история
type OrderHistory struct {
	ClientName string `gorm:"size:1000 ;notnull" json:"client_name"`
	//обмен
	ExchangeName        string  `gorm:"size:1000 ;notnull" json:"exchange_name"`
	Label               string  `gorm:"size:1000 ;notnull" json:"label"`
	Pair                string  `gorm:"size:1000 ;notnull" json:"pair"`
	Side                string  `gorm:"size:1000 ;notnull" json:"side"`
	Type                string  `gorm:"size:1000 ;notnull" json:"type"`
	BaseQty             float64 `gorm:"type:real ;notnull" json:"base_qty"`
	Price               float64 `gorm:"type:numeric ;notnull" json:"price"`
	AlgorithmNamePlaced string  `gorm:"size:1000 ;notnull" json:"algorithm_name_placed"`
	//самая низкая цена покупки
	LowestSellPrc float64 `gorm:"type:numeric ;notnull" json:"lowest_sell_prc"`
	//самая высокая цена покупки
	HighestBuyPrc float64 `gorm:"type:numeric ;notnull" json:"highest_buy_prc"`
	//комиссионное предложение кол-во
	CommissionQuoteQty float64 `gorm:"type:numeric ;notnull" json:"commission_quote_qty"`
	//время размещения
	TimePlaced time.Time `gorm:"autoUpdateTime:milli"`
}
