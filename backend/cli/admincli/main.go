package main

import (
	"github.com/teameosb/edex/backend/admin/cli"
	"github.com/teameosb/eosb-sdk-backend/utils"
	_ "github.com/joho/godotenv/autoload"
	"os"
)

func main() {
	cli := admincli.NewDexCli()
	err := cli.Run(os.Args)
	if err != nil {
		utils.Errorf(err.Error())
	}
}
