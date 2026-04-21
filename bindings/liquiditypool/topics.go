package liquiditypool

import "github.com/ethereum/go-ethereum/common"

// Event signature topics for LiquidityPool reconciliation subscriptions.
var (
	LiquidityPoolChannelFundedTopic   = common.HexToHash("0x174dd9cbed5b79ced30198029c21a27b4721fd0f4dd46ac813de82798e909818")
	LiquidityPoolChannelSettledTopic  = common.HexToHash("0xbb7a5a625951cb72f1e443bc7a3cd146dcdc284df4a5ed440993e5df803d2bc7")
	LiquidityPoolOperatorUpdatedTopic = common.HexToHash("0xfbe5b6cbafb274f445d7fed869dc77a838d8243a22c460de156560e8857cad03")
)
