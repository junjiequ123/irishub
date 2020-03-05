package types

import (
	service "github.com/irisnet/irishub/app/v3/service/exported"
	sdk "github.com/irisnet/irishub/types"
)

//expected Service keeper
type ServiceKeeper interface {
	RegisterResponseCallback(
		moduleName string,
		respCallback service.ResponseCallback,
	) sdk.Error

	GetRequestContext(
		ctx sdk.Context,
		requestContextID []byte,
	) (
		service.RequestContext,
		bool,
	)

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
		state service.RequestContextState,
		respThreshold uint16,
		respHandler string,
	) (
		[]byte,
		sdk.Error,
	)

	UpdateRequestContext(
		ctx sdk.Context,
		requestContextID []byte,
		providers []sdk.AccAddress,
		serviceFeeCap sdk.Coins,
		timeout int64,
		repeatedFreq uint64,
		repeatedTotal int64,
	) sdk.Error

	StartRequestContext(
		ctx sdk.Context,
		requestContextID []byte,
	) sdk.Error

	PauseRequestContext(
		ctx sdk.Context,
		requestContextID []byte,
	) sdk.Error
}
