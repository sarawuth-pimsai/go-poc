package main

import (
	"context"
	"go/poc/performance/database/increment/database"
	"log"
)

var ctx = context.Background()

func main() {
	r := database.NewRedis(ctx)
	result, ok := r.Increment(100)
	if ok != nil {
		log.Println(ok)
	}
	log.Println("Redis: ", result)
	m := database.NewMongo(ctx)
	result, ok = m.Increment(100)
	if ok != nil {
		log.Println(ok)
	}
	log.Println("Mongo: ", result)
	my := database.NewMySQL(ctx)
	result, ok = my.Increment(100)
	if ok != nil {
		log.Println(ok)
	}
	log.Println("MySQL: ", result)
}
