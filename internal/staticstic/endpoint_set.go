package staticstic

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	er "github.com/milara8888/collect_statistic/pkg/errors"
	"github.com/milara8888/collect_statistic/pkg/storage"
)

// сохранение заявки клиента
func (s *Statistic) SaveOrder(c *gin.Context) {

	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	data := struct {
		Client *storage.ClientSerializer       `json:"client"`
		Order  *storage.HistoryOrderSerializer `json:"history_order"`
	}{}

	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err = s.DB.SaveOrder(c.Request.Context(), data.Client, data.Order)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	c.AbortWithStatus(http.StatusCreated)
}

//добавление записи в книгу звявок
func (s *Statistic) SaveOrderBook(c *gin.Context) {

	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	data := struct {
		EchangeName string         `json:"exchange_name"`
		Pair        string         `json:"pair"`
		Data        *storage.JSONB `json:"data"`
	}{}

	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	err = s.DB.SaveOrderBook(c.Request.Context(), data.EchangeName, data.Pair, data.Data)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, er.ErrorSaving.Error())
		return
	}
	c.AbortWithStatus(http.StatusCreated)
}
