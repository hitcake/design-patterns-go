package db

import "fmt"

type Connection interface {
	Query(sql string) ([]interface{}, error)
	Close()
}

type MySqlConnection struct {
}

func (conn *MySqlConnection) Query(sql string) ([]interface{}, error) {
	fmt.Println("executing query: " + sql)
	return []interface{}{100}, nil
}
func (conn *MySqlConnection) Close() {
	fmt.Println("closing connection")
}
