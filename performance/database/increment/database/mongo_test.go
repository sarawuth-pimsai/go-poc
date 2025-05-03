package database

import (
	"context"
	"log"
	"testing"
)

func BenchmarkMongoIncrement(b *testing.B) {
	ctx := context.Background()
	mongo := NewMongo(ctx)
	b.ResetTimer()
	for b.Loop() {
		if _, ok := mongo.Increment(300); ok != nil {
			log.Println(ok)
		}
	}
}
