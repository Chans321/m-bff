package bff

import (
	"context"
	"strconv"

	pbgameengine "github.com/Chans321/m-apis/m-game-engine/v1"
	pbhighscore "github.com/Chans321/m-apis/m-highscore/v1"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

type gameResource struct {
	gameClient       pbhighscore.GameClient
	gameEngineClient pbgameengine.GameEngineClient
}

func NewGameResource(gameClient pbhighscore.GameClient, gameEngineClient pbgameengine.GameEngineClient) *gameResource {
	return &gameResource{
		gameClient:       gameClient,
		gameEngineClient: gameEngineClient,
	}
}
func NewGameServiceClient(address string) (pbhighscore.GameClient, error) {
	con, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal().Msgf("Failed to dial :%v", err)
	} else {
		log.Info().Msgf("Successfully connected to [%s]", address)
	}
	if con == nil {
		log.Info().Msg("m-highscore connection is nill in m-bff")
	}
	client := pbhighscore.NewGameClient(con)
	return client, nil

}
func NewGameEngineServiceClient(address string) (pbgameengine.GameEngineClient, error) {
	con, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal().Msgf("Failed to dial :%v", err)
	} else {
		log.Info().Msgf("Successfully connected to [%s]", address)
	}
	if con == nil {
		log.Info().Msg("m-game-engine connection is nill in m-bff")
	}
	client := pbgameengine.NewGameEngineClient(con)
	return client, nil

}

func (gr *gameResource) SetHighScore(c *gin.Context) {
	highScoreString := c.Param("hs")
	highScoreFloat64, err := strconv.ParseFloat(highScoreString, 64)
	if err != nil {
		log.Error().Err(err).Msg("failed to convert highscore to float")
	}
	gr.gameClient.SetHighScore(context.Background(), &pbhighscore.SetHighScoreRequest{
		HighScore: highScoreFloat64,
	})

}

func (gr *gameResource) GetHighScore(c *gin.Context) {
	highscoreResponse, err := gr.gameClient.GetHighScore(context.Background(), &pbhighscore.GetHighScoreRequest{})
	if err != nil {
		log.Error().Err(err).Msg("failed to convert highscore to float")
		return
	}
	hsString := strconv.FormatFloat(highscoreResponse.HighScore, 'e', -1, 64)
	c.JSONP(200, gin.H{
		"hs": hsString,
	})

}

func (gr *gameResource) GetSize(c *gin.Context) {
	sizeResponse, err := gr.gameEngineClient.GetSize(context.Background(), &pbgameengine.GetSizeRequest{})
	if err != nil {
		log.Error().Err(err).Msg("failed to get size in m-bff")
		return
	}
	c.JSONP(200, gin.H{
		"size": sizeResponse.GetSize(),
	})

}

func (gr *gameResource) SetScore(c *gin.Context) {
	scoreString := c.Param("score")
	scoreFloat64, _ := strconv.ParseFloat(scoreString, 64)
	_, err := gr.gameEngineClient.SetScore(context.Background(), &pbgameengine.SetScoreRequest{
		Score: scoreFloat64,
	})
	if err != nil {
		log.Error().Err(err).Msg("failed to set score in m-bff")
		return
	}

}
