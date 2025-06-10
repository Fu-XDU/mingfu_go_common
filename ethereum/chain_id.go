package ethereum

import (
	"context"
	"github.com/Fu-XDU/mingfu_go_common/flags"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/labstack/gommon/log"
	"math/big"
)

func GetChainId(client *ethclient.Client) (id *big.Int, err error) {
	id, err = client.ChainID(context.Background())
	if err != nil {
		return
	}

	if flags.EvmChainID.Cmp(id) != 0 {
		log.Warnf("Configured EVM chain ID (%d) does not match chain ID from RPC (%d); using RPC value", flags.EvmChainID.Uint64(), id.Uint64())
		flags.EvmChainID = id
	}

	return
}
