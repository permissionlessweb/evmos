package types

import (
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	errorsmod "cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgIBCSolidityContractCall = "ibc_solidity_contract_call"

var _ sdk.Msg = &MsgIBCSolidityContractCall{}

func NewMsgIBCSolidityContractCall(creator string, amount sdkmath.Int, lsmTokenIbcDenom string) *MsgIBCSolidityContractCall {
	return &MsgIBCSolidityContractCall{
		// TODO
	}
}

func (msg *MsgIBCSolidityContractCall) Route() string {
	return RouterKey
}

func (msg *MsgIBCSolidityContractCall) Type() string {
	return TypeMsgIBCSolidityContractCall
}

func (msg *MsgIBCSolidityContractCall) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgIBCSolidityContractCall) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgIBCSolidityContractCall) ValidateBasic() error {

	// check valid caller address 
	// check valid host-id 
	// check valid solidity msg
	// ??


	
	// // check valid caller address
	// _, err := sdk.AccAddressFromBech32(msg.Creator)
	// if err != nil {
	// 	return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	// }
	// // ensure amount is a nonzero positive integer
	// if msg.Amount.LTE(sdkmath.ZeroInt()) {
	// 	return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid amount (%v)", msg.Amount)
	// }
	// // validate host denom is not empty
	// if msg.LsmTokenIbcDenom == "" {
	// 	return errorsmod.Wrapf(ErrRequiredFieldEmpty, "LSM token denom cannot be empty")
	// }
	// // lsm token denom must be a valid asset denom matching regex
	// if err := sdk.ValidateDenom(msg.LsmTokenIbcDenom); err != nil {
	// 	return errorsmod.Wrapf(err, "invalid LSM token denom")
	// }
	return nil
}
