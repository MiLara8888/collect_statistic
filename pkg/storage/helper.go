package storage

import (

	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)


type JSON json.RawMessage

// Сканировать массив в Jsonb, описывает интерфейс sql.Scanner
func (j *JSON) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Ошибка распаковки значения JSONB:", value))
	}

	result := json.RawMessage{}
	err := json.Unmarshal(bytes, &result)
	*j = JSON(result)
	return err
}

// Возвращает значение json, описывает интерфейс driver.Valuer
func (j JSON) Value() (driver.Value, error) {
	if len(j) == 0 {
		return nil, nil
	}
	return json.RawMessage(j).MarshalJSON()
}



type JSONB []map[string]interface{}

func (a JSONB) Value() (driver.Value, error) {
	return json.Marshal(a)
}

func (p *JSONB) Scan(src interface{}) error {

	source, ok := src.([]byte)

	if !ok {
		return errors.New("type assertion .([]byte) failed")
	}

	var i interface{}
	err := json.Unmarshal(source, &i)
	if err != nil {
		return err
	}

	*p, ok = i.([]map[string]interface{})
	if !ok {
		return errors.New("type assertion .(map[string]interface{}) failed")
	}

	return nil
}
