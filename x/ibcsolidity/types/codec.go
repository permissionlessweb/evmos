package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgSolidityCall{}, "ibcsolidity/SolidityCall", nil)
	cdc.RegisterConcrete(&MsgIBCSolidityCall{}, "ibcsolidity/IBCSolidityCall", nil)
	cdc.RegisterConcrete(&MsgRegisterHostZone{}, "ibcsolidity/RegisterHostZone", nil)
	cdc.RegisterConcrete(&ToggleIBCSolidityProposal{}, "ibcsolidity/ToggleIBCSolidityProposal", nil)
	// cdc.RegisterConcrete(&MsgLiquidStake{}, "ibcsolidity/LiquidStake", nil)
	// cdc.RegisterConcrete(&MsgLSMLiquidStake{}, "ibcsolidity/LSMLiquidStake", nil)
	// cdc.RegisterConcrete(&MsgRegisterHostZone{}, "ibcsolidity/RegisterHostZone", nil)
	// cdc.RegisterConcrete(&MsgRedeemStake{}, "ibcsolidity/RedeemStake", nil)
	// cdc.RegisterConcrete(&MsgClaimUndelegatedTokens{}, "ibcsolidity/ClaimUndelegatedTokens", nil)
	// cdc.RegisterConcrete(&MsgRebalanceValidators{}, "ibcsolidity/RebalanceValidators", nil)
	// cdc.RegisterConcrete(&MsgAddValidators{}, "ibcsolidity/AddValidators", nil)
	// cdc.RegisterConcrete(&MsgChangeValidatorWeight{}, "ibcsolidity/ChangeValidatorWeight", nil)
	// cdc.RegisterConcrete(&MsgDeleteValidator{}, "ibcsolidity/DeleteValidator", nil)
	// cdc.RegisterConcrete(&AddValidatorsProposal{}, "ibcsolidity/AddValidatorsProposal", nil)

	// cdc.RegisterConcrete(&MsgRestoreInterchainAccount{}, "ibcsolidity/RestoreInterchainAccount", nil)
	// cdc.RegisterConcrete(&MsgUpdateValidatorSharesExchRate{}, "ibcsolidity/UpdateValidatorSharesExchRate", nil)
	// cdc.RegisterConcrete(&MsgUndelegateHost{}, "ibcsolidity/UndelegateHost", nil)
	// cdc.RegisterConcrete(&MsgCalibrateDelegation{}, "ibcsolidity/CalibrateDelegation", nil)
	// cdc.RegisterConcrete(&MsgUpdateInnerRedemptionRateBounds{}, "ibcsolidity/UpdateInnerRedemptionRateBounds", nil)
	// cdc.RegisterConcrete(&MsgResumeHostZone{}, "ibcsolidity/ResumeHostZone", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSolidityCall{},
		&MsgIBCSolidityCall{},
		&MsgRegisterHostZone{},
	)

	registry.RegisterImplementations((*govtypes.Content)(nil),
		&ToggleIBCSolidityProposal{},
	)


	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(Amino)
)

func init() {
	RegisterCodec(Amino)
	cryptocodec.RegisterCrypto(Amino)
	sdk.RegisterLegacyAminoCodec(Amino)
}
