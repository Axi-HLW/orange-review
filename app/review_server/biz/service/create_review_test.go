package service

import (
	"context"
	"testing"
	review "github.com/yzc/orange-review/app/review_server/kitex_gen/review"
)

func TestCreateReview_Run(t *testing.T) {
	ctx := context.Background()
	s := NewCreateReviewService(ctx)
	// init req and assert value

	req := &review.CreateReviewRequest{
		UserId:        1,
		OrderId:       1,
		ItemScore:     6,
		ServiceScore:  5,
		ExpressScore:  5,
		Content:       "12345678",
		PicInfo:       "12345678",
		VideoInfo:     "12345678",
		Anonymous:     true,
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
