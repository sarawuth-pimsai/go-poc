# Go Language with increment data performance
Test performance increment function
## Setup
### Redis Driver
```go
go get github.com/redis/go-redis/v9
```
### Mongo Driver
```go
go get go.mongodb.org/mongo-driver/v2/mongo
```
### Mongo JSON
```json
{
  "_id": {
    "$oid": "68161b3a6d2cade9cf931e1f"
  },
  "product_id": 1,
  "quantity": 0.12
}
```
### MySQL Driver
```go
go get -u github.com/go-sql-driver/mysql
```
### MySQL Table
```sql
CREATE TABLE `increment` (
  `id` int NOT NULL AUTO_INCREMENT,
  `product_id` int NOT NULL DEFAULT '0',
  `quantity` decimal(10,4) NOT NULL DEFAULT '0.0000',
  PRIMARY KEY (`id`),
  KEY `idx_product_id` (`product_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
```
### MySQL Data
```sql
INSERT INTO `increment`(`product_id`, `quantity`)VALUES(1, 0.01)
```
## Benchmark
### All Memory
```go
go test -v ./... -bench=. -benchmem
```
### Redis Benchmark
```go
go test -v ./... -bench=BenchmarkRedisIncrement
```
### Mongo Benchmark
```go
go test -v ./... -bench=BenchmarkMongoIncrement
```
### MySQL Benchmark
```go
go test -v ./... -bench=BenchmarkMongoIncrement
```
