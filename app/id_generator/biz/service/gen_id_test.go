package service

import (
	"context"
	"testing"

	generator "github.com/yzc/orange-review/app/id_generator/kitex_gen/generator"
)

func TestGenID_Run(t *testing.T) {
	ctx := context.Background()
	s := NewGenIDService(ctx)
	// init req and assert value

	for i := 0; i < 10; i++ {
		req := &generator.GenIDRequest{}
		resp, _ := s.Run(req)
		t.Logf("resp: %+v", resp)
	}

	// todo: edit your unit test
}
