package cli

import (
    "errors"
    "fmt"
    "strconv"
    "github.com/spf13/cobra"

    //"github.com/cosmos/cosmos-sdk/client"
    "github.com/cosmos/cosmos-sdk/client/context"
    "github.com/cosmos/cosmos-sdk/x/auth/client/utils"
    "github.com/cosmos/cosmos-sdk/codec"
    sdk "github.com/cosmos/cosmos-sdk/types"
    
    "github.com/cosmos/cosmos-sdk/x/auth"
    "github.com/cosmos/peggy/x/swaps/types"
)

// Gets a list of orders
func GetOrdersCmd(cdc *codec.Codec) *cobra.Command {
    cmd := &cobra.Command{
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

    return cmd
}

// Places an order 
func PlaceOrderCmd(cdc *codec.Codec) *cobra.Command {
    return &cobra.Command{
        Use:    "place [type] [address] [denom] [amount] [price]",
        Short:  "Place an order for amount at price, the type must be [bid|offer]",
        Args:   cobra.ExactArgs(5),
        RunE:   func(cmd *cobra.Command, args[] string) error {
            cliCtx := context.NewCLIContext().WithCodec(cdc)
            txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))


            orderType := args[0]
            if !(orderType == "bid" || orderType == "offer") {
                return errors.New("Order type must be bid or offer")
            }

            address, err := sdk.AccAddressFromBech32(args[1])
            if err != nil {
                return err
            }

            denom := args[2]
            amount, err := strconv.Atoi(args[3])
            if err != nil {
                return err
            }
            tokens := sdk.NewCoin(denom, sdk.NewInt(int64(amount)))

            price, err := strconv.Atoi(args[4])
            if err != nil {
                return err
            }
            price_uint := sdk.NewUint(uint64(price))


            msg := types.NewMsgOrder(orderType, tokens, price_uint, address)
            err = msg.ValidateBasic()
            if err != nil {
                return err
            }

            return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
        },
    }
}