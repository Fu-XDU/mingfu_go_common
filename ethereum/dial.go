package ethereum

import (
	"context"
	"github.com/Fu-XDU/mingfu_go_common/constants"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

func Dial(rawurl string) (ethClient *ethclient.Client, err error) {
	return ethclient.Dial(rawurl)
}

func DialWithAuth(rawurl string, auth string) (ethClient *ethclient.Client, err error) {
	clientOption := rpc.WithHeader(constants.Authorization, auth)

	c, err := rpc.DialOptions(context.Background(), rawurl, clientOption)
	if err != nil {
		return
	}

	ethClient = ethclient.NewClient(c)
	return
}
