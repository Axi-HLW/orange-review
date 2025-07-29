package model

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/redis/go-redis/v9"
	"github.com/yzc/orange-review/app/review/biz/dal/es"
	myRedis "github.com/yzc/orange-review/app/review/biz/dal/redis"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
)

const TableNameReviewInfo = "review_info"

// ReviewInfo 评价表
type ReviewInfo struct {
	ID             int64      `gorm:"column:id;primaryKey;autoIncrement:true;comment:主键" json:"id"`                                  // 主键
	CreateBy       string     `gorm:"column:create_by;not null;comment:创建方标识" json:"create_by"`                                      // 创建方标识
	UpdateBy       string     `gorm:"column:update_by;not null;comment:更新方标识" json:"update_by"`                                      // 更新方标识
	CreateAt       time.Time  `gorm:"column:create_at;not null;comment:创建时间" json:"create_at" time_format:"2006-01-02 15:04:05.999"` // 创建时间
	UpdateAt       time.Time  `gorm:"column:update_at;not null;comment:更新时间" json:"update_at" time_format:"2006-01-02 15:04:05.999"` // 更新时间
	DeleteAt       *time.Time `gorm:"column:delete_at;comment:逻辑删除标记" json:"delete_at"`                                              // 逻辑删除标记
	Version        int32      `gorm:"column:version;not null;comment:乐观锁标记" json:"version"`                                          // 乐观锁标记
	ReviewID       int64      `gorm:"column:review_id;not null;comment:评价id" json:"review_id"`                                       // 评价id
	Content        string     `gorm:"column:content;not null;comment:评价内容" json:"content"`                                           // 评价内容
	Score          int32      `gorm:"column:score;not null;comment:评分" json:"score"`                                                 // 评分
	ServiceScore   int32      `gorm:"column:service_score;not null;comment:商家服务评分" json:"service_score"`                             // 商家服务评分
	ExpressScore   int32      `gorm:"column:express_score;not null;comment:物流评分" json:"express_score"`                               // 物流评分
	HasMedia       int32      `gorm:"column:has_media;not null;comment:是否有图或视频" json:"has_media"`                                    // 是否有图或视频
	OrderID        int64      `gorm:"column:order_id;not null;comment:订单id" json:"order_id"`                                         // 订单id
	SkuID          int64      `gorm:"column:sku_id;not null;comment:sku id" json:"sku_id"`                                           // sku id
	SpuID          int64      `gorm:"column:spu_id;not null;comment:spu id" json:"spu_id"`                                           // spu id
	StoreID        int64      `gorm:"column:store_id;not null;comment:店铺id" json:"store_id"`                                         // 店铺id
	UserID         int64      `gorm:"column:user_id;not null;comment:用户id" json:"user_id"`                                           // 用户id
	Anonymous      int32      `gorm:"column:anonymous;not null;comment:是否匿名" json:"anonymous"`                                       // 是否匿名
	Tags           string     `gorm:"column:tags;not null;comment:标签json" json:"tags"`                                               // 标签json
	PicInfo        string     `gorm:"column:pic_info;not null;comment:媒体信息：图片" json:"pic_info"`                                      // 媒体信息：图片
	VideoInfo      string     `gorm:"column:video_info;not null;comment:媒体信息：视频" json:"video_info"`                                  // 媒体信息：视频
	Status         int32      `gorm:"column:status;not null;default:10;comment:状态:10待审核；20审核通过；30审核不通过；40隐藏" json:"status"`          // 状态:10待审核；20审核通过；30审核不通过；40隐藏
	IsDefault      int32      `gorm:"column:is_default;not null;comment:是否默认评价" json:"is_default"`                                   // 是否默认评价
	HasReply       int32      `gorm:"column:has_reply;not null;comment:是否有商家回复:0无;1有" json:"has_reply"`                              // 是否有商家回复:0无;1有
	OpReason       string     `gorm:"column:op_reason;not null;comment:运营审核拒绝原因" json:"op_reason"`                                   // 运营审核拒绝原因
	OpRemarks      string     `gorm:"column:op_remarks;not null;comment:运营备注" json:"op_remarks"`                                     // 运营备注
	OpUser         string     `gorm:"column:op_user;not null;comment:运营者标识" json:"op_user"`                                          // 运营者标识
	GoodsSnapshoot string     `gorm:"column:goods_snapshoot;not null;comment:商品快照信息" json:"goods_snapshoot"`                         // 商品快照信息
	ExtJSON        string     `gorm:"column:ext_json;not null;comment:信息扩展" json:"ext_json"`                                         // 信息扩展
	CtrlJSON       string     `gorm:"column:ctrl_json;not null;comment:控制扩展" json:"ctrl_json"`                                       // 控制扩展
}

