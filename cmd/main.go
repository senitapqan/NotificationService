package main

import (
	"goNotificationService/internal/handler"
	"goNotificationService/internal/service"
	"goNotificationService/server"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"github.com/joho/godotenv"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatal().Err(err).Msg("some error with initializiing")	
	}
	service := service.NewService()
	handler := handler.NewHandler(service)

	srv := new(server.Server)

	if err := srv.RunServer(viper.GetString("port"), handler.InitRoutes()); err != nil {
		log.Fatal().Err(err).Msg("error with run server")
	}
}

func initConfig() error {
	if err := godotenv.Load(); err != nil {
		return err
	}
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	viper.BindEnv("app_password", "APP_PASSWORD")
	viper.BindEnv("app_username", "APP_USERNAME")
	return nil
}