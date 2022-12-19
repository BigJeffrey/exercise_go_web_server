package persistence

import (
	"context"
	"database/sql"
	"log"
	"sync"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type PostgreSql struct {
	client *sql.DB
	ctx    context.Context
}

var once sync.Once
var instance *PostgreSql

func NewPostgreSql(ctx context.Context, conString string) *PostgreSql {

	once.Do(func() {
		instance = new(PostgreSql)
		instance.ctx = ctx
		err := instance.connect(conString)
		if err != nil {
			log.Println(err)
		}
	})

	return instance
}

func (m *PostgreSql) connect(conString string) error {
	var err error
	connectionstring := conString
	m.client, err = sql.Open("postgres", connectionstring)
	if err != nil {
		return err
	}
	return nil
}

func (m *PostgreSql) Disconnect() {
	err := m.client.Close()
	if err != nil {
		log.Println(err)
	}
}
