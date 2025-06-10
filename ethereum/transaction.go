package ethereum

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

// TransferEthTo 构造并发送一笔 ETH 转账交易，将以太币从指定私钥地址发送至目标地址。
//
// 参数：
//
//	client    - 已连接的以太坊客户端，用于查询链上数据和发送交易。
//	sender    - 发送方账户的私钥，用于签名交易。
//	receipt   - 接收方地址，类型为 *common.Address。
//	valueWei  - 转账金额（单位为 Wei）。若为 0，则自动转出账户余额减去手续费的全部金额（"转空账户"）。
//
// 返回值：
//
//	signedTx - 构造并签名完成的交易对象。
//	amount   - 实际转账金额（单位为 Wei）。
//	err      - 若发生错误，返回错误信息。
//
// 函数流程：
//  1. 获取发送方地址及账户余额；
//  2. 获取当前 nonce 和建议的 gas price；
//  3. 计算手续费（fee = gasPrice × 21000）；
//  4. 若 valueWei 为 0，表示要转出余额减手续费的全部金额；
//  5. 若 valueWei 非 0，检查余额是否足够覆盖转账金额与手续费；
//  6. 构造 Legacy 类型交易并使用 EIP-155 签名；
//  7. 广播交易至以太坊网络。
//
// 注意事项：
//   - 此函数使用标准 ETH 转账 gas 限额 21000。
//   - 如果余额不足以支付手续费或转账金额，会返回错误。
//   - 使用 types.LegacyTx 构造交易，适用于非 EIP-1559 网络或需兼容旧链的场景。
func TransferEthTo(client *ethclient.Client, sender *ecdsa.PrivateKey, receipt *common.Address, valueWei *big.Int) (signedTx *types.Transaction, amount *big.Int, err error) {
	ctx := context.Background()

	senderAddr := PrivateKeyToAddress(sender)
	// 获取账户当前余额
	balance, err := client.BalanceAt(ctx, senderAddr, nil)
	if err != nil {
		err = fmt.Errorf("failed to get balance: %w", err)
		return
	}
	if balance.Cmp(big.NewInt(0)) == 0 {
		err = fmt.Errorf("sender has no balance")
		return
	}

	// 获取 nonce
	nonce, err := client.PendingNonceAt(ctx, senderAddr)
	if err != nil {
		err = fmt.Errorf("failed to get sender nonce: %w", err)
		return
	}

	// 估算 gas price
	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		err = fmt.Errorf("failed to get gas price: %w", err)
		return
	}

	// 转账需要的 gas 限额，21000 是标准 ETH 转账固定值
	gasLimit := uint64(21000)
	fee := new(big.Int).Mul(gasPrice, new(big.Int).SetUint64(gasLimit))

	// 计算可转账金额 = 余额 - 手续费
	valueAvailable := new(big.Int).Sub(balance, fee)

	if valueWei.Cmp(big.NewInt(0)) == 0 {
		// 全部转出
		amount = valueAvailable

		if amount.Cmp(big.NewInt(0)) <= 0 {
			err = fmt.Errorf("not enough balance to cover gas fee")
			return
		}
	} else {
		// 转出 valueWei 数量的 Wei
		amount = valueWei

		// 扣除Gas费后余额不够转账
		if valueAvailable.Cmp(valueWei) == -1 {
			err = fmt.Errorf("the balance after deducting the gas fee is not enough to transfer")
			return
		}
	}

	chainID, err := client.ChainID(ctx)
	if err != nil {
		err = fmt.Errorf("failed to get network ID: %w", err)
		return
	}

	// 生成签名后的交易
	signedTx, err = types.SignNewTx(sender, types.NewEIP155Signer(chainID), &types.LegacyTx{
		Nonce:    nonce,
		GasPrice: gasPrice,
		Gas:      gasLimit,
		To:       receipt,
		Value:    amount,
		Data:     []byte{},
	})
	if err != nil {
		err = fmt.Errorf("failed to sign transaction: %w", err)
		return
	}

	// 发送交易
	err = client.SendTransaction(ctx, signedTx)
	if err != nil {
		err = fmt.Errorf("failed to send transaction: %w", err)
		return
	}

	return
}
