package model

// 评价表（主表）
type ReviewInfo struct {
	Base
	ReviewID      int64          `gorm:"not null;default:0;column:review_id;comment:评价id;uniqueIndex"`
	Content       string         `gorm:"size:512;not null;column:content;comment:评价内容"`
	Score         int8           `gorm:"not null;default:0;column:score;comment:评分"`
	ServiceScore  int8           `gorm:"not null;default:0;column:service_score;comment:商家服务评分"`
	ExpressScore  int8           `gorm:"not null;default:0;column:express_score;comment:物流评分"`
	HasMedia      int8           `gorm:"not null;default:0;column:has_media;comment:是否有图或视频"`
	OrderID       int64          `gorm:"not null;default:0;column:order_id;comment:订单id;index"`
	SkuID         int64          `gorm:"not null;default:0;column:sku_id;comment:sku id"`
	SpuID         int64          `gorm:"not null;default:0;column:spu_id;comment:spu id"`
	StoreID       int64          `gorm:"not null;default:0;column:store_id;comment:店铺id"`
	UserID        int64          `gorm:"not null;default:0;column:user_id;comment:用户id;index"`
	Anonymous     int8           `gorm:"not null;default:0;column:anonymous;comment:是否匿名"`
	Tags          string         `gorm:"size:1024;not null;default:'';column:tags;comment:标签json"`
	PicInfo       string         `gorm:"size:1024;not null;default:'';column:pic_info;comment:媒体信息：图片"`
	VideoInfo     string         `gorm:"size:1024;not null;default:'';column:video_info;comment:媒体信息：视频"`
	Status        int8           `gorm:"not null;default:10;column:status;comment:状态:10待审核；20审核通过；30审核不通过；40隐藏"`
	IsDefault     int8           `gorm:"not null;default:0;column:is_default;comment:是否默认评价"`
	HasReply      int8           `gorm:"not null;default:0;column:has_reply;comment:是否有商家回复:0无;1有"`
	OpReason      string         `gorm:"size:512;not null;default:'';column:op_reason;comment:运营审核拒绝原因"`
	OpRemarks     string         `gorm:"size:512;not null;default:'';column:op_remarks;comment:运营备注"`
	OpUser        string         `gorm:"size:64;not null;default:'';column:op_user;comment:运营者标识"`
	GoodsSnapshoot string        `gorm:"size:2048;not null;default:'';column:goods_snapshoot;comment:商品快照信息"`
	ExtJson       string         `gorm:"size:1024;not null;default:'';column:ext_json;comment:信息扩展"`
	CtrlJson      string         `gorm:"size:1024;not null;default:'';column:ctrl_json;comment:控制扩展"`
}

func (ReviewInfo) TableName() string {
	return "review_info"
}

// 评价回复表
type ReviewReplyInfo struct {
	Base
	ReplyID   int64          `gorm:"not null;default:0;column:reply_id;comment:回复id;uniqueIndex:uk_reply_id"`
	ReviewID  int64          `gorm:"not null;default:0;column:review_id;comment:评价id;index:idx_review_id"`
	StoreID   int64          `gorm:"not null;default:0;column:store_id;comment:店铺id;index:idx_store_id"`
	Content   string         `gorm:"size:512;not null;column:content;comment:回复内容"`
	PicInfo   string         `gorm:"size:1024;not null;default:'';column:pic_info;comment:媒体信息：图片"`
	VideoInfo string         `gorm:"size:1024;not null;default:'';column:video_info;comment:媒体信息：视频"`
	ExtJson   string         `gorm:"size:1024;not null;default:'';column:ext_json;comment:信息扩展"`
	CtrlJson  string         `gorm:"size:1024;not null;default:'';column:ctrl_json;comment:控制扩展"`
}

func (ReviewReplyInfo) TableName() string {
	return "review_reply_info"
}

// 评价申诉表（商家可以申诉恶意评价）
type ReviewAppealInfo struct {
	Base
	AppealID   int64          `gorm:"not null;default:0;column:appeal_id;comment:申诉id;index:idx_appeal_id"`
	ReviewID   int64          `gorm:"not null;default:0;column:review_id;comment:评价id;uniqueIndex:uk_review_id"`
	StoreID    int64          `gorm:"not null;default:0;column:store_id;comment:店铺id;index:idx_store_id"`
	Status     int8           `gorm:"not null;default:10;column:status;comment:状态:10待审核；20申诉通过；30申诉驳回"`
	Reason     string         `gorm:"size:255;not null;column:reason;comment:申诉原因类别"`
	Content    string         `gorm:"size:255;not null;column:content;comment:申诉内容描述"`
	PicInfo    string         `gorm:"size:1024;not null;default:'';column:pic_info;comment:媒体信息：图片"`
	VideoInfo  string         `gorm:"size:1024;not null;default:'';column:video_info;comment:媒体信息：视频"`
	OpRemarks  string         `gorm:"size:512;not null;default:'';column:op_remarks;comment:运营备注"`
	OpUser     string         `gorm:"size:64;not null;default:'';column:op_user;comment:运营者标识"`
	ExtJson    string         `gorm:"size:1024;not null;default:'';column:ext_json;comment:信息扩展"`
	CtrlJson   string         `gorm:"size:1024;not null;default:'';column:ctrl_json;comment:控制扩展"`
}

func (ReviewAppealInfo) TableName() string {
	return "review_appeal_info"
}