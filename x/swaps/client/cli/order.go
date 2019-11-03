package cli

import (
    "errors"
    "fmt"
    "strconv"
    "github.com/spf13/cobra"

    "github.com/cosmos/cosmos-sdk/client"
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
    cmd := &cobra.Command{
        Use:    "place [type] [amount] [price]",
        Short:  "Place an order for amount at price, the type must be [bid|offer]",
        Args:   cobra.ExactArgs(3),
        RunE:   func(cmd *cobra.Command, args[] string) error {
            cliCtx := context.NewCLIContext().WithCodec(cdc)
            txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))


            orderType := args[0]
            if !(orderType == "bid" || orderType == "offer") {
                return errors.New("Order type must be bid or offer")
            }

            tokens, err := sdk.ParseCoin(args[1])
            if err != nil {
                return err
            }

            price, err := strconv.Atoi(args[2])
            if err != nil {
                return err
            }
            price_uint := sdk.NewUint(uint64(price))


            msg := types.NewMsgOrder(orderType, tokens, price_uint, cliCtx.GetFromAddress())
            err = msg.ValidateBasic()
            if err != nil {
                return err
            }

            err = utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
            if err != nil {
                fmt.Println("cli", cliCtx.GetFromAddress())
                return err
            }

            return nil
        },
    }

    cmd = client.PostCommands(cmd)[0]

    return cmd

}