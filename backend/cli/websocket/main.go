package main

import (
	"context"
	"github.com/teameosb/edex/backend/cli"
	"github.com/teameosb/edex/backend/connection"
	"github.com/teameosb/eosb-sdk-backend/common"
	"github.com/teameosb/eosb-sdk-backend/utils"
	"github.com/teameosb/eosb-sdk-backend/websocket"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	os.Exit(run())
}

func run() int {
	ctx, stop := context.WithCancel(context.Background())

	redisClient := connection.NewRedisClient(os.Getenv("HSK_REDIS_URL"))
	redisClient = redisClient.WithContext(ctx)

	go cli.WaitExitSignal(stop)

	// new a source queue
	queue, err := common.InitQueue(&common.RedisQueueConfig{
		Name:   common.EOSB_WEBSOCKET_MESSAGES_QUEUE_KEY,
		Ctx:    ctx,
		Client: redisClient,
	})

	if err != nil {
		panic(err)
	}

	// new a websocket server
	wsServer := websocket.NewWSServer(":3002", queue)

	websocket.RegisterChannelCreator(
		common.MarketChannelPrefix,
		websocket.NewMarketChannelCreator(&websocket.DefaultHttpSnapshotFetcher{
			ApiUrl: os.Getenv("HSK_API_URL"),
		}),
	)

	// Start the server
	// It will block the current process to listen on the `addr` your provided.
	go utils.StartMetrics()
	wsServer.Start(ctx)

	return 0
}
