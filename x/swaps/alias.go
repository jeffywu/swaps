package swaps

import (
    "github.com/cosmos/peggy/x/swaps/types"
    "github.com/cosmos/peggy/x/swaps/keeper"
)

const (
    ModuleName = types.ModuleName
    RouterKey  = types.RouterKey
    StoreKey   = types.StoreKey
)

var (
    NewKeeper = keeper.NewKeeper
)

type (
    Keeper = keeper.Keeper
)