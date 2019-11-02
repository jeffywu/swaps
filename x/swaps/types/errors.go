package types

import (
    "fmt"
    sdk "github.com/cosmos/cosmos-sdk/types"
)

// CodeType local code type
type CodeType = sdk.CodeType

// Exported code type numbers
const (
    DefaultCodespace sdk.CodespaceType = ModuleName

    CodeInvalidOrderType    CodeType = 1
    CodeInvalidTokenAmount  CodeType = 2
    CodeInvalidPrice        CodeType = 3
)


func ErrInvalidOrderType(orderType string) sdk.Error {
    msg := fmt.Sprintf("invalid order type: %s", orderType)

    return sdk.NewError(DefaultCodespace, CodeInvalidOrderType, msg)
}

func ErrInvalidTokenAmount() sdk.Error {
    return sdk.NewError(DefaultCodespace, CodeInvalidTokenAmount, "token amount must be > 0")
}

func ErrInvalidPrice() sdk.Error {
    return sdk.NewError(DefaultCodespace, CodeInvalidPrice, "price amount must be > 0")
}