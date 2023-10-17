package ibcsolidity

import (
	"fmt"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"

	"github.com/evmos/evmos/v14/x/ibcsolidity/keeper"
	"github.com/evmos/evmos/v14/x/ibcsolidity/types"
)

// Handles ibcsolidity transactions
func NewMessageHandler(k keeper.Keeper) sdk.Handler {
	msgServer := keeper.NewMsgServerImpl(k)

	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		_ = ctx

		switch msg := msg.(type) {
		case *types.MsgIBCSolidityCall:
			res, err := msgServer.IBCSolidityCall(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgRegisterHostZone:
			res, err := msgServer.RegisterHostZone(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgRestoreInterchainAccount:
			res, err := msgServer.RestoreInterchainAccount(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgResumeHostZone:
			res, err := msgServer.ResumeHostZone(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, errorsmod.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

// // Handles stakeibc gov proposals
// func NewIbcSolidityProposalHandler(k keeper.Keeper) govtypes.Handler {
// 	return func(ctx sdk.Context, content govtypes.Content) error {
// 		switch c := content.(type) {
// 		case *types.Proposal:
// 			return k.AddValidatorsProposal(ctx, c)
// 		case *types.ToggleLSMProposal:
// 			return k.ToggleLSMProposal(ctx, c)

// 		default:
// 			return errorsmod.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized stakeibc proposal content type: %T", c)
// 		}
// 	}
// }
