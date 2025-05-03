package database

import (
	"context"
	"log"
	"testing"
)

func BenchmarkRedisIncrement(b *testing.B) {
	var ctx = context.Background()
	redis := NewRedis(ctx)
	b.ResetTimer()
	for b.Loop() {
		if _, ok := redis.Increment(3000); ok != nil {
			log.Println(ok)
		}
	}
}
