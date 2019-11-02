package types

import (
    "encoding/json"

    sdk "github.com/cosmos/cosmos-sdk/types"
)

type MsgOrder struct {
    OrderType       string              `json:"order_type"`
    Tokens          sdk.Coin           `json:"tokens"`
    Price           sdk.Uint           `json:"price"`
    Address         sdk.AccAddress      `json:"address"`
}

func NewMsgOrder(orderType string, tokens sdk.Coin, price sdk.Uint, address sdk.AccAddress) MsgOrder {
    return MsgOrder{
        OrderType: orderType,
        Tokens: tokens,
        Price: price,
        Address: address,
    }
}

// Route should be name of the module
func (msg MsgOrder) Route() string { return RouterKey }

func (msg MsgOrder) Type() string { return "create_order" }

func (msg MsgOrder) ValidateBasic() sdk.Error {
    // These checks need to be stateless
    if !(msg.OrderType == "bid" || msg.OrderType == "offer") {
        return ErrInvalidOrderType(msg.OrderType)
    }

    // NOTE: more complex amount checking will be done in the handler, comparing amount to risk profile
    if !msg.Tokens.IsValid() {
        return ErrInvalidTokenAmount()
    }

    if msg.Price.IsZero() {
        return ErrInvalidPrice()
    }

    if msg.Address.Empty() {
        return sdk.ErrInvalidAddress(msg.Address.String())
    }

    return nil
}

func (msg MsgOrder) GetSignBytes() []byte {
    b, err := json.Marshal(msg)
    if err != nil {
        panic(err)
    }
    return sdk.MustSortJSON(b)
}

func (msg MsgOrder) GetSigners() []sdk.AccAddress {
    return []sdk.AccAddress{msg.Address}
}
