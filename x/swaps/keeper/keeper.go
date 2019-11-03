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

func (k Keeper) GetOrders(ctx sdk.Context, auction string) []types.StoredOrder {
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

func (k Keeper) GetAuctionIterator(ctx sdk.Context, auction string) sdk.Iterator {
    store := ctx.KVStore(k.storeKey)
    return sdk.KVStorePrefixIterator(store, []byte(auction))
}

func (k Keeper) PutOrder(ctx sdk.Context, order types.StoredOrder) {
    store := ctx.KVStore(k.storeKey)
    // This ensures that we will iterate by the auction prefix
    key := []byte(order.Auction + order.OrderId)
    store.Set(key, k.cdc.MustMarshalBinaryBare(order))
}