package database

import (
	"context"
	"log"
	"testing"
)

func BenchmarkMySQLIncrement(b *testing.B) {
	ctx := context.Background()
	mysql := NewMySQL(ctx)
	b.ResetTimer()
	for b.Loop() {
		if _, ok := mysql.Increment(30); ok != nil {
			log.Println(ok)
		}
	}
}
