package client

import (
	"github.com/FnyaMing/nainaide/x/distribution/client/cli"
	"github.com/FnyaMing/nainaide/x/distribution/client/rest"
	govclient "github.com/FnyaMing/nainaide/x/gov/client"
)

// param change proposal handler
var (
	ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitProposal, rest.ProposalRESTHandler)
)
