package main

import (
	_ "github.com/joho/godotenv/autoload"
)

import (
	"context"
	"github.com/teameosb/edex/backend/cli"
	"github.com/teameosb/edex/backend/dex_engine"
	"github.com/teameosb/eosb-sdk-backend/utils"
	"os"
)

func run() int {
	ctx, stop := context.WithCancel(context.Background())
	go cli.WaitExitSignal(stop)

	dex_engine.Run(ctx, utils.StartMetrics)
	return 0
}

func main() {
	os.Exit(run())
}
