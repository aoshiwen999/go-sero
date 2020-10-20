package tx

import (
	"fmt"
	"testing"

	"github.com/sero-cash/go-czero-import/c_type"
	"github.com/sero-cash/go-sero/zero/utils"
)

func TestT_TokenCost(t *testing.T) {
	seroCy := utils.CurrencyToUint256("SERO")
	fmt.Printf("%t\n", seroCy)
	cy := utils.CurrencyToUint256("d")
	ret := make(map[c_type.Uint256]utils.U256)
	ret[seroCy] = utils.NewU256(24)
	if cost, ok := ret[seroCy]; ok {
		add := utils.NewU256(12)
		cost.AddU(&add)
		ret[seroCy] = cost
	} else {
		cost := utils.NewU256(48)
		ret[cy] = cost
	}

	fmt.Printf("%t", ret)

}
