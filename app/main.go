package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	subStart "github.com/gcp-iot/implementation/_start/pubsub"
	subService "github.com/gcp-iot/implementation/service/http"
	subUsecase "github.com/gcp-iot/implementation/usecase"
	"github.com/rs/zerolog/log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/spf13/viper"
)

func init() {
	path, err := os.Getwd()
	if err != nil {
		log.Error().Err(err).Msg("")
	}
	fmt.Println(`path: ` + path)
	viper.SetConfigType(`json`)
	viper.SetConfigName(`config`)
	viper.AddConfigPath(`./`)
	viper.AddConfigPath(`../`)

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("config file not found")
		}
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Info().Msg("Service RUN on DEBUG mode")
	}
}

func main() {

	fmt.Println("Go Time")

	flag.Parse()

	viper.AutomaticEnv()
	viper.SetEnvPrefix(viper.GetString("ENV"))
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodPatch, http.MethodDelete},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	timeoutContext := time.Duration(viper.GetInt("CONTEXT.TIMEOUT")) * time.Second
	baseUrl := viper.GetString("ENV_BASE_URL")
	if baseUrl == "" {
		log.Fatal().Msg("Base Url Not Found")

	}
	subId := viper.GetString("ENV_SUB_ID")
	if subId == "" {
		log.Fatal().Msg("Sub Id Not Found")

	}
	projectId := viper.GetString("ENV_PROJECT_ID")
	if projectId == "" {
		log.Fatal().Msg("PROJECT ID Not Found")

	}
	_subService := subService.NewSubService(timeoutContext)
	_subUsecase := subUsecase.NewSubUsecase(_subService, timeoutContext, baseUrl)
	subStart.NewSubHandler(subId, projectId, _subUsecase)

	//log.Fatal(e.Start(viper.GetString("ENV_AUTH_SERVER")))
}
