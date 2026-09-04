package futuapi

import (
	"testing"

	"github.com/771991673/go-futu-api/pb/qotcommon"
	"github.com/771991673/go-futu-api/pb/trdcommon"
)

func TestExtendedMarketEnums(t *testing.T) {
	cases := []struct {
		name string
		got  int32
		want int32
	}{
		{"QotMarket AU", int32(qotcommon.QotMarket_QotMarket_AU_Security), 51},
		{"QotMarket MY", int32(qotcommon.QotMarket_QotMarket_MY_Security), 61},
		{"QotMarket CA", int32(qotcommon.QotMarket_QotMarket_CA_Security), 71},
		{"QotMarket FX", int32(qotcommon.QotMarket_QotMarket_FX_Security), 81},
		{"QotMarket CC", int32(qotcommon.QotMarket_QotMarket_CC_Security), 91},
		{"TrdMarket Crypto", int32(trdcommon.TrdMarket_TrdMarket_Crypto), 7},
		{"TrdMarket AU", int32(trdcommon.TrdMarket_TrdMarket_AU), 8},
		{"TrdMarket JP", int32(trdcommon.TrdMarket_TrdMarket_JP), 15},
		{"TrdSecMarket FX", int32(trdcommon.TrdSecMarket_TrdSecMarket_FX), 91},
		{"TrdSecMarket CC", int32(trdcommon.TrdSecMarket_TrdSecMarket_CC), 101},
		{"TrdCategory Crypto", int32(trdcommon.TrdCategory_TrdCategory_Crypto), 3},
		{"TimeInForce IOC", int32(trdcommon.TimeInForce_TimeInForce_IOC), 2},
		{"RehabType None (unadjusted)", int32(qotcommon.RehabType_RehabType_None), 0},
	}
	for _, tc := range cases {
		if tc.got != tc.want {
			t.Errorf("%s: got %d want %d", tc.name, tc.got, tc.want)
		}
	}
}
