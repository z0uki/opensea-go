package opensea

import "testing"

func TestContract(t *testing.T) {

	keys := []string{
		"41be6f2c49714847a27780d6027f5421",
		"4c0bc196e3834d84a95c6b90ebab3b18",
		"a0672943ce854d16a94e4509aa388ef1",
		"051d1a31d15343cab4c4a4fdb123dd56",
		"d218315d31bb4d0bb4308c2cd45938a2",
		"852d96d562f646fd88184540dbd4f434",
		"57fcc64ba50648d6a93f21a0e05fc1a7",
		"2fe149332c264e008aa46f2fce2301d1",
		"40e4ba51dfb94a9489fdc848b1180d93",
		"1d93a8763d9d49bd995a4285cc7ac7aa",
		"e34a5960d01447d49a9b247d447cf076",
		"7825c4057b6044719021ed683c40ddf9",
		"560248ea4c5a46ef9f02e7ef321f6ff3",
		"50e679b3778542b39538b25379f1b9a5",
		"23eaf1789d914c6f91f1fa03ee65455c",
		"a4563dd083944ee2b6e9be567c27a805",
		"bea970cbbdae445a9f01b827f9ac227e",
	}

	for _, key := range keys {
		oc, _ := New(WithAPIKey(key))
		req := ContractRequest{
			AssetContractAddress: "0xed5af388653567af2f388e6224dc7c4b3241c544",
		}

		if contract, err := oc.Contract(&req); err != nil {
			t.Log("key: ", key, " error: ", err)
		} else {
			t.Log(contract.DevSellerFeeBasisPoints)
			t.Log(contract.Collection.Fees.SellerFees)
		}
	}

}
