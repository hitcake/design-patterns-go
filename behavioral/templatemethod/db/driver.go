package db

import "fmt"

type Driver interface {
	GetConnection() (Connection, error)
}

type MySqlDriver struct {
	dbUrl    string
	username string
	password string
}

func (d *MySqlDriver) GetConnection() (Connection, error) {
	fmt.Println("MySqlDriver.GetConnection")
	return &MySqlConnection{}, nil
}
