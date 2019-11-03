package swaps

import (
    "fmt"

    sdk "github.com/cosmos/cosmos-sdk/types"
    "github.com/cosmos/peggy/x/swaps/types"
)

// NewHandler returns a handler for "greeter" type messages.
func NewHandler(keeper Keeper) sdk.Handler {
    return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
        switch msg := msg.(type) {
        case types.MsgOrder:
            return handleMsgOrder(ctx, keeper, msg)
        default:
            errMsg := fmt.Sprintf("Unrecognized greeter Msg type: %v", msg.Type())
            return sdk.ErrUnknownRequest(errMsg).Result()
        }
    }
}

func handleMsgOrder(ctx sdk.Context, keeper Keeper, msg types.MsgOrder) sdk.Result {
    if msg.OrderType == "bid" {
        available_long_balance = getAvailableLongBalance(ctx, keeper, msg.address)
        if msg.Tokens.Amount > available_long_balance {
            return sdk.ErrUnauthorized("invalid order").Result() 
        }
        else {
            current_auction := keeper.GetCurrentAuction(ctx)
            order_id := keeper.GetOrderID(ctx)
            order := types.NewStoredOrderFromMsgOrder(msg, current_auction, order_id)
            keeper.PutOrder(ctx, order)
        }
    }
    else {
        available_short_balance = getAvailableShortBalance(ctx, keeper, msg.address)
        if msg.Tokens.Amount > available_short_balance {
            return sdk.ErrUnauthorized("invalid order").Result() 
        }
        else {
            current_auction := keeper.GetCurrentAuction(ctx)
            order_id := keeper.GetOrderID(ctx)
            order := types.NewStoredOrderFromMsgOrder(msg, current_auction, order_id)
            keeper.PutOrder(ctx, order)
        }
    }
    fmt.Println("inside here")
    return sdk.ErrUnauthorized("TODO, implement").Result() 

    return sdk.Result{}
}

func getAvailableLongBalance(ctx sdk.Context, keeper Keeper, address sdk.AccAddress) sdk.Int {
    current_auction := keeper.GetCurrentAuction(ctx)
    current_swap_num := keeper.GetCurrentSwapNum(ctx)
    current_swap_rate := keeper.GetCurrentSwapRate(ctx)
    last_auction_rate := keeper.GetLastAuctionRate(ctx)
    initial_block := keeper.GetInitialBlock(ctx)
    end_block := keeper.GetEndBlock(ctx)
    balance := keeper.CoinKeeper.GetCoins(ctx, address)
    collateralization_threshold := keeper.GetCollateralizationThreshold(ctx)
    current_trades := keeper.GetTradesByAddress(ctx, current_swap_num, address)
    future_trades := keeper.GetTradesByAddress(ctx, current_swap_num + 1, address)
    orders := keeper.GetOrdersByAddress(ctx, current_auction, address)
    current_trade_pnl := 0
    future_trade_pnl := 0
    current_position := 0
    future_position := 0

    for i := 0; i < len(current_trades); i++ {
        trade := current_trades[i]
        if trade.OrderType == "bid" {
            trade_pnl := trade.Tokens.Amount * (current_swap_rate - trade.Price)
            current_position += trade.Tokens.Amount
        }
        else {
            trade_pnl := trade.Tokens.Amount * (trade.Price - current_swap_rate)
            current_position -= trade.Tokens.Amount
        }
        current_trade_pnl += trade_pnl    
    }

    current_position_margin := (end_block - ctx.BlockHeight) / 1000 * abs(current_position) * collateralization_threshold

    for i := 0; i < len(future_trades); i++ {
        trade := future_trades[i]
        if trade.OrderType == "bid" {
            trade_pnl := trade.Tokens.Amount * (last_auction_rate - trade.Price)
            future_position += trade.Tokens.Amount
        }
        else {
            trade_pnl := trade.Tokens.Amount * (trade.Price - last_auction_rate)
            future_position -= trade.Tokens.Amount
        }
        future_trade_pnl += trade_pnl
    }

    for i := 0; i < len(orders); i++ {
        order := orders[i]
        if order.OrderType == "bid" {
            future_position += order.Tokens.Amount
        }
    }

    future_position_margin := abs(future_position) * collateralization_threshold

    available_long_balance := balance.Amount + current_trade_pnl + future_trade_pnl - current_position_margin - future_position_margin
    
    return available_long_balance
}

func getAvailableShortBalance(ctx sdk.Context, keeper Keeper, address sdk.AccAddress) sdk.Int {
    current_auction := keeper.GetCurrentAuction(ctx)
    current_swap_num := keeper.GetCurrentSwapNum(ctx)
    current_swap_rate := keeper.GetCurrentSwapRate(ctx)
    last_auction_rate := keeper.GetLastAuctionRate(ctx)
    initial_block := keeper.GetInitialBlock(ctx)
    end_block := keeper.GetEndBlock(ctx)
    balance := keeper.CoinKeeper.GetCoins(ctx, address)
    collateralization_threshold := keeper.GetCollateralizationThreshold(ctx)
    current_trades := keeper.GetTradesByAddress(ctx, current_swap_num, address)
    future_trades := keeper.GetTradesByAddress(ctx, current_swap_num + 1, address)
    orders := keeper.GetOrdersByAddress(ctx, current_auction, address)
    current_trade_pnl := 0
    future_trade_pnl := 0
    current_position := 0
    future_position := 0

    for i := 0; i < len(current_trades); i++ {
        trade := current_trades[i]
        if trade.OrderType == "bid" {
            trade_pnl := trade.Tokens.Amount * (current_swap_rate - trade.Price)
            current_position += trade.Tokens.Amount
        }
        else {
            trade_pnl := trade.Tokens.Amount * (trade.Price - current_swap_rate)
            current_position -= trade.Tokens.Amount
        }
        current_trade_pnl += trade_pnl    
    }

    current_position_margin := (end_block - ctx.BlockHeight) / 1000 * abs(current_position) * collateralization_threshold

    for i := 0; i < len(future_trades); i++ {
        trade := future_trades[i]
        if trade.OrderType == "bid" {
            trade_pnl := trade.Tokens.Amount * (last_auction_rate - trade.Price)
            future_position += trade.Tokens.Amount
        }
        else {
            trade_pnl := trade.Tokens.Amount * (trade.Price - last_auction_rate)
            future_position -= trade.Tokens.Amount
        }
        future_trade_pnl += trade_pnl
    }

    for i := 0; i < len(orders); i++ {
        order := orders[i]
        if order.OrderType == "offer" {
            future_position += order.Tokens.Amount
        }
    }

    future_position_margin := abs(future_position) * collateralization_threshold

    available_short_balance := balance.Amount + current_trade_pnl + future_trade_pnl - current_position_margin - future_position_margin
    
    return available_short_balance
}