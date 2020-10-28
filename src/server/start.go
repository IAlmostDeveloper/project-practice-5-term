package server

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"net/http"
	"server/src/storage/mysql_storage"
)

func Start(config *Config) error {
	db, err := newDb(config.DbConnection)
	if err != nil {
		return err
	}
	defer db.Close()
	storage := mysql_storage.New(db)
	redis, err := newRedisClient(config.RedisAddr, config.RedisPassword)
	if err != nil {
		return err
	}
	server := NewServer(storage, redis)

	port := ":8080"
	fmt.Println("Port " + port)

	return http.ListenAndServe(port, server.router)
}

func newDb(databaseURL string) (*sqlx.DB, error) {
	db, err := sqlx.Open("mysql", databaseURL)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func newRedisClient(redisAddr string, redisPassword string) (*redis.Client, error) {
	redis := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword,
		DB:       0,
	})
	if res := redis.Ping(context.Background()); res.Err() != nil {
		return nil, res.Err()
	}

	return redis, nil
}
