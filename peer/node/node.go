/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package node

import (
	"fmt"

	"github.com/hyperledger/fabric/common/flogging"
	"github.com/hyperledger/fabric/peer/common"
	"github.com/spf13/cobra"
)

const (
	nodeFuncName = "node"
	nodeCmdDes   = "Operate a peer node: start|status|reset|rollback."
)

var logger = flogging.MustGetLogger("nodeCmd")

// Cmd returns the cobra command for Node
func Cmd() *cobra.Command {
	nodeCmd.AddCommand(startCmd())
	nodeCmd.AddCommand(statusCmd())
	nodeCmd.AddCommand(resetCmd())
	nodeCmd.AddCommand(rollbackCmd())

	return nodeCmd
}

var nodeCmd = &cobra.Command{
	Use:              nodeFuncName,
	Short:            fmt.Sprint(nodeCmdDes),
	Long:             fmt.Sprint(nodeCmdDes),
	PersistentPreRun: common.InitCmd,
}
