package cli

import (
    "errors"
    "fmt"
    "strconv"
    "github.com/spf13/cobra"

    //"github.com/cosmos/cosmos-sdk/client"
    //"github.com/cosmos/cosmos-sdk/client/context"
    "github.com/cosmos/cosmos-sdk/codec"
    sdk "github.com/cosmos/cosmos-sdk/types"
    //"github.com/cosmos/peggy/x/swaps/types"
)

// Gets a list of orders
func GetOrdersCmd(cdc *codec.Codec) *cobra.Command {
    return &cobra.Command{
        Use:    "list [address]",
        Short:  "List orders for an address",
        Args:   cobra.ExactArgs(1),
        RunE:   func(cmd *cobra.Command, args[] string) error {
            //cliCtx := context.NewCLIContext().WithCodec(cdc)

            cosmosAddress, err := sdk.AccAddressFromBech32(args[0])
            if err != nil {
                return err
            }

            msg := fmt.Sprintf("cosmos address: %s", cosmosAddress)
            return errors.New(msg)
        },
    }
}

// Places an order 
func PlaceOrderCmd(cdc *codec.Codec) *cobra.Command {
    return &cobra.Command{
        Use:    "place [type] [address] [amount] [price]",
        Short:  "Place an order for amount at price, the type must be [bid|offer]",
        Args:   cobra.ExactArgs(4),
        RunE:   func(cmd *cobra.Command, args[] string) error {
            orderType := args[0]
            if !(orderType == "bid" || orderType == "offer") {
                return errors.New("Order type must be bid or offer")
            }

            cosmosAddress, err := sdk.AccAddressFromBech32(args[1])
            if err != nil {
                return err
            }

            amount, err := strconv.Atoi(args[2])
            if err != nil {
                return err
            }

            price, err := strconv.Atoi(args[3])
            if err != nil {
                return err
            }

            msg := fmt.Sprintf("type: %s, address: %s, amount: %d, price: %d", orderType, cosmosAddress, amount, price)
            return errors.New(msg)
        },
    }
}