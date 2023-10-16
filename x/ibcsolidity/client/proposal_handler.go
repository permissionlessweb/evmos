package client

import (
	"github.com/evmos/evmos/v14/x/ibcsolidity/client/cli"

	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
)

var (
	ToggleIBCSolidityCallProposalHandler     = govclient.NewProposalHandler(cli.CmdToggleIBCSolidityCallProposal)
)
