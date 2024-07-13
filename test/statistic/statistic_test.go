package test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/milara8888/collect_statistic/internal/staticstic"
	"github.com/milara8888/collect_statistic/pkg/storage"
	"github.com/milara8888/collect_statistic/pkg/storage/statisticdb"
)

func TestStatisticStart(t *testing.T) {

	service, err := staticstic.New(config)
	if err != nil {
		t.Fatal(err)
	}
	err = service.Start()
	if err != nil {
		t.Fatal(err)
	}
}

func TestMigrate(t *testing.T) {
	db, err := statisticdb.New(config)
	if err != nil {
		return
	}
	err = db.Migrate(context.TODO())
	if err != nil {
		return
	}
}

func TestSaveOrdesBook(t *testing.T) {

	var n *storage.JSONB

	data := struct {
		EchangeName string         `json:"exchange_name"`
		Pair        string         `json:"pair"`
		Data        *storage.JSONB `json:"data"`
	}{
		EchangeName: "test",
		Pair: "test",
		Data: 	n,
	}

	b, err := json.Marshal(&data)
	if err != nil {
		t.Fatal(err)
	}

	bodyReader := bytes.NewReader(b)

	req, err := http.NewRequest("POST", "http://0.0.0.0:5313/collect_statistic/save/order_book", bodyReader)
	if err != nil {
		t.Error(t)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{
		Timeout: 300 * time.Second,
	}
	res, err := client.Do(req)
	if err != nil {
		t.Error(err)
		return
	}
	if res == nil {
		t.Error(errors.New(" res is nil"))
		return
	}
	if res.StatusCode != http.StatusCreated {
		t.Error(errors.New(" != 201 SaveOrderBook"))
		return
	}
}

func TestSaveClient(t *testing.T) {

	data := struct {
		Client *storage.ClientSerializer       `json:"client"`
		Order  *storage.HistoryOrderSerializer `json:"history_order"`
	}{
		Client: &storage.ClientSerializer{ClientName: "test",
			ExchangeName: "test",
			Label:        "test",
			Pair:         "test"},
		Order: &storage.HistoryOrderSerializer{ClientName: "test",
			ExchangeName:        "test",
			Label:               "test",
			Pair:                "test",
			Side:                "test",
			Type:                "test",
			BaseQty:             1,
			Price:               1,
			AlgorithmNamePlaced: "test",
			LowestSellPrc:       1,
			HighestBuyPrc:       1,
			CommissionQuoteQty:  1,
			TimePlaced:          time.Now(),
		},
	}

	b, err := json.Marshal(&data)
	if err != nil {
		t.Fatal(err)
	}

	bodyReader := bytes.NewReader(b)

	req, err := http.NewRequest("POST", "http://0.0.0.0:5313/collect_statistic/save/order_client", bodyReader)
	if err != nil {
		t.Error(t)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{
		Timeout: 300 * time.Second,
	}
	res, err := client.Do(req)
	if err != nil {
		t.Error(err)
		return
	}
	if res == nil {
		t.Error(errors.New(" res is nil"))
		return
	}
	if res.StatusCode != http.StatusCreated {
		t.Error(errors.New(" != 201 SaveOrderBook"))
		return
	}
}

func TestGetOrderHistory(t *testing.T) {

	data := struct {
		Сlient *storage.ClientSerializer `json:"client"`
	}{
		Сlient: &storage.ClientSerializer{ClientName: "test",
			ExchangeName: "test",
			Label:        "test",
			Pair:         "test",
		},
	}
	b, err := json.Marshal(&data)
	if err != nil {
		t.Fatal(err)
	}

	bodyReader := bytes.NewReader(b)

	req, err := http.NewRequest("POST", "http://0.0.0.0:5313/collect_statistic/get/history_client/0/500", bodyReader)
	if err != nil {
		t.Error(t)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{
		Timeout: 300 * time.Second,
	}
	res, err := client.Do(req)
	if err != nil {
		t.Error(err)
		return
	}
	if res == nil {
		t.Error(errors.New(" res is nil"))
		return
	}
	if res.StatusCode != http.StatusOK {
		t.Error(errors.New(" != 200"))
		return
	}
	b, err = io.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
		return
	}

	result := struct {
		Data  []*storage.HistoryOrderSerializer `json:"data"`
		Count int                               `json:"count"`
	}{}

	err = json.Unmarshal(b, &result)
	if err != nil {
		t.Error(err)
		return
	}
}
