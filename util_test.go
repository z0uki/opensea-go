package opensea

import (
	"fmt"
	"math/big"
	"testing"
)

func TestCalcFees(t *testing.T) {
	fmt.Println(CalcEarnings(big.NewInt(2000000000000000000), big.NewInt(980)))
	fmt.Println(CalcOpenSeaFeeByBasePrice(big.NewInt(2000000000000000000)))
	fmt.Println(CalcFeeByBasisPoints(big.NewInt(2000000000000000000), big.NewInt(980)))
}
