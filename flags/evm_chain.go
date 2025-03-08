package flags

import (
	"github.com/urfave/cli/v2"
	"math/big"
)

var (
	EvmChainRPC string
	EvmChainID  *big.Int
)

var (
	evmChainRpcFlag = cli.StringFlag{
		Name:        "evm_chain.rpc",
		Usage:       "RPC URL used to connect to the RPC interface of the EVM chain",
		EnvVars:     []string{"EVM_CHAIN_RPC"},
		Destination: &EvmChainRPC,
		Required:    true,
		Action: func(ctx *cli.Context, url string) (err error) {
			return
		},
	}

	evmChainIDFlag = cli.Int64Flag{
		Name:     "evm_chain.id",
		Usage:    "The EVM chain ID which server will connect to",
		EnvVars:  []string{"EVM_CHAIN_ID"},
		Required: true,
		Action: func(ctx *cli.Context, chainID int64) (err error) {
			EvmChainID = big.NewInt(chainID)
			return
		},
	}
)

var EvmChainFlags = []cli.Flag{
	&evmChainRpcFlag,
	&evmChainIDFlag,
}
