package public

import (
	"sync"
	"time"
)

type FlowCounter struct {
	RedisFlowCountMap   map[string]*RedisFlowCountService
	RedisFlowCountSlice []*RedisFlowCountService
	Locker              sync.RWMutex
}

var FlowCounterHander *FlowCounter

func NewFlowCounter() *FlowCounter {
	return &FlowCounter{
		RedisFlowCountMap:   map[string]*RedisFlowCountService{},
		RedisFlowCountSlice: []*RedisFlowCountService{},
		Locker:              sync.RWMutex{},
	}
}

func init()  {
	FlowCounterHander=NewFlowCounter()
}

func (counter *FlowCounter)GetCounter(serverName string) (*RedisFlowCountService,error) {
	for _,item:=range counter.RedisFlowCountSlice{
		if item.AppID == serverName{
			return item,nil
		}
	}
	newCounter:=NewRedisFlowCountService(serverName,1*time.Second)
	counter.RedisFlowCountSlice=append(counter.RedisFlowCountSlice,newCounter)
	counter.Locker.Lock()
	defer counter.Locker.Unlock()
	counter.RedisFlowCountMap[serverName]=newCounter
	return newCounter,nil
}