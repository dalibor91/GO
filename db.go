package db

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"errors"
)

const HOST_PARAM		= "host"
const PORT_PARAM		= "port"
const USER_PARAM		= "username"
const DATABASE_PARAM	= "database"
const PASSWORD_PARAM	= "password"
const CONNECTION_STRING_PARAM = "connection_string"

type MysqlConnection struct {
	connection *sql.DB
	connectionData map[string]string
}

func (m*MysqlConnection) Get(name string) string {
	return m.connectionData[name]
}

func (m*MysqlConnection) Set(name string, value string) * MysqlConnection {
	m.connectionData[name] = value
	return m
}

func (m*MysqlConnection)init() {
	if m.Get(CONNECTION_STRING_PARAM) == "" {
		m.Set(CONNECTION_STRING_PARAM, "%s:%s@tcp(%s:%s)/%s?tls=skip-verify&autocommit=true")
	}

	if m.Get(PORT_PARAM) == "" {
		m.Set(PORT_PARAM, "3306")
	}

	if m.Get(HOST_PARAM) == "" {
		m.Set(HOST_PARAM, "localhost")
	}

	m.connection = nil
}

func (m*MysqlConnection)Connect() *MysqlConnection {
	var err error

	m.connection, err = sql.Open("mysql", fmt.Sprintf(m.Get(CONNECTION_STRING_PARAM), m.Get(USER_PARAM), m.Get(PASSWORD_PARAM), m.Get(HOST_PARAM), m.Get(PORT_PARAM), m.Get(DATABASE_PARAM)));

	if err != nil {
		log.Println("Unable to connect ")
	}

	return m
}

func (m*MysqlConnection)Disconnect() {
	if m.connection != nil {
		defer m.connection.Close()
	} else { log.Println("Disconnect() -> not connected !") }
}

func (m*MysqlConnection)GetConnection() (error, *sql.DB){
	if m.connection == nil {
		return errors.New("Connection not opened"), nil
	}

	return nil, m.connection
}

func (m*MysqlConnection)Dump() {
	for k,v := range m.connectionData {
		fmt.Printf("%s => %s\n", k, v)
	}
}

func New() *MysqlConnection{
	c := new(MysqlConnection);
	c.connectionData = make(map[string]string)

	c.init()

	return c
}



