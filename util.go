package opensea

import (
	"math/big"
)

func CalcEarnings(price, creatorBasisPoints *big.Int) int64 {
	return price.Int64() - CalcFeeByBasisPoints(price, creatorBasisPoints)
}

func CalcOpenSeaFeeByBasePrice(price *big.Int) int64 {
	m := new(big.Int)
	m.SetInt64(10000)

	feeBasisPoints := new(big.Int)
	feeBasisPoints.SetInt64(250)
	pfee := m.Div(feeBasisPoints.Mul(price, feeBasisPoints), m).Int64()
	return pfee
}

func CalcFeeByBasisPoints(price, basisPoints *big.Int) int64 {
	m := new(big.Int)
	m.SetInt64(10000)

	bp := new(big.Int)
	bp.SetInt64(basisPoints.Int64())

	cfee := m.Div(bp.Mul(price, bp), m).Int64()
	return cfee
}
