package main

import (
	"github.com/fungerouscode/go-ambassador/pkg/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	if err := initConfigs(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	db, err := repository.NewMysqlDB(repository.Config{
		Host:     "db",
		Port:     "3306",
		Username: "root",
		Password: "root",
		DBName:   "ambassador",
	})

	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	_ = repository.NewRepository(db)

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Listen(":3000")
}

func initConfigs() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