type ReviewInfoES struct {
	ID             int64  `gorm:"column:id;primaryKey;autoIncrement:true;comment:主键" json:"id,string"`                           // 主键
	CreateBy       string `gorm:"column:create_by;not null;comment:创建方标识" json:"create_by"`                                      // 创建方标识
	UpdateBy       string `gorm:"column:update_by;not null;comment:更新方标识" json:"update_by"`                                      // 更新方标识
	CreateAt       string `gorm:"column:create_at;not null;comment:创建时间" json:"create_at" time_format:"2006-01-02 15:04:05.999"` // 创建时间
	UpdateAt       string `gorm:"column:update_at;not null;comment:更新时间" json:"update_at" time_format:"2006-01-02 15:04:05.999"` // 更新时间
	DeleteAt       string `gorm:"column:delete_at;comment:逻辑删除标记" json:"delete_at"`                                              // 逻辑删除标记
	Version        int32  `gorm:"column:version;not null;comment:乐观锁标记" json:"version,string"`                                   // 乐观锁标记
	ReviewID       int64  `gorm:"column:review_id;not null;comment:评价id" json:"review_id,string"`                                // 评价id
	Content        string `gorm:"column:content;not null;comment:评价内容" json:"content"`                                           // 评价内容
	Score          int32  `gorm:"column:score;not null;comment:评分" json:"score,string"`                                          // 评分
	ServiceScore   int32  `gorm:"column:service_score;not null;comment:商家服务评分" json:"service_score,string"`                      // 商家服务评分
	ExpressScore   int32  `gorm:"column:express_score;not null;comment:物流评分" json:"express_score,string"`                        // 物流评分
	HasMedia       int32  `gorm:"column:has_media;not null;comment:是否有图或视频" json:"has_media,string"`                             // 是否有图或视频
	OrderID        int64  `gorm:"column:order_id;not null;comment:订单id" json:"order_id,string"`                                  // 订单id
	SkuID          int64  `gorm:"column:sku_id;not null;comment:sku id" json:"sku_id,string"`                                    // sku id
	SpuID          int64  `gorm:"column:spu_id;not null;comment:spu id" json:"spu_id,string"`                                    // spu id
	StoreID        int64  `gorm:"column:store_id;not null;comment:店铺id" json:"store_id,string"`                                  // 店铺id
	UserID         int64  `gorm:"column:user_id;not null;comment:用户id" json:"user_id,string"`                                    // 用户id
	Anonymous      int32  `gorm:"column:anonymous;not null;comment:是否匿名" json:"anonymous,string"`                                // 是否匿名
	Tags           string `gorm:"column:tags;not null;comment:标签json" json:"tags"`                                               // 标签json
	PicInfo        string `gorm:"column:pic_info;not null;comment:媒体信息：图片" json:"pic_info"`                                      // 媒体信息：图片
	VideoInfo      string `gorm:"column:video_info;not null;comment:媒体信息：视频" json:"video_info"`                                  // 媒体信息：视频
	Status         int32  `gorm:"column:status;not null;default:10;comment:状态:10待审核；20审核通过；30审核不通过；40隐藏" json:"status,string"`   // 状态:10待审核；20审核通过；30审核不通过；40隐藏
	IsDefault      int32  `gorm:"column:is_default;not null;comment:是否默认评价" json:"is_default,string"`                            // 是否默认评价
	HasReply       int32  `gorm:"column:has_reply;not null;comment:是否有商家回复:0无;1有" json:"has_reply,string"`                       // 是否有商家回复:0无;1有
	OpReason       string `gorm:"column:op_reason;not null;comment:运营审核拒绝原因" json:"op_reason"`                                   // 运营审核拒绝原因
	OpRemarks      string `gorm:"column:op_remarks;not null;comment:运营备注" json:"op_remarks"`                                     // 运营备注
	OpUser         string `gorm:"column:op_user;not null;comment:运营者标识" json:"op_user"`                                          // 运营者标识
	GoodsSnapshoot string `gorm:"column:goods_snapshoot;not null;comment:商品快照信息" json:"goods_snapshoot"`                         // 商品快照信息
	ExtJSON        string `gorm:"column:ext_json;not null;comment:信息扩展" json:"ext_json"`                                         // 信息扩展
	CtrlJSON       string `gorm:"column:ctrl_json;not null;comment:控制扩展" json:"ctrl_json"`                                       // 控制扩展
}

