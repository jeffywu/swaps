package keeper

import (
    //"github.com/cosmos/cosmos-sdk/codec"

    sdk "github.com/cosmos/cosmos-sdk/types"
    abci "github.com/tendermint/tendermint/abci/types"
)

// NewQuerier is the module level router for state queries
func NewQuerier(keeper Keeper) sdk.Querier {
    return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
        switch path[0] {
        default:
            return nil, sdk.ErrUnknownRequest("unknown nameservice query endpoint")
        }
    }
}
