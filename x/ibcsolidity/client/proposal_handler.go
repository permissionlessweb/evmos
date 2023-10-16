package client

import (
	"github.com/Stride-Labs/stride/v16/x/stakeibc/client/cli"

	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
)

var (
	ToggleIBCSolidityCallProposalHandler     = govclient.NewProposalHandler(cli.CmdToggleIBCSolidityCallProposal)
)
