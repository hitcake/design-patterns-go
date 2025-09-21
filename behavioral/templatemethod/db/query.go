package db

import (
	"errors"
	"fmt"
	"strconv"
)

type Query interface {
	QueryForInt(sql string) (int64, error)
}

type AbstractQuery struct {
}

// 定义抽象方法，需要子类实现
func (q *AbstractQuery) getConnection() (Connection, error) {
	return nil, errors.New("not implemented")
}

// 定义抽象方法，需要子类实现
func (q *AbstractQuery) closeConnection(c Connection) error {
	return errors.New("not implemented")
}
func (q *AbstractQuery) QueryForInt(sql string) (int64, error) {
	conn, err := q.getConnection()
	if err != nil {
		return 0, err
	}
	defer func(q *AbstractQuery, c Connection) {
		err := q.closeConnection(c)
		if err != nil {
			fmt.Println(err)
		}
	}(q, conn)
	result, err := conn.Query(sql)
	if err != nil {
		return 0, err
	}
	size := len(result)
	if size == 0 {
		return 0, errors.New("no result")
	} else if size > 1 {
		return 0, errors.New("more than 1 result")
	} else {
		obj := result[0]
		switch t := obj.(type) {
		case int:
			return int64(t), nil
		case int64:
			return t, nil
		case string:
			return strconv.ParseInt(t, 10, 64)
		default:
			return 0, errors.New("unknown type")
		}
	}
}

type DefaultQuery struct {
	AbstractQuery
	driver Driver
}

func NewDefaultQuery(driver Driver) *DefaultQuery {
	return &DefaultQuery{driver: driver}
}
func (q *DefaultQuery) getConnection() (Connection, error) {
	return q.driver.GetConnection()
}
func (q *DefaultQuery) closeConnection(c Connection) error {
	c.Close()
	return nil
}
