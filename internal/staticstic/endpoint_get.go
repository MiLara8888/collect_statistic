package staticstic

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/milara8888/collect_statistic/pkg/storage"
)

// извлечение записи по имени и паре  из таблицы OrderBook
func (s *Statistic) GetOrderBook(c *gin.Context) {

	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	data := struct {
		ExchangeName string `json:"exchange_name"`
		Pair         string `json:"pair"`
	}{}

	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	res, err := s.DB.GetOrderBook(c.Request.Context(), data.ExchangeName, data.Pair)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, res)
}

//извлечение заявки по клиенту
func (s *Statistic) GetOrderHistory(c *gin.Context) {

	limit, err := strconv.Atoi(c.Param("limit"))
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	offset, err := strconv.Atoi(c.Param("offset"))
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	data := struct {
		Сlient *storage.ClientSerializer `json:"client"`
	}{}

	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	res, cnt, err := s.DB.GetOrderHistory(c.Request.Context(), data.Сlient, offset, limit)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	result := struct {
		Data  []*storage.HistoryOrderSerializer `json:"data"`
		Count int                               `json:"count"`
	}{
		Data:  res,
		Count: cnt,
	}
	c.AbortWithStatusJSON(http.StatusOK, result)
}
