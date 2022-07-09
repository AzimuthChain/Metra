package cli

import (
	"github.com/mitchellh/cli"
)

// PeersCommand is the command to group the peers commands
type PeersCommand struct {
	UI cli.Ui
}

// Help implements the cli.Command interface
func (c *PeersCommand) Help() string {
	return `Usage: metra peers <subcommand>

  This command groups actions to interact with peers.
	
  List the connected peers:
  
    $ metra peers list
	
  Add a new peer by enode:
  
    $ metra peers add <enode>

  Remove a connected peer by enode:

    $ metra peers remove <enode>

  Display information about a peer:

    $ metra peers status <peer id>`
}

// Synopsis implements the cli.Command interface
func (c *PeersCommand) Synopsis() string {
	return "Interact with peers"
}

// Run implements the cli.Command interface
func (c *PeersCommand) Run(args []string) int {
	return cli.RunResultHelp
}
