package greeter

import (
    "fmt"

    sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewHandler returns a handler for "greeter" type messages.
func NewHandler(keeper Keeper) sdk.Handler {
    return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
        switch msg := msg.(type) {
        case MsgOrder:
            return handleMsgOrder(ctx, keeper, msg)
        default:
            errMsg := fmt.Sprintf("Unrecognized greeter Msg type: %v", msg.Type())
            return sdk.ErrUnknownRequest(errMsg).Result()
        }
    }
}

func handleMsgOrder(ctx sdk.Context, keeper Keeper, msg MsgOrder) sdk.Result {
    if true {
        // TODO: implement
        fmt.Println("inside here")
        return sdk.ErrUnauthorized("TODO, implement").Result() 
    }

    return sdk.Result{}
}