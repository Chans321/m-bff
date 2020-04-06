package main

import (
	"flag"

	"github.com/Chans321/m-bff/bff"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func main() {
	addressHighScore := flag.String("address-m-highscore", "localhost:50051", "The grpc address for m-highscore microservice")
	addressGameEngine := flag.String("address-m-game-engine", "localhost:60051", "The grpc address for m-game-engine microservice")
	serverAddress := flag.String("address-http", ":8081", "HTTP server address")
	flag.Parse()
	gameClient, err := bff.NewGameServiceClient(*addressHighScore)
	if err != nil {
		log.Info().Msg("error in creating client for m-highscore")
	}
	gameEngineClient, err := bff.NewGameEngineServiceClient(*addressGameEngine)
	if err != nil {
		log.Info().Msg("error in creating client for m-game-engine")
	}
	gr := bff.NewGameResource(gameClient, gameEngineClient)
	router := gin.Default()
	router.GET("/geths", gr.GetHighScore)
	router.GET("/seths/:hs", gr.SetHighScore)
	router.GET("/getsize", gr.GetSize)
	router.GET("/setscore/:score", gr.SetScore)
	err = router.Run(*serverAddress)
	if err != nil {
		log.Info().Msg("Couldn't start bff")
	}

	log.Info().Msgf("started http server at %v", *serverAddress)

}
