package main

import (
	"log"
	"os"

	"clean-arc/configs"

	"clean-arc/modules/servers"
	databases "clean-arc/pkg/databases"
	redisclient "clean-arc/pkg/redis"

	"github.com/joho/godotenv"
)

func main() {
	// Load dotenv config
	if err := godotenv.Load(".env"); err != nil {
		panic(err.Error())
	}
	cfg := new(configs.Configs)

	redisClient := redisclient.GetNewRedisClient()
	// Fiber configs
	cfg.App.Host = os.Getenv("FIBER_HOST")
	cfg.App.Port = os.Getenv("FIBER_PORT")

	// Database Configs
	cfg.PostgreSQL.Host = os.Getenv("DB_HOST")
	cfg.PostgreSQL.Port = os.Getenv("DB_PORT")
	cfg.PostgreSQL.Protocol = os.Getenv("DB_PROTOCOL")
	cfg.PostgreSQL.Username = os.Getenv("DB_USERNAME")
	cfg.PostgreSQL.Password = os.Getenv("DB_PASSWORD")
	cfg.PostgreSQL.Database = os.Getenv("DB_DATABASE")

	// New Database
	db, err := databases.NewPostgreSQLDBConnection(cfg)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer db.Close()

	s := servers.NewServer(cfg, db, redisClient)
	s.Start()
}
