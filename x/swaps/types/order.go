package types

import (
    sdk "github.com/cosmos/cosmos-sdk/types"
)

type StoredOrder struct {
    OrderType       string              `json:"order_type"`
    Tokens          sdk.Coin            `json:"tokens"`
    Price           sdk.Uint            `json:"price"`
    Address         sdk.AccAddress      `json:"address"`
    Auction         string              `json:"auction"`
    OrderId         string              `json:"order_id"`
}

func NewStoredOrder(orderType string, tokens sdk.Coin, price sdk.Uint, address sdk.AccAddress, auction string, orderId string) StoredOrder {
    return StoredOrder{
        OrderType: orderType,
        Tokens: tokens,
        Price: price,
        Address: address,
        Auction: auction,
        OrderId: orderId,
    }
}

func NewStoredOrderFromMsgOrder(msg MsgOrder, auction string, orderId string) StoredOrder {
    return NewStoredOrder(msg.OrderType, msg.Tokens, msg.Price, msg.Address, auction, orderId)
}