package keeper

import (
    "github.com/cosmos/cosmos-sdk/codec"
    sdk "github.com/cosmos/cosmos-sdk/types"
    "github.com/cosmos/cosmos-sdk/x/bank"
    "github.com/cosmos/peggy/x/swaps/types"
)

// Keeper maintains the link to storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
    CoinKeeper bank.Keeper

    storeKey sdk.StoreKey // Unexposed key to access store from sdk.Context

    cdc *codec.Codec // The wire codec for binary encoding/decoding.
}

// NewKeeper creates new instances of the nameservice Keeper
func NewKeeper(coinKeeper bank.Keeper, storeKey sdk.StoreKey, cdc *codec.Codec) Keeper {
    return Keeper{
        CoinKeeper: coinKeeper,
        storeKey:   storeKey,
        cdc:        cdc,
    }
}

func (k Keeper) GetOrders(ctx sdk.Context, auction sdk.Uint) []types.StoredOrder {
    iterator := k.GetAuctionIterator(ctx, auction)
    orders := make([]types.StoredOrder, 0)

    defer iterator.Close()

    for iterator.Valid() {
        var o types.StoredOrder
        v := iterator.Value() 
        k.cdc.MustUnmarshalBinaryBare(v, &o)
        orders = append(orders, o)
        iterator.Next()
    }

    return orders
}

func (k Keeper) GetOrdersByAddress(ctx sdk.Context, auction sdk.Uint, address sdk.AccAddress) []types.StoredOrder {
    iterator := k.GetAuctionIterator(ctx, auction)
    orders := make([]types.StoredOrder, 0)

    defer iterator.Close()

    for iterator.Valid() {
        var o types.StoredOrder
        v := iterator.Value() 
        k.cdc.MustUnmarshalBinaryBare(v, &o)
        if o.Address == address {
            orders = append(orders, o)
        }
        iterator.Next()
    }

    return orders
}

func (k Keeper) GetTrades(ctx sdk.Context, swap_period sdk.Uint) []types.StoredTrade {
    iterator := k.GetTradeIterator(ctx, swap_period)
    trades := make([]types.StoredTrade, 0)

    defer iterator.Close()

    for iterator.Valid() {
        var o types.StoredTrade
        v := iterator.Value() 
        k.cdc.MustUnmarshalBinaryBare(v, &o)
        trades = append(trades, o)
        iterator.Next()
    }

    return trades
}

func (k Keeper) GetTradesByAddress(ctx sdk.Context, swap_period sdk.Uint, address sdk.AccAddress) []types.StoredTrade {
    iterator := k.GetTradeIterator(ctx, swap_period)
    trades := make([]types.StoredTrade, 0)

    defer iterator.Close()

    for iterator.Valid() {
        var o types.StoredTrade
        v := iterator.Value() 
        k.cdc.MustUnmarshalBinaryBare(v, &o)
        if o.Address == address {
            trades = append(trades, o)
        }
        iterator.Next()
    }

    return trades
}

func (k Keeper) GetAuctionIterator(ctx sdk.Context, auction sdk.Uint) sdk.Iterator {
    store := ctx.KVStore(k.storeKey)
    return sdk.KVStorePrefixIterator(store, []byte(auction))
}

func (k Keeper) GetTradeIterator(ctx sdk.Context, swap_period sdk.Uint) sdk.Iterator {
    store := ctx.KVStore(k.storeKey)
    return sdk.KVStorePrefixIterator(store, []byte(swap_period))
}

func (k Keeper) PutOrder(ctx sdk.Context, order types.StoredOrder) {
    store := ctx.KVStore(k.storeKey)
    // This ensures that we will iterate by the auction prefix
    key := []byte(order.Auction + order.OrderId)
    store.Set(key, k.cdc.MustMarshalBinaryBare(order))
}

func (k Keeper) PutTrade(ctx sdk.Context, trade types.StoredTrade) {
    store := ctx.KVStore(k.storeKey)
    // This ensures that we will iterate by the swap_period prefix
    key := []byte(trade.swap_period + trade.TradeId)
    store.Set(key, k.cdc.MustMarshalBinaryBare(trade))
}

func (k Keeper) GetCurrentAuction(ctx sdk.Context) sdk.Uint {

}

func (k Keeper) GetCurrentSwapNum(ctx sdk.Context) sdk.Uint {

}

func (k Keeper) GetCurrentSwapRate(ctx sdk.Context) sdk.Uint {

}

func (k Keeper) GetLastAuctionRate(ctx sdk.Context) sdk.Uint {

}

func (k Keeper) GetOrderID(ctx sdk.Context) sdk.Uint {

}

func (k Keeper) GetTradeID(ctx sdk.Context) sdk.Uint {
    
}