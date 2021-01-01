package main

import (
	"database/sql"
	"flag"
	"github.com/pressly/goose"
	"github.com/spf13/viper"
	"log"
	"server/src/server"
	"time"
)

func main() {
	time.Sleep(time.Second * 10)
	flag.Parse()
	config := &server.Config{}

	viper.SetConfigName("configs")
	viper.AddConfigPath("./configs")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(config)
	if err != nil {
		log.Fatalf("unable to decode config into struct, %v", err)
	}

	if err := migrate(config.DbConnection); err != nil {
		log.Fatalf("migration error: %s", err)
	}
	if err := server.Start(config); err != nil {
		log.Fatal(err)
	}
}

func migrate(dbConnection string) error {
	command := "up"
	dir := "./migrations"
	db, err := sql.Open("mysql", dbConnection)
	if err != nil {
		return err
	}
	defer db.Close()
	if err := goose.SetDialect("mysql"); err != nil {
		return err
	}

	if err := goose.Run(command, db, dir); err != nil {
		return err
	}
	return nil
}