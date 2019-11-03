package types

import (
    sdk "github.com/cosmos/cosmos-sdk/types"
)

type StoredTrade struct {
    OrderType       string              `json:"trade_type"`
    Tokens          sdk.Coin            `json:"tokens"`
    Price           sdk.Uint            `json:"price"`
    Address         sdk.AccAddress      `json:"address"`
    Swap_period         sdk.Uint              `json:"auction"`
    TradeId         sdk.Uint              `json:"trade_id"`
}

func NewStoredTrade(tradeType string, tokens sdk.Coin, price sdk.Uint, address sdk.AccAddress, swap_period sdk.Uint, tradeId sdk.Uint) StoredTrade {
    return StoredTrade{
        TradeType: tradeType,
        Tokens: tokens,
        Price: price,
        Address: address,
        Swap_period: swap_period,
        TradeId: tradeId,
    }
}

