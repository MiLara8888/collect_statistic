# Микросервис на golang для сбора статистики

локально собрать проект команда : make local-build-run


в микросервисе реализовано 4 api ручки

-/collect_statistic/save/order_book - сохранение записи в таблицу заказов

в теле нужно передать структуру вида :

{
    "exchange_name": "test",
    "pair": "test",
    "data": [
        {
            "price":1,
            "base_qty":2
        },
         {
            "price":3,
            "base_qty":4
        }
    ]
}

если запись с таким exchange_name и pair существует, то данные дополнятся информацией из "data", если нет, то запись создасться


-/collect_statistic/get/book извлечение записи из таблицы заказов заиписи

в теле передать:

{
    "exchange_name": "test",
    "pair": "test"
}

вернётся структура вида:
[
    {
        "price": 1,
        "base_qty": 2
    },
    {
        "price": 3,
        "base_qty": 4
    }
]


-/collect_statistic/save/order_client добавление записи о заказе клиента
для добавления передать струкуру вида:

{
    "client": {
        "exchange_name": "test",
        "client_name": "test",
        "label": "test",
        "pair": "test"
    },
    "history_order": {
        "client_name": "test",
        "exchange_name": "test",
        "label": "test",
        "pair": "test",
        "side": "test",
        "type": "test",
        "base_qty": 1,
        "price": 1,
        "algorithm_name_placed": "test",
        "lowest_sell_prc": 1,
        "highest_buy_prc": 1,
        "commission_quote_qty": 1,
        "time_placed": "0001-01-01T02:57:40+02:57"
    }
}

-/collect_statistic/get/history_client/:offset/:limit получение записей о заказах клиента

передать структуру, а так же использовать offset и limit
  "client": {
        "exchange_name": "test",
        "client_name": "test",
        "label": "test",
        "pair": "test"
    }

ответ

{
    "data": [
        {
            "client_name": "test",
            "exchange_name": "test",
            "label": "test",
            "pair": "test",
            "side": "test",
            "type": "test",
            "base_qty": 1,
            "price": 1,
            "algorithm_name_placed": "test",
            "lowest_sell_prc": 1,
            "highest_buy_prc": 1,
            "commission_quote_qty": 1,
            "time_placed": "0001-01-01T02:58:20+02:57"
        }
    ],
    "count": 1
}

