package etherscan

import (
	"github.com/nanmu42/etherscan-api"
	"math/big"
)

// SumGasFees 计算一组以太坊交易的总 Gas 费用（总手续费）。
//
// 对于每一笔交易，费用按公式：GasPrice × GasUsed 计算，
// 所有交易费用将累加为最终总费用，单位为 Wei（以太坊最小单位）。
//
// 参数：
//
//	txs - etherscan.NormalTx 类型的交易列表，每笔交易应包含有效的 GasPrice 和 GasUsed 值。
//
// 返回值：
//
//	sum - 所有交易总费用之和，类型为 *big.Int，单位为 Wei。
//
// 注意：此函数使用 big.Int 处理任意精度运算，避免了整数溢出。
func SumGasFees(txs []etherscan.NormalTx) (sum *big.Int) {
	sum = new(big.Int)
	for _, tx := range txs {
		sum.Add(sum, new(big.Int).Mul(tx.GasPrice.Int(), big.NewInt(int64(tx.GasUsed))))
	}
	return
}
