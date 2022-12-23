package opensea

import "testing"

func TestContract(t *testing.T) {

	req := ContractRequest{
		AssetContractAddress: "0x2bdbe17a6f148e686581bb3fb78186aadad9e73d",
	}

	if contract, err := client.Contract(&req); err != nil {
		t.Error(err)
	} else {
		t.Log(contract.Collection.Fees.SellerFees)
	}

}
