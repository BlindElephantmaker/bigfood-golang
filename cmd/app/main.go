package main

import (
	"bigfood/internal/controller"
	"bigfood/internal/infrastructure"
	"bigfood/pkg/database"
	"bigfood/pkg/server"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
	"os"
)

// @title        Bigfood
// @version      1.0
// @description  Internal API

// @host      localhost:8000
// @BasePath  /

// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        Authorization
func main() {
	// todo: add env production and development

	// todo: move initializing from here
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	// todo: test server can't read config, fix it
	//initConfig()

	if err := godotenv.Load(); err != nil { // todo: it is use global. How move to local?
		log.Fatalf("failed to load env variables: %s", err.Error())
	}

	logLevel := viper.GetString("server.log-level")
	if err := initLogger(logLevel); err != nil {
		log.Fatalf("error initializing logger: %s", err.Error())
	}

	configPsql := database.NewConfig(
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DATABASE"),
		os.Getenv("POSTGRES_SSL"),
	)
	db, err := database.NewPostgresDB(configPsql)
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	repositories := infrastructure.NewRepositories(db)
	services := infrastructure.NewServices(repositories, db)
	handlers := infrastructure.NewHandlers(repositories, services)
	controllers := controller.NewController(handlers)

	serverPort := viper.GetString("server.port")
	s := server.NewServer(serverPort, controllers.InitRoutes())

	if err := s.Run(); err != nil {
		logrus.Fatalf("error occured while running server server: %s", err.Error())
	}

	// todo: add shutdown handler
}

// todo: it is use global. How move to local?
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}

// todo: test server can't read config, fix it
//func initConfig() {
//	viper.Set("server.port", "8000")
//	viper.Set("server.log-level", "debug")
//}

// todo: it is use global. How move to local?
func initLogger(logLevel string) error {
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		return err
	}

	logrus.SetFormatter(new(logrus.JSONFormatter))
	logrus.SetLevel(level)

	return nil
}