// TableName ReviewInfo's table name
func (*ReviewInfo) TableName() string {
	return TableNameReviewInfo
}

// CreateReview 创建评价
func CreateReview(db *gorm.DB, ctx context.Context, review *ReviewInfo) (*ReviewInfo, error) {
	tx := db.Model(&ReviewInfo{}).Create(review)
	return review, tx.Error
}

// GetReviewByOrderId 根据订单id获取评价
func GetReviewByOrderId(db *gorm.DB, ctx context.Context, orderId int64) (*[]ReviewInfo, error) {
	var reviews []ReviewInfo
	tx := db.Model(&ReviewInfo{}).Where("order_id = ?", orderId).Find(&reviews)
	return &reviews, tx.Error
}

// ListReviewByStoreID 根据storeID 分页查询评价
func ListReviewByStoreID(ctx context.Context, storeID int64, offset, limit int) ([]*ReviewInfo, error) {
	// return listReviewByStoreIDFromES(ctx, storeID, offset, limit) // 第一版直接查ES
	return listReviewByStoreIDFromESV2(ctx, storeID, offset, limit) // 第二版增加缓存和singleflight
}

func listReviewByStoreIDFromES(ctx context.Context, storeID int64, offset, limit int) ([]*ReviewInfo, error) {
	// 去ES里面查询评价
	resp, err := es.ESClient.Search().
		Index("review").
		From(offset).
		Size(limit).
		Query(&types.Query{
			Bool: &types.BoolQuery{
				Filter: []types.Query{
					{
						Term: map[string]types.TermQuery{
							"store_id": {Value: storeID},
						},
					},
				},
			},
		}).
		Do(ctx)
	// klog.Infof("--> es search: %+v %+v\n", resp, err)
	if err != nil {
		return nil, err
	}
	klog.Infof("es result total:%v\n", resp.Hits.Total.Value)

	// 反序列化数据
	list := make([]*ReviewInfo, 0, resp.Hits.Total.Value)

	for _, hit := range resp.Hits.Hits {
		tmp := &ReviewInfoES{}
		if err := json.Unmarshal(hit.Source_, tmp); err != nil {
			klog.Errorf("json.Unmarshal(hit.Source_, tmp) failed, err:%v", err)
			continue
		}
		// klog.CtxInfof(ctx, "es result: %+v\n", tmp)
		ri := &ReviewInfo{
			ID:             tmp.ID,
			CreateBy:       tmp.CreateBy,
			UpdateBy:       tmp.UpdateBy,
			Version:        tmp.Version,
			ReviewID:       tmp.ReviewID,
			OrderID:        tmp.OrderID,
			SkuID:          tmp.SkuID,
			SpuID:          tmp.SpuID,
			StoreID:        tmp.StoreID,
			UserID:         tmp.UserID,
			Content:        tmp.Content,
			Score:          tmp.Score,
			ServiceScore:   tmp.ServiceScore,
			ExpressScore:   tmp.ExpressScore,
			HasMedia:       tmp.HasMedia,
			Anonymous:      tmp.Anonymous,
			Tags:           tmp.Tags,
			PicInfo:        tmp.PicInfo,
			VideoInfo:      tmp.VideoInfo,
			Status:         tmp.Status,
			IsDefault:      tmp.IsDefault,
			HasReply:       tmp.HasReply,
			OpReason:       tmp.OpReason,
			OpRemarks:      tmp.OpRemarks,
			OpUser:         tmp.OpUser,
			GoodsSnapshoot: tmp.GoodsSnapshoot,
			ExtJSON:        tmp.ExtJSON,
			CtrlJSON:       tmp.CtrlJSON,
		}
		layout := "2006-01-02 15:04:05.999"
		createAt, _ := time.Parse(layout, tmp.CreateAt)
		updateAt, _ := time.Parse(layout, tmp.UpdateAt)
		ri.CreateAt = createAt
		ri.UpdateAt = updateAt
		list = append(list, ri)
	}

	return list, nil
}

