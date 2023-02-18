package opensea

import (
	"math"
	"math/big"
	"testing"
)

func TestCollection(t *testing.T) {
	req := CollectionRequest{
		CollectionSlug: "mineablepunks",
	}

	if collection, err := client.Collection(&req); err != nil {
		t.Error(err)
	} else {
		t.Log(EthToWei(collection.Stats.FloorPrice))
	}
}

// EthToWei converts Gwei as a float to Wei as a big int
func EthToWei(n float64) *big.Int {
	return floatToBigInt(n, 18)
}

func floatToBigInt(amount float64, decimal int64) *big.Int {
	// 6 is our smallest precision
	if decimal < 6 {
		return big.NewInt(int64(amount * math.Pow10(int(decimal))))
	}
	result := big.NewInt(int64(amount * math.Pow10(6)))
	return result.Mul(result, big.NewInt(0).Exp(big.NewInt(10), big.NewInt(decimal-6), nil))
}
