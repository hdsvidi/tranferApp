package manager

import "github.com/jmoiron/sqlx"
import _ "github.com/lib/pq"


type Infra interface {
	SqlDb() *sqlx.DB
}
type infra struct {
	db *sqlx.DB
}

func (i *infra) SqlDb() *sqlx.DB {
	return i.db
}

func NewInfra(dataSourceName string) Infra {
	conn, err := sqlx.Connect("postgres",dataSourceName)
	if err != nil {
		panic(err)
	}
	return &infra{
		db: conn,
	}
}