func listReviewByStoreIDFromESV2(ctx context.Context, storeID int64, offset, limit int) ([]*ReviewInfo, error) {
	key := fmt.Sprintf("review:%d:%d:%d", storeID, offset, limit)
	dataByte, err := getDataBySingleFlight(ctx, key)
	if err != nil {
		return nil, err
	}
	var reviewList []*ReviewInfo
	err = json.Unmarshal(dataByte, &reviewList)
	if err != nil {
		return nil, err
	}
	return reviewList, nil
}

var g singleflight.Group

func getDataBySingleFlight(ctx context.Context, key string) ([]byte, error) {
	val, err, _ := g.Do(key, func() (interface{}, error) {
		data, err := getDataFromCache(ctx, key)
		if err == nil {
			klog.Infof("getDataBySingleFlight cache hit. key: %s", key)
			return data, nil
		}
		klog.Infof("getDataFromCache err: %v", err)
		if errors.Is(err, redis.Nil) {
			klog.Infof("getDataBySingleFlight cache miss. key: %s", key)
			_, storeID, offset, limit, err := parseKey(key)
			if err != nil {
				return nil, err
			}
			reviewList, err := listReviewByStoreIDFromES(ctx, storeID, offset, limit)
			if err == nil {
				data, err = json.Marshal(reviewList)
				if err != nil {
					return nil, err
				}
				err = setCache(ctx, key, data)
				return data, err
			}
			return nil, err
		}
		klog.Infof("getDataBySingleFlight err: %v", err)
		return nil, err
	})
	// klog.Infof("getDataBySingleFlight ret: v:%v err:%v shared:%v\n", val, err, shared)
	if err != nil {
		return nil, err
	}
	return val.([]byte), nil
}

func setCache(ctx context.Context, key string, data []byte) error {
	return myRedis.RedisClient.Set(ctx, key, data, time.Second*60).Err()
}

func getDataFromCache(ctx context.Context, key string) ([]byte, error) {
	klog.Infof("getDataFromCache key: %s\n", key)
	return myRedis.RedisClient.Get(ctx, key).Bytes()
}

func parseKey(key string) (string, int64, int, int, error) {
	values := strings.Split(key, ":")
	if len(values) < 4 {
		return "", 0, 0, 0, errors.New("invalid key")
	}
	index, storeID, offsetStr, limitStr := values[0], values[1], values[2], values[3]
	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		return "", 0, 0, 0, err
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		return "", 0, 0, 0, err
	}
	storeIDInt, err := strconv.ParseInt(storeID, 10, 64)
	if err != nil {
		return "", 0, 0, 0, err
	}
	return index, storeIDInt, offset, limit, nil
}
