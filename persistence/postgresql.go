package persistence

import (
	"context"
	"database/sql"
	"log"
	"sync"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Postgresql struct {
	client *sql.DB
	ctx    context.Context
}

var once sync.Once
var instance *Postgresql

func NewPostgresql(ctx context.Context, conString string) *Postgresql {

	once.Do(func() {
		instance = new(Postgresql)
		instance.ctx = ctx
		err := instance.connect(conString)
		if err != nil {
			log.Println(err)
		}
	})

	return instance
}

func (m *Postgresql) connect(conString string) error {
	var err error
	connectionString := conString
	m.client, err = sql.Open("postgres", connectionString)
	if err != nil {
		return err
	}
	return nil
}

func (m *Postgresql) Disconnect() {
	err := m.client.Close()
	if err != nil {
		log.Println(err)
	}
}
