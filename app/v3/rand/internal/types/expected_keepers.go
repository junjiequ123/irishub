package types

import (
	"github.com/irisnet/irishub/app/v3/service"
	"github.com/irisnet/irishub/app/v3/service/exported"
	sdk "github.com/irisnet/irishub/types"
)

//expected Service keeper
type ServiceKeeper interface {
	RegisterResponseCallback(
		moduleName string,
		respCallback exported.ResponseCallback,
	) sdk.Error

	GetRequestContext(
		ctx sdk.Context,
		requestContextID []byte,
	) (exported.RequestContext, bool)

	CreateRequestContext(
		ctx sdk.Context,
		serviceName string,
		providers []sdk.AccAddress,
		consumer sdk.AccAddress,
		input string,
		serviceFeeCap sdk.Coins,
		timeout int64,
		superMode bool,
		repeated bool,
		repeatedFrequency uint64,
		repeatedTotal int64,
		state exported.RequestContextState,
		respThreshold uint16,
		respHandler string,
	) ([]byte, sdk.Error)

	ServiceBindingsIterator(ctx sdk.Context, serviceName string) sdk.Iterator

	GetParamSet(ctx sdk.Context) service.Params
}