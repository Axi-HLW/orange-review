package gateway

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/yzc/orange-review/app/gateway/biz/service"
	"github.com/yzc/orange-review/app/gateway/biz/utils"
	gateway "github.com/yzc/orange-review/app/gateway/hertz_gen/gateway"
)

// CreateReview .
// @router /v1/review [POST]
func CreateReview(ctx context.Context, c *app.RequestContext) {
	var err error
	var req gateway.CreateReviewReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &gateway.CreateReviewResp{}
	resp, err = service.NewCreateReviewService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// GetReview .
// @router /v1/review/{review_id} [GET]
func GetReview(ctx context.Context, c *app.RequestContext) {
	var err error
	var req gateway.GetReviewReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &gateway.GetReviewResp{}
	resp, err = service.NewGetReviewService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// AuditReview .
// @router /v1/review/audit [POST]
func AuditReview(ctx context.Context, c *app.RequestContext) {
	var err error
	var req gateway.AuditReviewReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &gateway.AuditReviewResp{}
	resp, err = service.NewAuditReviewService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// ReplyReview .
// @router /v1/review/reply [POST]
func ReplyReview(ctx context.Context, c *app.RequestContext) {
	var err error
	var req gateway.ReplyReviewReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &gateway.ReplyReviewResp{}
	resp, err = service.NewReplyReviewService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// AppealReview .
// @router /v1/review/appeal [POST]
func AppealReview(ctx context.Context, c *app.RequestContext) {
	var err error
	var req gateway.AppealReviewReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &gateway.AppealReviewResp{}
	resp, err = service.NewAppealReviewService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// AuditAppeal .
// @router /v1/appeal/audit [POST]
func AuditAppeal(ctx context.Context, c *app.RequestContext) {
	var err error
	var req gateway.AuditAppealReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &gateway.AuditAppealResp{}
	resp, err = service.NewAuditAppealService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// ListReviewByUserID .
// @router /v1/{user_id}/reviews [GET]
func ListReviewByUserID(ctx context.Context, c *app.RequestContext) {
	var err error
	var req gateway.ListReviewByUserIDReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &gateway.ListReviewByUserIDResp{}
	resp, err = service.NewListReviewByUserIDService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// ListReviewByStoreID .
// @router /v1/store/{store_id}/reviews [GET]
func ListReviewByStoreID(ctx context.Context, c *app.RequestContext) {
	var err error
	var req gateway.ListReviewByStoreIDReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewListReviewByStoreIDService(ctx, c).Run(&req)

	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// ListReviewBySpuID .
// @router /v1/spu/{spu_id}/reviews [GET]
func ListReviewBySpuID(ctx context.Context, c *app.RequestContext) {
	var err error
	var req gateway.ListReviewBySpuIDReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewListReviewBySpuIDService(ctx, c).Run(&req)

	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// ListReviewBySkuID .
// @router /v1/sku/{sku_id}/reviews [GET]
func ListReviewBySkuID(ctx context.Context, c *app.RequestContext) {
	var err error
	var req gateway.ListReviewBySkuIDReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewListReviewBySkuIDService(ctx, c).Run(&req)

	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
