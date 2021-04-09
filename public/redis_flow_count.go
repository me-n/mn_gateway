package public

import (
	"time"
)

//redis 流量计数服务
type RedisFlowCountService struct {
	AppID       string
	Interval    time.Duration
	QPS         int64
	Unix        int64
	TickerCount int64
	TotalCount  int64
}

func NewRedisFlowCountService(appID string,interval time.Duration) *RedisFlowCountService {
	return nil
}
