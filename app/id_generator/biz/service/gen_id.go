package service

import (
	"context"
	"sync"
	"time"

	"github.com/yzc/orange-review/app/id_generator/conf"
	generator "github.com/yzc/orange-review/app/id_generator/kitex_gen/generator"
)

type GenIDService struct {
	ctx          context.Context
	lock         sync.Mutex
	dataCenterId int64     // 机房ID：读配置
	workerId     int64     // 服务实例ID：通过数据库主键生成
	startTime    time.Time // 系统初始时间，人为指定
	millsPassed  int64     // 上次生成ID的时间戳
	concurrency  int64     // 并发数
}

// NewGenIDService new GenIDService
func NewGenIDService(ctx context.Context) *GenIDService {
	return &GenIDService{
		ctx:          ctx,
		startTime:    time.Date(2025, 07, 15, 0, 0, 0, 0, time.Local),
		dataCenterId: conf.GetConf().DataCenterId,
		workerId:     0, // 假装是从数据库中获得的
	}
}

// 雪花算法生成全局唯一ID
// 一共64位，1位符号位，40位系统运行总毫秒数，2位机房id，11位服务实例id，10位并发数
func (s *GenIDService) Run(req *generator.GenIDRequest) (resp *generator.GenIDResponse, err error) {
	// Finish your business logic.
	resp = &generator.GenIDResponse{}

	millis := time.Since(s.startTime).Milliseconds()
	s.lock.Lock()
	defer s.lock.Unlock()

	var concurrencyValue int64
	if millis == s.millsPassed {
		if s.concurrency >= 1<<10 {
			resp.Id = -1
			resp.Message = "reach concurrency limit"
			return resp, nil
		}
		concurrencyValue = s.concurrency
		s.concurrency++
	} else {
		concurrencyValue = 0
		s.concurrency = 1
		s.millsPassed = millis
	}

	resp.Id = (millis << 23) | (s.dataCenterId << 11) | (s.workerId << 10) | concurrencyValue

	return resp, nil
}
