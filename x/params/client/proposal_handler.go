package client

import (
	govclient "github.com/FnyaMing/nainaide/x/gov/client"
	"github.com/FnyaMing/nainaide/x/params/client/cli"
	"github.com/FnyaMing/nainaide/x/params/client/rest"
)

// param change proposal handler
var ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitProposal, rest.ProposalRESTHandler)
