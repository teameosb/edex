package dex_engine

import (
	"github.com/teameosb/edex/backend/models"
	"github.com/teameosb/eosb-sdk-backend/sdk"
	"github.com/teameosb/eosb-sdk-backend/utils"
)

func geteosbOrderFromModelOrder(orderJSON *models.OrderJSON) *sdk.Order {
	return sdk.NewOrderWithData(
		orderJSON.Trader,
		orderJSON.Relayer,
		orderJSON.BaseCurrency,
		orderJSON.QuoteCurrency,
		utils.DecimalToBigInt(orderJSON.BaseCurrencyHugeAmount),
		utils.DecimalToBigInt(orderJSON.QuoteCurrencyHugeAmount),
		utils.DecimalToBigInt(orderJSON.GasTokenHugeAmount),
		orderJSON.Data,
		orderJSON.Signature,
	)
}

func geteosbOrderHashHexFromOrderJson(orderJSON *models.OrderJSON) string {
	order := sdk.NewOrderWithData(
		orderJSON.Trader,
		orderJSON.Relayer,
		orderJSON.BaseCurrency,
		orderJSON.QuoteCurrency,
		utils.DecimalToBigInt(orderJSON.BaseCurrencyHugeAmount),
		utils.DecimalToBigInt(orderJSON.QuoteCurrencyHugeAmount),
		utils.DecimalToBigInt(orderJSON.GasTokenHugeAmount),
		orderJSON.Data,
		"",
	)

	return utils.Bytes2HexP(eosbProtocol.GetOrderHash(order))
}
