package staticstic

import (
	"net/http"
)

// доступ к двум методам
var postopt = []string{http.MethodOptions, http.MethodPost}

func (s *Statistic) initializeRoutes() {

	statisticGroup := s.Routes.Group("/collect_statistic", s.HostMiddleware(), CORSMiddleware())

	getGroup := statisticGroup.Group("/get")

	{
		// извлечение записи по имени и паре  из таблицы OrderBook
		getGroup.Match(postopt, "/book", s.GetOrderBook)

		//извлечение заявки по клиенту
		getGroup.Match(postopt, "/history_client/:offset/:limit", s.GetOrderHistory)
	}

	saveGroup := statisticGroup.Group("/save")

	{
		//сохранение заявки клиента
		saveGroup.Match(postopt, "/order_client", s.SaveOrder)

		//добавление записи в книгу звявок
		saveGroup.Match(postopt, "/order_book", s.SaveOrderBook)
	}

}

// GetOrderBook(exchange_name, pair string) ([]*DepthOrder, error)
// GetOrderHistory(client *Client) ([]*HistoryOrder, error)

// SaveOrder(client *Client, order *HistoryOrder) error
// SaveOrderBook(exchange_name, pair string, orderBook []*DepthOrder) error
