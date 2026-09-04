package futuapi

import (
	"context"
	"testing"

	"github.com/771991673/go-futu-api/pb/qotcommon"
	"google.golang.org/protobuf/proto"
)

func TestGetCurKLineRejectsMissingSecurity(t *testing.T) {
	api := NewFutuAPI()
	_, err := api.GetCurKLine(context.Background(), nil, 10, qotcommon.RehabType_RehabType_None, qotcommon.KLType_KLType_Day)
	if err != ErrParameters {
		t.Fatalf("got %v", err)
	}
}

func TestGetCurKLineRejectsZeroNumNotUnadjustedRehab(t *testing.T) {
	api := NewFutuAPI()
	sec := &qotcommon.Security{Market: proto.Int32(1), Code: proto.String("00700")}
	_, err := api.GetCurKLine(context.Background(), sec, 0, qotcommon.RehabType_RehabType_None, qotcommon.KLType_KLType_Day)
	if err != ErrParameters {
		t.Fatalf("zero num should be invalid, got %v", err)
	}
}

func TestRequestHistoryKLineRejectsUnknownKLType(t *testing.T) {
	api := NewFutuAPI()
	sec := &qotcommon.Security{Market: proto.Int32(1), Code: proto.String("00700")}
	_, err := api.RequestHistoryKLine(context.Background(), sec, "2020-01-01", "2020-01-02",
		qotcommon.KLType_KLType_Unknown, qotcommon.RehabType_RehabType_None,
		nil, qotcommon.KLFields_KLFields_None, nil, nil)
	if err != ErrParameters {
		t.Fatalf("unknown KLType should be invalid, got %v", err)
	}
}
