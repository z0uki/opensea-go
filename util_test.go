package opensea

import (
	"fmt"
	"math/big"
	"testing"
)

func TestCalcFees(t *testing.T) {
	fmt.Println(CalcEarnings(big.NewInt(1450000000000000000), big.NewInt(50)))
	fmt.Println(CalcFeeByBasisPoints(big.NewInt(1450000000000000000), big.NewInt(50)))
}
