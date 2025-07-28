package service

import (
	"context"
	"testing"
	review "github.com/yzc/orange-review/rpc_gen/kitex_gen/review"
)

func TestListReviewByStoreID_Run(t *testing.T) {
	ctx := context.Background()
	s := NewListReviewByStoreIDService(ctx)
	// init req and assert value

	req := &review.ListReviewByStoreIDReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
