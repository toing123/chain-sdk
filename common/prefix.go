package common

import "fmt"

const (
	BucketChain         = "chains:%s"     //chain:链简称：链id，存储链信息
	BucketCoin          = "coins:%s"      //coin:链名：链id:币种简称，存储币种信息
	BucketCoinSetKey    = "coinset:xrc20" //coinset:xrc20，erc20 bep20 hrc20 frc20等fork以太坊链的子币集合
	BucketChainRc20Keys = "contract:%s"   //以太坊系列链下的erc20子币的key集合
)

func BuildChainKey(chain string) string {
	return fmt.Sprintf(BucketChain, chain)
}

func BuildCoinKey(coin string) string {
	return fmt.Sprintf(BucketCoin, coin)
}

func BuildChainContractsKey(chain string) string {
	return fmt.Sprintf(BucketChainRc20Keys, chain)
}
