package etherscan

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/nanmu42/etherscan-api"
	"strings"
)

// AllNormalTxByAddress 获取指定地址的所有普通交易（Normal Transactions）。
//
// 该函数通过分页调用 Etherscan 的 NormalTxByAddress 接口，遍历从区块 0 到指定的 endBlock（如果为 nil 则默认为 latest）的交易记录，
// 自动处理分页并累计结果，返回所有与该地址相关的正常交易（仅包含该地址发出的交易记录）。
//
// 注意事项：
//   - 每次请求最多获取 offset 条交易，默认设置为 1000；
//   - 函数使用 tx.BlockNumber 控制分页进度；
//   - 如果请求返回的交易数量少于 offset，说明已到达末尾，终止循环。
//
// 参数：
//   - client: 已初始化的 Etherscan 客户端对象
//   - endBlock: 指定查询的结束区块号（若为 nil 则查询至最新区块）
//   - address: 需要查询的以太坊地址（建议使用 EIP-55 格式）
//
// 返回值：
//   - txs: 包含所有查询结果的 NormalTx 列表
//   - err: 查询过程中发生的错误
//
// 文档地址：
//   - https://docs.etherscan.io/etherscan-v2/api-endpoints/accounts#get-a-list-of-normal-transactions-by-address
//   - https://docs.etherscan.io/etherscan-v2/get-an-addresss-full-transaction-history
//   - https://gist.github.com/0xV4L3NT1N3/0aecbb6a1040ecc1d4d4d31dd235e915
func AllNormalTxByAddress(client *etherscan.Client, endBlock *int, address common.Address) (txs []etherscan.NormalTx, err error) {
	nextBlock := 0
	offset := 1000
	var requestTxs []etherscan.NormalTx
	for {
		requestTxs, err = client.NormalTxByAddress(address.String(), &nextBlock, endBlock, 1, offset, false)
		if err != nil {
			if err.Error() == "etherscan server: No transactions found" {
				err = nil
			}
			return
		}

		if len(requestTxs) == 0 {
			break
		}

		nextBlock = requestTxs[len(requestTxs)-1].BlockNumber

		for _, tx := range requestTxs {
			if tx.BlockNumber != nextBlock || len(requestTxs) != offset {
				txs = append(txs, tx)
			}
		}

		if len(requestTxs) < offset {
			break
		}
	}
	return
}

func AllNormalTxFromAddress(client *etherscan.Client, endBlock *int, address common.Address) (txs []etherscan.NormalTx, err error) {
	allTxs, err := AllNormalTxByAddress(client, endBlock, address)
	if err != nil {
		return
	}

	addressString := address.String()
	for _, tx := range allTxs {
		if strings.EqualFold(addressString, tx.From) {
			txs = append(txs, tx)
		}
	}
	return
}

func AllNormalTxToAddress(client *etherscan.Client, endBlock *int, address common.Address) (txs []etherscan.NormalTx, err error) {
	allTxs, err := AllNormalTxByAddress(client, endBlock, address)
	if err != nil {
		return
	}

	addressString := address.String()
	for _, tx := range allTxs {
		if strings.EqualFold(addressString, tx.To) {
			txs = append(txs, tx)
		}
	}
	return
}
