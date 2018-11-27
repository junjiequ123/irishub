package iparam

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/irisnet/irishub/modules/params"
)

const (
	SignalParamspace  = "Sig"
	GovParamspace     = "Gov"
)

type Parameter interface {
	InitGenesis(interface{})

	GetStoreKey() []byte

	SetReadWriter(paramSpace params.Subspace)

	SaveValue(ctx sdk.Context)

	LoadValue(ctx sdk.Context) bool
}

type SignalParameter interface {
	Parameter
}

type GovParameter interface {
	Parameter

	Valid(json string) sdk.Error

	GetValueFromRawData(cdc *codec.Codec, res []byte) interface{}

	Update(ctx sdk.Context, json string)

	ToJson(string) string
}

type GovArrayParameter interface {
	GovParameter

	LoadValueByKey(ctx sdk.Context, key string) bool

	Insert(ctx sdk.Context, json string)
}
