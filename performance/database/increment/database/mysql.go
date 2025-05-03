package database

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type MySQL struct {
	ctx context.Context
	db  *sql.DB
}

func NewMySQL(ctx context.Context) *MySQL {
	db, err := sql.Open("mysql", "user_performance:pa9^J#s9NOqmGv3C@/performance")
	if err != nil {
		log.Println(err)
	}
	db.SetConnMaxIdleTime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	mysql := new(MySQL)
	mysql.db = db
	mysql.ctx = ctx
	return mysql
}

func (m *MySQL) Increment(value uint16) (uint16, error) {
	query := "UPDATE increment SET quantity = quantity + ? WHERE product_id = 1"
	ctx, cancel := context.WithTimeout(m.ctx, 5*time.Second)
	defer cancel()
	_, err := m.db.ExecContext(ctx, query, value)
	if err != nil {
		log.Println(err)
	}
	return value, nil
}
